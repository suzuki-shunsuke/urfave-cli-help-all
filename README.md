# urfave-cli-help-all

[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/urfave-cli-help-all/main/LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/urfave-cli-help-all.svg)](https://pkg.go.dev/github.com/suzuki-shunsuke/urfave-cli-help-all/helpall)

A Go library to show the help of all commands of CLIs built with [urfave/cli/v2](https://pkg.go.dev/github.com/urfave/cli/v2).
This is useful if you want to put the usage of CLI built with urfave/cli/v2 into the document.

## How To Use

Using this library, you can add a command `help-all` showing the help of all commands.

```go
import (
	"github.com/suzuki-shunsuke/urfave-cli-help-all/helpall"
	"github.com/urfave/cli/v2"
)

app := cli.NewApp()
app.Commands = []*cli.Command{
	// ...
	helpall.New(nil),
}
```

This example adds the command to [tfcmt](https://github.com/suzuki-shunsuke/tfcmt):

e.g. https://github.com/suzuki-shunsuke/tfcmt/blob/9120e07afae826dab4dc28bebd0ce1304350cc5d/pkg/cli/app.go#L117

`help-all` command outputs the help message.
You can put it into the document.

```console
$ tfcmt help-all
```

Output: https://github.com/suzuki-shunsuke/tfcmt-docs/blob/05b02ee32aeb2935dc096e8e4c1eb76f4f830e4b/docs/usage.md?plain=1#L9-L94

e.g. https://github.com/suzuki-shunsuke/tfcmt-docs/blob/05b02ee32aeb2935dc096e8e4c1eb76f4f830e4b/scripts/generate-usage.sh

By default, the command is hidden, so it isn't shown in the help message.
The function `helpall.New()` returns a `*cli.Command`. You can customize the returned value.

e.g. Change the command name

```go
cmd := helpall.New(nil)
cmd.Name = "help-markdown" // Change the command name
app.Commands = append(app.Commands, cmd)
```
