package main

import (
	"MyApp/internal/app"
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.1"

func main() {
	ctx := context.Background()
	app := cli.NewApp()
	app.Name = "MyAPP"
	app.Version = VERSION
	app.Usage = "GIN + Ent ORM + Wire"
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(VERSION))
		},
	}
}
