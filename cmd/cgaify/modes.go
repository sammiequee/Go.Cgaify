package main

//go:generate go run gen\makevga.go

import (
	"image/color"
)

type mode struct {
	desc   string
	width  uint
	height uint
	colors color.Palette
}

func (m *mode) aspectRatio() float64 {
	return float64(m.width) / float64(m.height)
}

var egacolors = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // 0
	color.RGBA{0x00, 0x00, 0xAA, 0xff}, // 1
	color.RGBA{0x00, 0xAA, 0x00, 0xff}, // 2
	color.RGBA{0x00, 0xAA, 0xAA, 0xff}, // 3
	color.RGBA{0xAA, 0x00, 0x00, 0xff}, // 4
	color.RGBA{0xAA, 0x00, 0xAA, 0xff}, // 5
	color.RGBA{0xAA, 0x55, 0x00, 0xff}, // 6
	color.RGBA{0xAA, 0xAA, 0xAA, 0xff}, // 7
	color.RGBA{0x55, 0x55, 0x55, 0xff}, // 8
	color.RGBA{0x55, 0x55, 0xFF, 0xff}, // 9
	color.RGBA{0x55, 0xFF, 0x55, 0xff}, // 10
	color.RGBA{0x55, 0xFF, 0xFF, 0xff}, // 11
	color.RGBA{0xFF, 0x55, 0x55, 0xff}, // 12
	color.RGBA{0xFF, 0x55, 0xFF, 0xff}, // 13
	color.RGBA{0xFF, 0xFF, 0x55, 0xff}, // 14
	color.RGBA{0xFF, 0xFF, 0xFF, 0xff}, // 15
}

// define the modes/palettes we support
var modes = map[string]*mode{
	"CGA1": {
		"CGA 320x200 Palette 1, Low Intensity, Cyan Magenta White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0xAA, 0xAA, 0xff},
			color.RGBA{0xAA, 0x00, 0xAA, 0xff},
			color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
		},
	},
	"CGA1H": {
		"CGA 320x200 Palette 1, High Intensity, Cyan Magenta White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x55, 0xFF, 0xFF, 0xff},
			color.RGBA{0xFF, 0x55, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"CGA0": {
		"CGA 320x200 Palette 0, Low Intensity, Green Red Brown",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0xAA, 0x00, 0xff},
			color.RGBA{0xAA, 0x00, 0x00, 0xff},
			color.RGBA{0xAA, 0x55, 0x00, 0xff},
		},
	},
	"CGA0H": {
		"CGA 320x200 Palette 0, High Intensity, Green Red Yellow",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x55, 0xFF, 0x55, 0xff},
			color.RGBA{0xFF, 0x55, 0x55, 0xff},
			color.RGBA{0xFF, 0xFF, 0x55, 0xff},
		},
	},
	"EGA": {
		"EGA 320x200 16-colors",
		320,
		200,
		egacolors,
	},
	"EGA640": {
		"EGA 640x350 16-colors",
		640,
		350,
		egacolors,
	},
	"VGA": {
		"VGA 320x200 256-colors",
		320,
		200,
		vgacolors,
	},
	"CLAY": {
		"CGA 320x200 Palette 1, High Intensity, Cyan Magenta White (Alternate)",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0xFF, 0xFF, 0xff},
			color.RGBA{0xFF, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"FLAY": {
		"Orange Magenta White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xFF, 0x5A, 0x00, 0xff},
			color.RGBA{0xFF, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"ZANE": {
		"Indigo Purple Magenta Pink LightBlue White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x80, 0x00, 0xFF, 0xff},
			color.RGBA{0xbb, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0x80, 0xFF, 0xff},
			color.RGBA{0xBB, 0xDD, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"UMBRA": {
		"Black DarkGrey Magenta White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x35, 0x35, 0x35, 0xff},
			color.RGBA{0xCC, 0x00, 0xFB, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"MAGENTA": {
		"Purple Magenta Pink",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xbb, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0x80, 0xFF, 0xff},
		},
	},
	"TRANS": {
		"Transgender flag colours",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xF5, 0xA9, 0xB8, 0xff},
			color.RGBA{0x5B, 0xCE, 0xFA, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"LESBIAN": {
		"Lesbian flag colours",
		320,
		200,
		color.Palette{
			color.RGBA{0xA3, 0x02, 0x62, 0xff},
			color.RGBA{0xD5, 0x2D, 0x00, 0xff},
			color.RGBA{0xD3, 0x62, 0xA4, 0xff},
			color.RGBA{0xFF, 0x9A, 0x56, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"BW": {
		"Black and White",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"BWR": {
		"Black White Red",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xFF, 0x00, 0x00, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"BWB": {
		"Black White Blue",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"RBW": {
		"Red Blue White",
		320,
		200,
		color.Palette{
			color.RGBA{0xFF, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"BWRB": {
		"Black White Red Blue",
		320,
		200,
		color.Palette{
			color.RGBA{0x00, 0x00, 0x00, 0xff},
			color.RGBA{0xFF, 0x00, 0x00, 0xff},
			color.RGBA{0x00, 0x00, 0xFF, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
	"8088": {
		"16-color palette used for 8088+CGA, 160x200",
		160,
		200,
		color.Palette{
			color.RGBA{0x09, 0x09, 0x09, 0xff},
			color.RGBA{0x00, 0x9c, 0x3a, 0xff},
			color.RGBA{0x40, 0x18, 0xff, 0xff},
			color.RGBA{0x00, 0xc5, 0xff, 0xff},
			color.RGBA{0xe3, 0x05, 0x4b, 0xff},
			color.RGBA{0x9a, 0x9d, 0x96, 0xff},
			color.RGBA{0xff, 0x1d, 0xff, 0xff},
			color.RGBA{0xdd, 0xc5, 0xff, 0xff},
			color.RGBA{0x4c, 0x74, 0x00, 0xff},
			color.RGBA{0x07, 0xff, 0x00, 0xff},
			color.RGBA{0x9e, 0x9a, 0x9a, 0xff},
			color.RGBA{0x53, 0xff, 0xd2, 0xff},
			color.RGBA{0xff, 0x78, 0x00, 0xff},
			color.RGBA{0xf5, 0xff, 0x00, 0xff},
			color.RGBA{0xff, 0x98, 0xfa, 0xff},
			color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
		},
	},
}
