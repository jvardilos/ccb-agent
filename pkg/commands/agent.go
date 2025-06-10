package commands

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/jvardilos/ccb-agent/pkg/url"
	"github.com/urfave/cli/v3"
)

func Agent(ctx context.Context, c *cli.Command) error {
	log.Info(c.String("username"))
	log.Debug("this is a debug")
	log.Info("this is an info")
	log.Warn("this is a warning")
	log.Error("this is an error")
	// log.Fatal("this is a fatal")

	out, err := url.GetFromDB("GET", "https://httpbin.org/anything")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Println(out)

	return nil
}
