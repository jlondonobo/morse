<p align="center">
<img width="1628" height="677" alt="Untitled" src="https://github.com/user-attachments/assets/ce58515f-e25e-4e50-aa08-1438f3676f7f" />
</p>


## Features

- Output string representation to `stdout`.
- Output sound representation to a `.wav` file via the `--output` flag.
- Supports short flags like `-s` (sound) and `-o` (output).
- Configure sound pitch `--pitch`

## Quickstart

To translate text to Morse code, run:

```console
morse 'Vamos' -s
```

This command will write the string representation to standard output and play it on your speakers.

To name the ouput file something different, use the `--file-name` (`-f`) flag. For example:

```console
morse 'Lets go, Carlos, lets go' -sf 'carlitos.wav'
```


## Configuring sounds

By default, the sound output is a 700hz sine wave. With a speed of 20 words per minute (WPM).

Morse supports configuring the pitch, and speed of sound using the `--pitch` and `--speed` flags.

### Pitch
You can set the pitch anywhere from 300hz to 1000hz via the `--pitch` 

```console
morse 'Ace' -s --pitch 500
```

### Speed
Morse code speed is measured in **words per minute** (wpm). Because characters might have different lengths, the convention is to meausre words per minute using the word "PARIS " with a space at the end.

`morse` produces sound outputs at 20wpm, try adjusting it to anything between 5-40 wpm.

```console
morse 'What a magnificent shot' -s --speed 35
```

### Tone
Set the output tone to one of `sine`, `triangle`, `sawtooth` and `square`. Via the `--tone` flag.

```console
morse 'And the champion of the 2025 Nitto ATP finals is Jannik Sinner' -s --tone triangle
```
- `sine`: The base, smooth. Makes up all other waves.
- `square`: Richer and buzzier.
- `triangle`: Between sine and square.
- `sawtooth`: More friction and the buzziest of all.


## Roadmap

- [x] Enable saving sound to file
- [x] Make functions run in parallel
- [x] Set up short-version flags
- [x] Extend punctuation
- [x] Add a default file name for better ergonomics
- [x] Enable editing sound qualities
  - [x] Speed
  - [x] Pitch
  - [x] Tone

- [ ] Improve efficiency of steream construction / duplication.
- [ ] Improve error handling in goroutine
