package main

import (
	"github.com/nsf/termbox-go"
)

func writeString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
