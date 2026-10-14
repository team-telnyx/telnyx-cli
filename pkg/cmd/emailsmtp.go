package cmd

// Hand-written, not generated from the OpenAPI spec: `telnyx email smtp
// verify` exercises SMTP submission end to end with the caller's own API
// key, so there is no REST resource behind it.

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify"
	"github.com/urfave/cli/v3"
)

var emailCommand = newEmailCommand()

// emailSMTPVerifyTLSConfig lets tests trust an in-process mock's
// self-signed certificate; nil (system roots) everywhere else.
var emailSMTPVerifyTLSConfig *tls.Config

func newEmailCommand() *cli.Command {
	return &cli.Command{
		Name:     "email",
		Usage:    "Email tools that are not REST resources",
		Category: "TOOLS",
		Suggest:  true,
		Commands: []*cli.Command{
			{
				Name:     "smtp",
				Usage:    "SMTP submission tools",
				Suggest:  true,
				Commands: []*cli.Command{newEmailSMTPVerifyCommand()},
			},
		},
	}
}

func newEmailSMTPVerifyCommand() *cli.Command {
	return &cli.Command{
		Name:  "verify",
		Usage: "Send one real message through SMTP submission and print the transcript",
		Description: "Connects to --host, upgrades with STARTTLS (required: credentials are never sent\n" +
			"without TLS), authenticates with AUTH PLAIN using the literal username \"apikey\"\n" +
			"and your API key as the password, then submits one short test message from\n" +
			"--from to --to (default: the --from address). Every server reply is printed\n" +
			"verbatim, stage by stage. The final line is either\n" +
			"\"VERIFIED: message UUID <uuid>\" or \"FAILED at <stage>: <reply>\", followed by\n" +
			"a hint. The session stops at the first failing reply. Exit status is 0 only\n" +
			"when verified.\n\n" +
			"The API key comes from the global --api-key flag or TELNYX_API_KEY.\n" +
			"This sends a real message, billed like any other send.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Usage:   "SMTP submission host (required; there is no default host yet)",
				Sources: cli.EnvVars("TELNYX_SMTP_HOST"),
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "SMTP submission port (STARTTLS)",
				Value: smtpverify.DefaultPort,
			},
			&cli.StringFlag{
				Name:  "from",
				Usage: "Sender address, on a verified sending domain on the API key's account (required)",
			},
			&cli.StringFlag{
				Name:  "to",
				Usage: "Recipient address (default: the --from address)",
			},
			&cli.DurationFlag{
				Name:  "timeout",
				Usage: "How long to wait for each server reply",
				Value: smtpverify.DefaultCommandTimeout,
			},
			&cli.DurationFlag{
				Name:  "data-timeout",
				Usage: "How long to wait for the server to accept the message after DATA",
				Value: smtpverify.DefaultDataTimeout,
			},
		},
		Action:          handleEmailSMTPVerify,
		HideHelpCommand: true,
	}
}

func handleEmailSMTPVerify(ctx context.Context, cmd *cli.Command) error {
	if args := cmd.Args().Slice(); len(args) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", args)
	}
	cfg, err := emailSMTPVerifyConfig(cmd)
	if err != nil {
		return err
	}
	out := cmd.Root().Writer
	if out == nil {
		out = os.Stdout
	}
	res, err := smtpverify.Run(ctx, cfg, out)
	if err != nil {
		return err
	}
	if !res.Verified {
		return fmt.Errorf("SMTP verification failed at %s", res.Stage)
	}
	return nil
}

func emailSMTPVerifyConfig(cmd *cli.Command) (smtpverify.Config, error) {
	host := strings.TrimSpace(cmd.String("host"))
	if host == "" {
		return smtpverify.Config{}, errors.New("--host is required (or set TELNYX_SMTP_HOST): there is no default SMTP submission host yet")
	}
	apiKey := cmd.String("api-key")
	if apiKey == "" {
		return smtpverify.Config{}, errors.New("an API key is required: pass --api-key or set TELNYX_API_KEY")
	}
	from := strings.TrimSpace(cmd.String("from"))
	if from == "" {
		return smtpverify.Config{}, errors.New("--from is required: a sender address on a verified sending domain on the API key's account")
	}
	return smtpverify.Config{
		Host:           host,
		Port:           cmd.Int("port"),
		APIKey:         apiKey,
		From:           from,
		To:             cmd.String("to"),
		CommandTimeout: cmd.Duration("timeout"),
		DataTimeout:    cmd.Duration("data-timeout"),
		TLSConfig:      emailSMTPVerifyTLSConfig,
	}, nil
}
