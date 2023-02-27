package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mczaplinski/joe-client/joectl/pkg/joe"
	"github.com/urfave/cli/v2"
)

const (
	flagFile   string = "file"
	flagData   string = "data"
	flagOutput string = "output"
	flagAPIKey string = "apikey"
)

func main() {
	app := &cli.App{
		Name:     "joectl",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Marcel Czaplinski",
				Email: "marcel.czaplinski@gmx.de",
			},
		},
		Copyright: "(c) 2023 Marcel Czaplinski",
		HelpName:  "joectl",
		Usage:     "tool to communicate with the JOE API",
		ArgsUsage: "(-f <file> -o <output>)",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagFile, Value: "", Usage: "filepath to read", Aliases: []string{"f"}},
			&cli.StringFlag{Name: flagData, Value: "", Usage: "in-line OpenTrans2.1 XML data", Aliases: []string{"d"}},
			&cli.StringFlag{Name: flagOutput, Value: "", Usage: "filepath to output", Aliases: []string{"o"}},
			&cli.StringFlag{Name: flagAPIKey, Value: "", Usage: "API key", Aliases: []string{"k"}},
		},
		HideHelp:    false,
		HideVersion: false,
		Commands: []*cli.Command{
			{
				Name:            "order",
				SkipFlagParsing: true,
				Aliases:         []string{"o"},
				Usage:           "place an order",
				Action:          cmdOrder,
			},
			{
				Name:            "list",
				SkipFlagParsing: true,
				Aliases:         []string{"l"},
				Usage:           "lists all orders",
				Action:          cmdListOrders,
			},
			{
				Name:            "get",
				SkipFlagParsing: true,
				Aliases:         []string{"g"},
				Usage:           "get an order",
				Action:          cmdGetOrder,
			},
		},
		CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Fprintf(cCtx.App.Writer, "Command %q doesn't exist.\n", command)
		},
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(cCtx.App.Writer, "Wrong usage: %#v\n", err)
			return nil
		},
		// Action: cmdOrder,
	}

	app.Run(os.Args)
}

func processInput(cCtx *cli.Context, needsInput bool) ([]byte, string, error) {
	// validate input flags
	if needsInput {
		if cCtx.String(flagFile) == "" && cCtx.String(flagData) == "" {
			return nil, "", fmt.Errorf("no input data. please specify input using --file or --data")
		}

		if cCtx.String(flagFile) != "" && cCtx.String(flagData) != "" {
			return nil, "", fmt.Errorf("cannot specify both --file and --data")
		}
	}

	// read content from file or argument
	var (
		content []byte
		err     error
	)
	if cCtx.String(flagFile) != "" {
		content, err = os.ReadFile(cCtx.String(flagFile))
		if err != nil {
			return nil, "", err
		}
	} else {
		content = []byte(cCtx.String(flagData))
	}

	// api key from flag or env
	apiKey := cCtx.String(flagAPIKey)
	if apiKey == "" {
		apiKey = os.Getenv("JOE_API_KEY")
		if apiKey == "" {
			return nil, "", fmt.Errorf("no API key specified. please specify using --apikey or set JOE_API_KEY environment variable")
		}
	}

	return content, apiKey, nil
}

func processOutput(cCtx *cli.Context, result []byte) error {
	// write result to file or stdout (default)
	if cCtx.String(flagOutput) != "" {
		fmt.Printf("output to file %s\n", cCtx.String(flagOutput))
		err := os.WriteFile(cCtx.String(flagOutput), result, 0644)
		if err != nil {
			return err
		}
	} else {
		fmt.Fprintf(cCtx.App.Writer, "%s\n", result)
	}

	return nil
}

func cmdOrder(cCtx *cli.Context) error {
	// process input
	content, apiKey, err := processInput(cCtx, true)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// send to JOE
	result, err := joe.Order(content, apiKey)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// process output
	if err = processOutput(cCtx, result); err != nil {
		return cli.Exit(err.Error(), 1)
	}

	return nil
}

func cmdListOrders(cCtx *cli.Context) error {
	// process input
	_, apiKey, err := processInput(cCtx, false)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// send to JOE
	result, err := joe.Orders(apiKey)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// process output
	if err = processOutput(cCtx, result); err != nil {
		return cli.Exit(err.Error(), 1)
	}

	return nil
}

func cmdGetOrder(cCtx *cli.Context) error {
	// process input
	_, apiKey, err := processInput(cCtx, false)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// send to JOE
	result, err := joe.Get(cCtx.Args().First(), apiKey)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// process output
	if err = processOutput(cCtx, result); err != nil {
		return cli.Exit(err.Error(), 1)
	}

	return nil
}
