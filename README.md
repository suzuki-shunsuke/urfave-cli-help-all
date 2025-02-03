# urfave-cli-help-all

Go library to show the help of all commands of CLIs built with urfave/cli/v2

## How To Use

e.g. https://github.com/suzuki-shunsuke/tfcmt/blob/9120e07afae826dab4dc28bebd0ce1304350cc5d/pkg/cli/app.go#L117

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

```console
$ tfcmt help-all
```

Output: https://github.com/suzuki-shunsuke/tfcmt-docs/blob/05b02ee32aeb2935dc096e8e4c1eb76f4f830e4b/docs/usage.md?plain=1#L9-L94

e.g. https://github.com/suzuki-shunsuke/tfcmt-docs/blob/05b02ee32aeb2935dc096e8e4c1eb76f4f830e4b/scripts/generate-usage.sh
