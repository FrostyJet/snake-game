package audio

import (
	_ "embed"
)

var (
	//go:embed dull-click.ogg
	DullClick_ogg []byte

	//go:embed bleep.ogg
	Bleep_ogg []byte

	//go:embed snake-self-bite.ogg
	SnakeSelfBite_ogg []byte

	//go:embed game-start.ogg
	GameStart_ogg []byte
)
