package commands

import (
	"context"
	"fmt"

	"github.com/jvardilos/ccb-agent/pkg/fileio"
	"github.com/jvardilos/ccb-agent/pkg/llm"
	"github.com/urfave/cli/v3"
)

func Agent(ctx context.Context, c *cli.Command) error {

	out, err := llm.Ask("Hi There, please respond")
	if err != nil {
		return err
	}

	if err := fileio.WriteJSON(out, "name.json"); err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}

	return nil
}
