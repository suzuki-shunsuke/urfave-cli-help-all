package helpall

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

type Options struct{}

// New returns a new command to show the help of all commands.
func New(_ *Options) *cli.Command {
	return &cli.Command{
		Name:   "help-all",
		Hidden: true,
		Usage:  "show all help",
		Action: func(ctx *cli.Context) error {
			app := ctx.App
			fmt.Fprintln(app.Writer, "```console")
			fmt.Fprintf(app.Writer, "$ %s --help\n", app.Name)
			if err := cli.ShowAppHelp(ctx); err != nil {
				return err
			}
			fmt.Fprintln(app.Writer, "```")

			subcommands := ctx.Command.Subcommands
			ctx.Command.Subcommands = nil
			defer func() {
				ctx.Command.Subcommands = subcommands
			}()

			cmdName := "help-all"
			if ctx.Command != nil && ctx.Command.Name != "" {
				cmdName = ctx.Command.Name
			}

			for _, cmd := range app.Commands {
				if cmd.Name == cmdName {
					continue
				}
				if err := showCommandHelp(ctx, cmd, app.Name, 2); err != nil { //nolint:mnd
					return err
				}
			}
			return nil
		},
	}
}

func showCommandHelp(ctx *cli.Context, cmd *cli.Command, parentCommand string, level int) error {
	if cmd.Hidden || cmd.Name == "help" {
		return nil
	}
	command := parentCommand + " " + cmd.Name
	fmt.Fprintf(ctx.App.Writer, "\n%s %s\n\n", strings.Repeat("#", level), command)
	fmt.Fprintln(ctx.App.Writer, "```console")
	fmt.Fprintf(ctx.App.Writer, "$ %s --help\n", command)

	if err := cli.ShowCommandHelp(ctx, cmd.Name); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "```")

	a := ctx.Command
	ctx.Command = cmd
	defer func() {
		ctx.Command = a
	}()

	level2 := level + 1
	for _, subcmd := range cmd.Subcommands {
		if err := showCommandHelp(ctx, subcmd, command, level2); err != nil {
			return err
		}
	}
	return nil
}
