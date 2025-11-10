package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"joselondono/morse/internal/sound"

	"github.com/urfave/cli/v3"
)

var morseDictionary = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": "-.-",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	" ": "/",
}

func toMorse(s string) string {
	// edge case: spaces
	var sb strings.Builder
	for _, v := range s {
		chr := string(unicode.ToLower(v))
		repl, ok := morseDictionary[chr]
		if !ok {
			log.Fatalf("Unrecognized character %s", string(v))
		}

		sb.WriteString(repl)
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	var translateInput string
	var play bool
	var output string

	cmd := &cli.Command{
		Name:  "morse",
		Usage: "beep beep beeep",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "input",
				Destination: &translateInput,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "sound",
				Value:       false,
				Usage:       "plays on speakers",
				Destination: &play,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "outputs as wav file",
				Destination: &output,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			seq := toMorse(translateInput)
			fmt.Println(seq)
			if play {
				sound.Play(seq)
			}
			if len(output) > 0 {
				sound.Write(seq, output)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
