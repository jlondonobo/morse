package main

import "joselondono/morse/internal/sound"

func main() {
	sound.Play("....   .. / -   ....   .   . ..   .", 700, 20, "sine")
}

// Next steps: Use this in the CLI to produce single words
// Then allow creating sentences.
