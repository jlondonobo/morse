package main

import (
	"context"
	"errors"
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
	DefaultPitch          = 700
	DefaultWordsPerMinute = 20
	DefaultTone           = "sine"
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

var ErrInvalidParameter = errors.New("invalid parameter")

func main() {
	var wg sync.WaitGroup

	var translateInput string
	var play bool
	var output bool
	var file string
	var pitch uint16
	var wpm uint8
	var waveType string

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
			&cli.Uint8Flag{
				Name:        "speed",
				Usage:       "sets the ouput speed in words per minute, higher means faster",
				Destination: &wpm,
				Value:       DefaultWordsPerMinute,
			},
			&cli.StringFlag{
				Name:        "tone",
				Usage:       "sets the ouput tone, possible values are 'sine', 'triangle', 'sawtooth', 'square'",
				Destination: &waveType,
				Value:       DefaultTone,
				Action: func(ctxt context.Context, cmd *cli.Command, v string) error {
					_, ok := sound.ToneGenerator[v]
					if !ok {
						return fmt.Errorf("invalid tone '%v' must be one of 'sine' 'triangle' 'sawtooth' 'square'", v)
					}
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if pitch < 300 || pitch > 1000 {
				return fmt.Errorf("%w: pitch must be between 300 and 1000, got %d", ErrInvalidParameter, pitch)
			}
			if wpm < 5 || wpm > 40 {
				return fmt.Errorf("%w: speed must be between 5 and 40, got %d", ErrInvalidParameter, wpm)
			}

			seq := toMorse(translateInput)
			fmt.Println(seq)
			if play {
				wg.Go(func() { sound.Play(seq, pitch, wpm, waveType) })
			}
			if output && (len(file) > 0) {
				log.Fatal("Cannot use --output and --output-file at the same time.")
				return nil
			}
			if output {
				wg.Go(func() { sound.Write(seq, "sound.wav", pitch, wpm, waveType) })
			} else if len(file) > 0 {
				wg.Go(func() {
					sound.Write(seq, file, pitch, wpm, waveType)
				})
			}

			wg.Wait()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
