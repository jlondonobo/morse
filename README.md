# Go Morse

A go-based CLI for translating strings to morse code.

## Features

- Output string representation to `stdout`.
- Output sound representation to a `.wav` file via the `--output` flag.
- Supports short flags like `-s` (sound) and `-o` (output).

## Quickstart

This command will print the morse code string representation to standard output and write the sound to the file `carlitos.wav`.

```
morse 'Let's go Carlos lets go' -so 'carlitos.wav'
```

## Roadmap

- [x] Enable saving sound to file
- [x] Make functions run in parallel
- [x] Set up short-version flags
- [x] Extend punctuation
- [ ] Add a default file name for better ergonomics
- [ ] Enable editing sound qualities
  - [ ] Speed
  - [ ] Pitch
  - [ ] Volume

- [ ] Improve efficiency of steream construction / duplication.
