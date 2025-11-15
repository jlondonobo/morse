package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"joselondono/morse/internal/sound"

	"sync"

	"github.com/urfave/cli/v3"
)

const (
	DefaultPitch = 700
)

var morseDictionary = map[string]string{
	"a":  ".-",
	"b":  "-...",
	"c":  "-.-.",
	"d":  "-..",
	"e":  ".",
	"f":  "..-.",
	"g":  "--.",
	"h":  "....",
	"i":  "..",
	"j":  ".---",
	"k":  "-.-",
	"l":  ".-..",
	"m":  "--",
	"n":  "-.",
	"o":  "---",
	"p":  ".--.",
	"q":  "--.-",
	"r":  ".-.",
	"s":  "...",
	"t":  "-",
	"u":  "..-",
	"v":  "...-",
	"w":  ".--",
	"x":  "-..-",
	"y":  "-.--",
	"z":  "--..",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	" ":  "/",
	".":  ".-.-.-",
	",":  "--..--",
	"?":  "..--..",
	"'":  ".----.",
	"/":  "-..-.",
	"(":  "-.--.",
	")":  "-.--.-",
	":":  "---...",
	"=":  "-...-",
	"+":  ".-.-.",
	"-":  "-....-",
	"\"": ".-..-.",
	"@":  ".--.-.",
}

func toMorse(s string) string {
	var sb strings.Builder
	for _, v := range s {
		chr := string(unicode.ToLower(v))
		repl, ok := morseDictionary[chr]
		if !ok {
			log.Fatalf(
				"Character `%s` not defined in the International Morse Code Recommendation.",
				string(v),
			)
		}

		sb.WriteString(repl)
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	var wg sync.WaitGroup

	var translateInput string
	var play bool
	var output bool
	var file string
	var pitch uint16

	cmd := &cli.Command{
		UseShortOptionHandling: true,
		Name:                   "morse",
		Usage:                  "beep beep beeep",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "input",
				Destination: &translateInput,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "sound",
				Aliases:     []string{"s"},
				Value:       false,
				Usage:       "plays on speakers",
				Destination: &play,
			},
			&cli.BoolFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "writes sound to sound.wav file",
				Destination: &output,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"f"},
				Usage:       "writes sound to wav file",
				Destination: &file,
			},
			&cli.Uint16Flag{
				Name:        "pitch",
				Aliases:     []string{"p"},
				Usage:       "sets the ouput pitch",
				Destination: &pitch,
				Value:       DefaultPitch,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			seq := toMorse(translateInput)
			fmt.Println(seq)
			if play {
				wg.Go(func() { sound.Play(seq, pitch) })
			}
			if output && (len(file) > 0) {
				log.Fatal("Cannot use --output and --output-file at the same time.")
				return nil
			}
			if output {
				wg.Go(func() { sound.Write(seq, "sound.wav", pitch) })
			} else if len(file) > 0 {
				wg.Go(func() { sound.Write(seq, file, pitch) })
			}

			wg.Wait()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
