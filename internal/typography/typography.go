package typography

import (
	"app/resources/fonts"
	"log"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	once sync.Once

	CocolaFont font.Face
	fontsMap   map[string]*sfnt.Font
)

func init() {
	once.Do(func() {
		fontsMap = make(map[string]*sfnt.Font, 1)

		parsedCocola, err := opentype.Parse(fonts.Cocola_ttf)
		if err != nil {
			log.Fatal(err)
		}
		fontsMap["Cocola"] = parsedCocola
	})
}

func GetFront(fName string, size float64) (font.Face, error) {
	const dpi = 72

	return opentype.NewFace(fontsMap[fName], &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
