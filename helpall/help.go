package helpall

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type Options struct {
	// Header is a function to return the header of the help.
	Header func(app *cli.App) (string, error)
	// Footer is a function to return the footer of the help.
	Footer func(app *cli.App) (string, error)
}

// New returns a new command to show the help of all commands.
func New(app *cli.App, opts *Options) *cli.Command {
	return &cli.Command{
		Name:   "help-all",
		Hidden: true,
		Usage:  "show all help",
		Action: func(ctx *cli.Context) error {
			if opts != nil && opts.Header != nil {
				header, err := opts.Header(app)
				if err != nil {
					return err
				}
				fmt.Fprintln(app.Writer, header)
			}
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
			ignoredCommands := map[string]struct{}{
				"help":     {},
				"help-all": {},
			}
			for _, cmd := range app.Commands {
				if _, ok := ignoredCommands[cmd.Name]; ok {
					continue
				}
				fmt.Fprintf(app.Writer, "\n## %s %s\n\n", app.Name, cmd.Name)
				fmt.Fprintln(app.Writer, "```console")
				fmt.Fprintf(app.Writer, "$ %s %s --help\n", app.Name, cmd.Name)
				if err := cli.ShowCommandHelp(ctx, cmd.Name); err != nil {
					return err
				}
				fmt.Fprintln(app.Writer, "```")
			}
			if opts != nil && opts.Footer != nil {
				footer, err := opts.Footer(app)
				if err != nil {
					return err
				}
				fmt.Fprintln(app.Writer, footer)
			}
			return nil
		},
	}
}
