package commands

import (
	"context"

	"github.com/apex/log"
	"github.com/urfave/cli/v3"
)

func Agent(ctx context.Context, c *cli.Command) error {
	log.Info(c.String("username"))
	log.Debug("this is a debug")
	log.Info("this is an info")
	log.Warn("this is a warning")
	log.Error("this is an error")
	// log.Fatal("this is a fatal")

	return nil
}
