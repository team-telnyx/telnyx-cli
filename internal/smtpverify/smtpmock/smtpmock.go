// Package smtpmock is an in-process, scriptable SMTP submission server for
// tests: a local listener with a self-signed certificate that speaks
// STARTTLS and records every command it receives. Default replies are the
// Telnyx SMTP submission contract's published strings.
package smtpmock

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"math/big"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

// Contract reply strings (Telnyx SMTP submission contract).
const (
	Banner          = "220 email-submission.telnyx.com ESMTP email_submission"
	ReplyAuthOK     = "235 2.7.0 Authentication successful"
	ReplyEncryption = "538 5.7.11 Encryption required"
	ReplyQuit       = "221 Bye"
	QueuedUUID      = "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"
	ReplyQueued     = "250 2.0.0 Ok: queued as " + QueuedUUID
)

// Echo in AuthReply is replaced with the AUTH command line as received,
// credentials included: some proxies and MTAs echo it back.
const Echo = "{echo}"

// Script controls the server's replies. Zero values give a fully
// successful contract-shaped session. Reply fields may hold multi-line
// replies separated by "\n".
type Script struct {
	Banner        string
	NoSTARTTLS    bool   // don't advertise STARTTLS (and offer AUTH in plaintext)
	NoAuthAdvert  bool   // don't advertise AUTH after STARTTLS
	EHLOReply     string // replaces the whole EHLO reply
	STARTTLSReply string
	// GarbageAfterSTARTTLS is written in plaintext right after the STARTTLS
	// reply, before the handshake (a command-injection attempt).
	GarbageAfterSTARTTLS string
	AuthReply            string
	MailReply            string
	RcptReply            string
	DataReply            string // reply to DATA itself (354 by default)
	DataEndReply         string // reply after the terminating "."
	// AuthPartial replaces AuthReply with bytes written as is, with no
	// final line ending; the server then closes the connection (a reply
	// cut off mid-line).
	AuthPartial string
	// DataEndPartial replaces DataEndReply with bytes written as is, with no
	// line ending: the server then closes the connection or, with HangOn
	// ".", stalls (a reply cut off mid-line).
	DataEndPartial string
	QuitReply      string
	// HangOn: after receiving this verb, never reply (timeout tests).
	HangOn string
	// CloseOn: after receiving this verb, close the connection without replying.
	CloseOn string
	// CloseAfter: after replying to this verb, close the connection.
	CloseAfter string
	// StallTLS: answer STARTTLS with 220, then never start the handshake.
	StallTLS bool
	// CertName replaces the certificate's DNS name ("localhost").
	CertName string
}

// Command is one line the server received.
type Command struct {
	Line string
	TLS  bool
}

// Server is a running mock.
type Server struct {
	Addr string // 127.0.0.1:port
	Host string
	Port int
	// ClientTLS trusts the server's certificate (for 127.0.0.1).
	ClientTLS *tls.Config

	script Script
	ln     net.Listener
	tlsCfg *tls.Config

	mu       sync.Mutex
	commands []Command
	bodies   []string
	raw      []byte // every byte received off the wire, TLS records included
	wg       sync.WaitGroup
}

// Start launches a mock on 127.0.0.1; it is closed on test cleanup.
func Start(t testing.TB, script Script) *Server {
	t.Helper()
	cert, pool := selfSigned(t, or(script.CertName, "localhost"))
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("smtpmock: listen: %v", err)
	}
	addr := ln.Addr().(*net.TCPAddr)
	s := &Server{
		Addr:      addr.String(),
		Host:      addr.IP.String(),
		Port:      addr.Port,
		ClientTLS: &tls.Config{RootCAs: pool},
		script:    script,
		ln:        ln,
		tlsCfg:    &tls.Config{Certificates: []tls.Certificate{cert}},
	}
	s.wg.Add(1)
	go s.serve()
	t.Cleanup(func() {
		ln.Close()
		s.wg.Wait()
	})
	return s
}

// Wait stops accepting connections and waits for every session to end
// (sessions end when the client closes), so the recordings are complete.
func (s *Server) Wait() {
	s.ln.Close()
	s.wg.Wait()
}

// Raw returns every byte received on the wire, before and after STARTTLS
// (after it, only TLS records), including bytes a client sent after a
// failed handshake, which the server drains rather than drops.
func (s *Server) Raw() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]byte(nil), s.raw...)
}

type teeConn struct {
	net.Conn
	s *Server
}

func (c teeConn) Read(p []byte) (int, error) {
	n, err := c.Conn.Read(p)
	c.s.mu.Lock()
	c.s.raw = append(c.s.raw, p[:n]...)
	c.s.mu.Unlock()
	return n, err
}

// Commands returns the received command lines in order.
func (s *Server) Commands() []Command {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Command(nil), s.commands...)
}

// Verbs returns the upper-cased first word of each received command.
func (s *Server) Verbs() []string {
	var verbs []string
	for _, c := range s.Commands() {
		verb, _, _ := strings.Cut(c.Line, " ")
		verbs = append(verbs, strings.ToUpper(verb))
	}
	return verbs
}

// Bodies returns each message received via DATA (CRLF line endings,
// dot-unstuffed, without the terminating ".").
func (s *Server) Bodies() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.bodies...)
}

