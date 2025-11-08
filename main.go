package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

var morseDictionary = map[string]string{
	"a": "·-",
	"b": "-···",
	"c": "-·-·",
	"d": "-··",
	"e": "·",
	"f": "··-·",
	"g": "--·",
	"h": "····",
	"i": "··",
	"j": "·---",
	"k": "-·-",
	"l": "·-··",
	"m": "--",
	"n": "-·",
	"o": "---",
	"p": "·--·",
	"q": "--·-",
	"r": "-·-",
	"s": "···",
	"t": "-",
	"u": "··-",
	"v": "···-",
	"w": "·--",
	"x": "-··-",
	"y": "-·--",
	"z": "--··",
	"1": "·----",
	"2": "··---",
	"3": "···--",
	"4": "····-",
	"5": "·····",
	"6": "-····",
	"7": "--···",
	"8": "---··",
	"9": "----·",
	"0": "-----",
	" ": "/",
}

func toMorse(s string) string {
	// edge case: spaces
	var sb strings.Builder
	for _, v := range s {
		sb.WriteString(morseDictionary[string(v)])
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	var translateInput string

	cmd := &cli.Command{
		Name:  "morse",
		Usage: "beep beep beeep",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "input",
				Destination: &translateInput,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println(toMorse(translateInput))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
