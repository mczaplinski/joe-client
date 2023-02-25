package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mczaplinski/joe-client/joeconvert/pkg/convert"

	"github.com/urfave/cli/v2"
)

const (
	flagFile   string = "file"
	flagData   string = "data"
	flagOutput string = "output"
)

func main() {
	app := &cli.App{
		Name:     "joeconvert",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Marcel Czaplinski",
				Email: "marcel.czaplinski@gmx.de",
			},
		},
		Copyright: "(c) 2023 Marcel Czaplinski",
		HelpName:  "joeconvert",
		Usage:     "convert other json formats for usage with joectl",
		ArgsUsage: "-f <file> (-o <output>)",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagFile, Value: "", Usage: "filepath to read", Aliases: []string{"f"}},
			&cli.StringFlag{Name: flagData, Value: "", Usage: "in-line json data", Aliases: []string{"d"}},
			&cli.StringFlag{Name: flagOutput, Value: "", Usage: "filepath to output", Aliases: []string{"o"}},
		},
		HideHelp:    false,
		HideVersion: false,
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
		Action: cmdConvert,
	}

	app.Run(os.Args)
}

func cmdConvert(cCtx *cli.Context) error {
	// validate input flags
	if cCtx.String(flagFile) == "" && cCtx.String(flagData) == "" {
		return cli.Exit("no input data. please specify input using --file or --data", 1)
	}

	if cCtx.String(flagFile) != "" && cCtx.String(flagData) != "" {
		return cli.Exit("cannot specify both --file and --data", 1)
	}

	// read content from file or argument
	var (
		content []byte
		err     error
	)
	if cCtx.String(flagFile) != "" {
		content, err = os.ReadFile(cCtx.String(flagFile))
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
	} else {
		content = []byte(cCtx.String(flagData))
	}

	// convert content
	result, err := convert.Convert(content)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// write result to file or stdout (default)
	if cCtx.String(flagOutput) != "" {
		fmt.Printf("output to file %s\n", cCtx.String(flagOutput))
		err := os.WriteFile(cCtx.String(flagOutput), result, 0644)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
	} else {
		fmt.Fprintf(cCtx.App.Writer, "%s", result)
	}

	return nil
}