func (s *Server) serve() {
	defer s.wg.Done()
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.handle(teeConn{conn, s})
		}()
	}
}

func (s *Server) record(line string, secure bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.commands = append(s.commands, Command{Line: line, TLS: secure})
}

func or(v, def string) string {
	if v != "" {
		return v
	}
	return def
}

func (s *Server) handle(raw net.Conn) {
	conn := raw
	defer func() { conn.Close() }()
	_ = conn.SetDeadline(time.Now().Add(30 * time.Second))
	r := bufio.NewReader(conn)
	secure, authed := false, false
	closeNow := false
	send := func(reply string) bool { // one write per reply, like a real server
		_, err := conn.Write([]byte(strings.ReplaceAll(reply, "\n", "\r\n") + "\r\n"))
		if closeNow {
			conn.Close()
		}
		return err == nil
	}
	if !send(or(s.script.Banner, Banner)) {
		return
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\r\n")
		s.record(line, secure)
		verb, _, _ := strings.Cut(line, " ")
		verb = strings.ToUpper(verb)
		if verb == s.script.HangOn {
			continue // read on until the client gives up and closes
		}
		if verb == s.script.CloseOn {
			return
		}
		closeNow = s.script.CloseAfter != "" && verb == s.script.CloseAfter
		switch verb {
		case "EHLO":
			reply := s.script.EHLOReply
			if reply == "" {
				caps := []string{"email-submission.telnyx.com", "SIZE 25000000", "8BITMIME", "SMTPUTF8", "PIPELINING"}
				if !secure && !s.script.NoSTARTTLS {
					caps = append(caps, "STARTTLS")
				}
				if (secure && !s.script.NoAuthAdvert) || s.script.NoSTARTTLS {
					caps = append(caps, "AUTH PLAIN LOGIN")
				}
				for i, c := range caps {
					sep := "-"
					if i == len(caps)-1 {
						sep = " "
					}
					caps[i] = "250" + sep + c
				}
				reply = strings.Join(caps, "\n")
			}
			send(reply)
		case "STARTTLS":
			if secure || s.script.NoSTARTTLS {
				send("502 5.5.1 Unrecognized command")
				continue
			}
			reply := or(s.script.STARTTLSReply, "220 2.0.0 Ready to start TLS")
			if s.script.GarbageAfterSTARTTLS != "" {
				reply += "\n" + s.script.GarbageAfterSTARTTLS // same write as the 220
			}
			send(reply)
			if !strings.HasPrefix(reply, "220") {
				continue
			}
			if s.script.StallTLS {
				_, _ = io.Copy(io.Discard, raw) // record until the client hangs up
				return
			}
			tc := tls.Server(raw, s.tlsCfg)
			if err := tc.Handshake(); err != nil {
				// Keep reading so anything the client sends after a failed
				// handshake (e.g. plaintext AUTH) is recorded in Raw.
				_ = raw.SetDeadline(time.Now().Add(5 * time.Second))
				_, _ = io.Copy(io.Discard, raw)
				return
			}
			conn, r, secure = tc, bufio.NewReader(tc), true
		case "AUTH":
			if !secure && !s.script.NoSTARTTLS {
				send(ReplyEncryption)
				continue
			}
			if s.script.AuthPartial != "" {
				_, _ = conn.Write([]byte(s.script.AuthPartial))
				return
			}
			reply := strings.ReplaceAll(or(s.script.AuthReply, ReplyAuthOK), Echo, line)
			authed = strings.HasPrefix(reply, "235")
			send(reply)
		case "MAIL":
			if !authed {
				send("530 5.7.0 Authentication required")
				continue
			}
			send(or(s.script.MailReply, "250 2.1.0 Ok"))
		case "RCPT":
			if !authed {
				send("530 5.7.0 Authentication required")
				continue
			}
			send(or(s.script.RcptReply, "250 2.1.5 Ok"))
		case "DATA":
			reply := or(s.script.DataReply, "354 End data with <CR><LF>.<CR><LF>")
			send(reply)
			if !strings.HasPrefix(reply, "354") {
				continue
			}
			var body strings.Builder
			for {
				l, err := r.ReadString('\n')
				if err != nil {
					return
				}
				if l == ".\r\n" {
					break
				}
				body.WriteString(strings.TrimPrefix(l, "."))
			}
			s.mu.Lock()
			s.bodies = append(s.bodies, body.String())
			s.mu.Unlock()
			if s.script.DataEndPartial != "" {
				_, _ = conn.Write([]byte(s.script.DataEndPartial))
				if s.script.HangOn != "." {
					return
				}
			}
			if s.script.HangOn == "." {
				continue
			}
			send(or(s.script.DataEndReply, ReplyQueued))
		case "QUIT":
			send(or(s.script.QuitReply, ReplyQuit))
			return
		case "RSET", "NOOP":
			send("250 Ok")
		default:
			send("502 5.5.1 Unrecognized command")
		}
	}
}

func selfSigned(t testing.TB, dnsName string) (tls.Certificate, *x509.CertPool) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("smtpmock: key: %v", err)
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "smtpmock"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
		DNSNames:     []string{dnsName},
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("smtpmock: cert: %v", err)
	}
	leaf, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("smtpmock: parse cert: %v", err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(leaf)
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key, Leaf: leaf}, pool
}
