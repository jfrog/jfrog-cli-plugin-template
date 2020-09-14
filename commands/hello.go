package commands

import (
	"errors"
	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"os"
	"strconv"
	"strings"
)

func GetHelloCommand() components.Command {
	return components.Command{
		Name:        "hello",
		Description: "Says Hello.",
		Aliases:     []string{"hi"},
		Arguments:   getHelloArguments(),
		Flags:       getHelloFlags(),
		EnvArgs:     getHelloEnvVar(),
		Action: func(c *components.Context) error {
			return helloCmd(c)
		},
	}
}

func getHelloArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags() []components.Flag {
	return []components.Flag{
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
	}
}

func getHelloEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_EXAMPLE_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

func helloCmd(c *components.Context) error {
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	prefix := os.Getenv("HELLO_EXAMPLE_GREET_PREFIX")
	if prefix == "" {
		prefix = "A new greet from your plugin template: "
	}
	greet := prefix + "Hello " + c.Arguments[0] + "!"

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
