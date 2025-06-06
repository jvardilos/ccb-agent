package main

import (
	"context"
	"os"
	"strings"

	"github.com/apex/log"
	apexcli "github.com/apex/log/handlers/cli"
	"github.com/jvardilos/ccb-agent/pkg/commands"
	cli "github.com/urfave/cli/v3"
)

func parseLogLevel(s string) log.Level {
	switch strings.ToLower(s) {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn", "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}

func initLog(level string) {
	log.SetHandler(apexcli.Default)
	log.SetLevel(parseLogLevel(level))
	log.WithField("level", level).Info("Logger initialized")
}

func main() {
	app := &cli.Command{
		Name:  "ccb-agent",
		Usage: "CCB AI Agent is your wizard to grabbing your data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Usage:    "CHMS username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "CHMS password",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "church",
				Aliases:  []string{"c"},
				Usage:    "Your CCB subdomain",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Usage:   "Logging level (debug, info, warn, error, fatal)",
				Value:   "info",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			initLog(c.String("log-level"))
			return commands.Agent(ctx, c)
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.WithError(err).Fatal("Failed with error")
	}
}
