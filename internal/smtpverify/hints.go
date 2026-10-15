package smtpverify

import (
	"fmt"
	"time"
)

// Hints are user guidance printed just before a FAILED verdict. They key
// off the reply code + enhanced code the Telnyx submission tier publishes
// (its SMTP protocol contract); the server's own reply text is always
// printed verbatim and is never rewritten.

type hintKey struct {
	stage    Stage // "" matches any stage
	code     int
	enhanced string
}

var replyHints = map[hintKey]string{
	{"", 538, "5.7.11"}:           "The server requires STARTTLS before AUTH. This command always issues STARTTLS first, so something on the network path is interfering with TLS.",
	{"", 535, "5.7.8"}:            "The API key was rejected. The username is always the literal \"apikey\" (sent by this command); check that --api-key or TELNYX_API_KEY holds a valid Telnyx v2 API key.",
	{StageAuth, 535, "5.7.1"}:     "The API key is valid but lacks the email send capability. Use a key that is allowed to send email.",
	{StageData, 535, "5.7.1"}:     "Send authorization failed when the message was submitted. Check the key's email send capability, then retry.",
	{StageAuth, 421, "4.7.0"}:     "Authentication is temporarily unavailable or throttled (repeated failures from your IP or for this key). Nothing was sent; wait and retry.",
	{"", 421, "4.7.0"}:            "The server is temporarily refusing service. Nothing was sent; retry later.",
	{StageMailFrom, 421, "4.7.0"}: hintEntitlementLost,
	{StageRcptTo, 421, "4.7.0"}:   hintEntitlementLost,
	{StageData, 421, "4.7.0"}:     hintEntitlementLost,
	{"", 421, ""}:                 "The server closed the session (for example after its 30 s idle timeout between commands). Nothing was accepted; retry.",
	{"", 552, ""}:                 "The message is larger than the server's advertised SIZE limit.",
	{"", 421, "4.4.2"}:            "The session reached its maximum lifetime. Nothing was sent in the expired session; run the command again.",
	{"", 530, "5.7.0"}:            "The server treated the session as unauthenticated.",
	{"", 550, "5.7.1"}:            "The sender was refused. Use a --from address on a verified sending domain that belongs to the same account as the API key; the reply text says which check failed.",
	{"", 451, "4.3.0"}:            "Temporary failure on the server side; nothing was accepted. Retry later.",
	{StageRcptTo, 452, "4.5.3"}:   "Too many recipients in this session.",
	{StageData, 452, "4.5.3"}:     "The account's send limit or quota is exhausted. Retry later.",
	{"", 552, "5.3.4"}:            "The message is larger than the server's advertised SIZE limit.",
	{"", 554, "5.6.0"}:            "The message failed the send API's validation.",
}

func hintForReply(stage Stage, r Reply) string {
	if h, ok := replyHints[hintKey{stage, r.Code, r.Enhanced}]; ok {
		return h
	}
	if h, ok := replyHints[hintKey{"", r.Code, r.Enhanced}]; ok {
		return h
	}
	switch {
	case r.Code >= 400 && r.Code < 500:
		return "Temporary failure (4xx): retrying later may succeed."
	case r.Code >= 500:
		return "Permanent failure (5xx): retrying unchanged will fail again."
	default:
		return fmt.Sprintf("Unexpected %d reply at %s.", r.Code, stage)
	}
}

const (
	hintNoSTARTTLS      = "Telnyx requires STARTTLS on port 587 and rejects AUTH before TLS with \"538 5.7.11 Encryption required\". Check --host and --port, and that nothing on your network strips STARTTLS."
	hintNoUUID          = "The server accepted the message, but not with the Telnyx acceptance reply \"250 2.0.0 Ok: queued as <uuid>\", so there is no message UUID to look up (a UUID elsewhere in a reply is not one). Check that --host points at a Telnyx SMTP submission host."
	hintIncomplete      = "The server's reply was cut off before its line ending (SMTP replies end in CRLF), so it does not count, even if it began like an acceptance. If it came after the message was sent, the message may still have been queued: check the portal message log before retrying."
	hintTLS             = "The TLS handshake failed. Check that --host is the name on the server's certificate."
	hintEntitlementLost = "The server ended the session after authentication: the key may have been revoked or lost its email send capability, or the service is temporarily unavailable. Retry; if it persists, check the key."
	hintInjected        = "The server sent unexpected data before the TLS handshake: something on the network path may be tampering with the connection. Credentials were not sent."
	hintDNS             = "The host name does not resolve. Check --host."
	hintDial            = "Could not connect. Check --host and --port, and that your network allows outbound connections to that port."
	hintWithheld        = "The server sent more reply text than could be checked for the API key in time, so the transcript is not shown. Check --host and --port."
)

func hintTimeout(d time.Duration) string {
	return fmt.Sprintf("No reply within %s. Check connectivity, or raise --timeout / --data-timeout.", d)
}
