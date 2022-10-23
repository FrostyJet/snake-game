package sounds

import (
	"bytes"
	"log"
	"sync"

	audioResources "app/resources/audio"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var (
	audioContext        *audio.Context
	GameStartPlayer     *audio.Player
	BleepPlayer         *audio.Player
	SnakeSelfBitePlayer *audio.Player

	once sync.Once
)

func init() {
	once.Do(func() {
		audioContext = audio.NewContext(48_000)

		// GameStart_ogg
		{
			gameStart, err := vorbis.DecodeWithoutResampling(bytes.NewReader(audioResources.GameStart_ogg))
			if err != nil {
				log.Fatal(err)
			}

			GameStartPlayer, err = audioContext.NewPlayer(gameStart)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Bleep_ogg
		{
			bleep, err := vorbis.DecodeWithoutResampling(bytes.NewReader(audioResources.Bleep_ogg))
			if err != nil {
				log.Fatal(err)
			}
			BleepPlayer, err = audioContext.NewPlayer(bleep)
			if err != nil {
				log.Fatal(err)
			}
			BleepPlayer.SetVolume(0.5)
		}

		// SnakeSelfBite_ogg
		{
			snakeSelfBite, err := vorbis.DecodeWithoutResampling(bytes.NewReader(audioResources.SnakeSelfBite_ogg))
			if err != nil {
				log.Fatal(err)
			}

			SnakeSelfBitePlayer, err = audioContext.NewPlayer(snakeSelfBite)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}
