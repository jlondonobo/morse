package sound

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
)

const (
	WordsPerMinute = 20
	SampleRate     = beep.SampleRate(48000)
)

// Todo: Would be a lot more efficiento to seek Streamer to 0 instead of recreating it.
func generate(s string, pitch uint16) beep.Streamer {
	sine, err := generators.SineTone(SampleRate, float64(pitch))
	if err != nil {
		log.Fatal("Error generating sine tone")
	}
	silence := generators.Silence(-1)
	// based on: https://morsecode.world/international/timing/
	spd := time.Duration(60 * 1000 / (50 * WordsPerMinute)) // in milliseconds
	dit := SampleRate.N(time.Millisecond * spd)
	dah := dit * 3

	// todo: this is not exact correspondence. Just for convenience assuming
	// 1-dit silence after every symbol.
	var m = map[string]func() beep.Streamer{
		".": func() beep.Streamer { return beep.Take(dit, sine) },
		"-": func() beep.Streamer { return beep.Take(dah, sine) },
		" ": func() beep.Streamer { return beep.Take(dit*2, silence) },
		"/": func() beep.Streamer { return beep.Take(0, silence) },
	}
	var sounds []beep.Streamer

	for _, v := range s {
		st, ok := m[string(v)]
		if !ok {
			log.Fatal("Unrecognize symbol.")
		}
		sounds = append(sounds, st(), beep.Take(dit, silence))
	}
	return beep.Seq(sounds...)
}

func Play(s string, pitch uint16) {
	speaker.Init(SampleRate, 4800)

	ch := make(chan struct{})
	seq := generate(s, pitch)

	sounds := beep.Seq(seq, beep.Callback(func() { ch <- struct{}{} }))
	speaker.Play(sounds)
	<-ch
	time.Sleep(200 * time.Millisecond) // to ensure last signal plays
}

func Write(s string, name string, pitch uint16) {
	finalStreamer := generate(s, pitch)
	outFile, err := os.Create(name)
	if err != nil {
		log.Fatal("Unable to create file.")
	}
	defer outFile.Close()
	fmt := beep.Format{SampleRate: 48000, NumChannels: 2, Precision: 2}
	err = wav.Encode(outFile, finalStreamer, fmt)
	if err != nil {
		log.Fatal(err)
	}
}
