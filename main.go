package main

import (
	"errors"
	"github.com/jfrog/jfrog-cli-core/plugins"
	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"strconv"
	"strings"
)

func main() {
	plugins.JfrogPluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "hello-example"
	app.Description = "Easily greet anyone."
	app.Version = "0.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		{
			Name:        "hello",
			Description: "Says Hello.",
			Aliases:     []string{"hi"},
			Arguments: []components.Argument{
				{
					Name:        "addressee",
					Description: "The name of the person you would like to greet.",
				},
			},
			Flags: []components.Flag{
				components.BoolFlag{
					Name:         "shout",
					Usage:        "Makes output all uppercase.",
					DefaultValue: false,
				},
				components.StringFlag{
					Name:         "repeat",
					Usage:        "Greets multiple times.",
					DefaultValue: "1",
				},
			},
			Action: func(c *components.Context) error {
				return helloCmd(c)
			},
		},
	}
}

func helloCmd(c *components.Context) error {
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	greet := "Hello " + c.Arguments[0] + "!"

	if c.GetBoolFlagValue("shout") {
		greet = strings.ToUpper(greet)
	}

	repeat, err := strconv.Atoi(c.GetStringFlagValue("repeat"))
	if err != nil {
		return err
	}
	for i := 0; i < repeat; i++ {
		log.Output(greet)
	}

	return nil
}
