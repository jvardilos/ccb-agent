package commands

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/jvardilos/ccb-agent/pkg/fileio"
	"github.com/jvardilos/ccb-agent/pkg/url"
	"github.com/urfave/cli/v3"
)

func Agent(ctx context.Context, c *cli.Command) error {

	log.Warn("test")

	out, err := url.GetFromDB("GET", "https://httpbin.org/anything")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := fileio.WriteJSON(out, "name.json"); err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}

	return nil
}
