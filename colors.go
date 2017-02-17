package main

import (
	"github.com/nsf/termbox-go"
	//	"log"
)

/*
func colorToChar(c float64) rune {
	if c < 0.35 {
		return '░'
	}
	if c < 0.6 {
		return '▒'
	}
	if c < 0.85 {
		return '▓'
	}
	if c <= 1 {
		return '█'
	}
	return '-'
}
*/

func colorToChar(c float64) rune {
	if c < 0.35 {
		return '·'
	}
	if c < 0.6 {
		return '*'
	}
	if c < 0.85 {
		return '⁙'
	}
	if c <= 1 {
		return '※'
	}
	return '-'
}
func shadeColor(c float64, off, num int) termbox.Attribute {
	if c == 0 {
		return termbox.Attribute(0)
	}
	ret := round(float64(num-1)*c) + float64(off)
	//log.Printf("shadeColor(): %v, %v, %v: %v\n", c, off, num, ret)
	return termbox.Attribute(ret)
}
