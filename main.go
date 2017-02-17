package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"log"
	"math"
	"os"
	"time"
)

const (
	whRatio  = 2.3
	indicesL = 0.1 // indices length
	minutesL = 0.8 // minute hand length
	hoursL   = 0.6 // hour hand length
)

// converts polar coordinates r, ph to cartesian
// coordinates x, y with offset oX, oY
func polToCart(r, ph, oX, oY float64) (x, y float64) {
	x = r*math.Cos(ph)*whRatio + oX
	y = r*math.Sin(ph) + oY
	return
}

func drawPalette() {
	for col := 0; col < 8; col++ {
		for row := 0; row < 32; row++ {
			writeString(col*4, row, fmt.Sprintf("%d", col*32+row), termbox.Attribute(col*32+row), termbox.ColorBlack)
		}
	}
}

func drawClassicClock() {
	var x, y, x2, y2, ph float64

	maxX, maxY := termbox.Size()
	midX := float64(maxX / 2)
	midY := float64(maxY / 2)
	r := midY - 2
	rad := math.Pi * 2

	// draw clock face
	for ph := -0.25 * rad; ph < 0.75*rad; ph += rad / 12 {
		x, y = polToCart(r, ph, midX, midY)
		x2, y2 = polToCart((1-indicesL)*r, ph, midX, midY)
		drawLine(x, y, x2, y2, 1, 237, 10)
	}

	// draw clock hands
	t := time.Now()

	// Minutes
	min := float64(t.Minute())
	ph = (rad / 60 * min) - 0.25*rad
	x, y = polToCart(minutesL*r, ph, midX, midY)
	drawLine(x, y, midX, midY, 1, 248, 6)

	// Hours
	hrs := float64(t.Hour())
	ph = (rad / 12 * hrs) - 0.25*rad
	ph += (rad / 12) / 60 * min
	x, y = polToCart(hoursL*r, ph, midX, midY)
	drawLine(x, y, midX, midY, 1, 248, 6)

	// Seconds
	sec := float64(t.Second())
	ph = (rad / 60 * sec) - 0.25*rad
	x, y = polToCart(r, ph, midX, midY)
	drawLine(x, y, midX, midY, 1, 197, 4)

	// Date
	writeString(3, maxY-3, t.Format("Mon Jan 2 2006"), termbox.ColorWhite, termbox.ColorBlack)
	// colors
	//drawPalette()
}

func drawAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)
	drawClassicClock()
	termbox.Flush()
}

func main() {
	f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	drawAll()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			termbox.Interrupt()
		}
	}()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		case termbox.EventInterrupt:
			drawAll()
		case termbox.EventResize:
			drawAll()
		}
	}
}
