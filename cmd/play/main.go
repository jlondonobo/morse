package main

import "github.com/jlondonobo/morse/internal/sound"

func main() {
	conf := &sound.Config{Pitch: 700, Wpm: 20, WaveType: "sine"}
	sound.Play("....   .. / -   ....   .   . ..   .", conf)
}

// Next steps: Use this in the CLI to produce single words
// Then allow creating sentences.
