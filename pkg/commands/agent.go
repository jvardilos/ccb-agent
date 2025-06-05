package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Agent(ctx context.Context, c *cli.Command) error {
	fmt.Println(c)
	fmt.Println("hello world")

	return nil
}
