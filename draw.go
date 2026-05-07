package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const DIAMOND_SIZE = 10
const DIAMOND_OFFSET = 3
const DIGIT_OFFSET = 34
const DOT_SIZE = 10

var DIAMOND_CLR_ON = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var DIAMOND_CLR_OFF = color.RGBA{R: 20, G: 0, B: 0, A: 255}

func DrawClock(screen *ebiten.Image, win_w, win_h float32, bluePrint [6][7]bool) {
	var center = win_w / 2
	space := float32(DIAMOND_SIZE*2 + DIGIT_OFFSET)

	drawDigit(screen, center-(space*6), win_h/2, bluePrint[0])
	drawDigit(screen, center-(space*4), win_h/2, bluePrint[1])

	vector.FillCircle(screen, center-(space*2.5), win_h/2+DIGIT_OFFSET, DOT_SIZE, DIAMOND_CLR_ON, false)
	vector.FillCircle(screen, center-(space*2.5), win_h/2-DIGIT_OFFSET, DOT_SIZE, DIAMOND_CLR_ON, false)

	drawDigit(screen, center-space, win_h/2, bluePrint[2])
	drawDigit(screen, center+space, win_h/2, bluePrint[3])

	vector.FillCircle(screen, center+(space*2.5), win_h/2+DIGIT_OFFSET, DOT_SIZE, DIAMOND_CLR_ON, false)
	vector.FillCircle(screen, center+(space*2.5), win_h/2-DIGIT_OFFSET, DOT_SIZE, DIAMOND_CLR_ON, false)

	drawDigit(screen, center+(space*4), win_h/2, bluePrint[4])
	drawDigit(screen, center+(space*6), win_h/2, bluePrint[5])
}

func drawDigit(screen *ebiten.Image, pos_x, pos_y float32, bluePrint [7]bool) {

	// 1
	drawDiamond(screen, pos_x, pos_y-(DIAMOND_SIZE*6+DIAMOND_OFFSET*2), DIAMOND_SIZE, false, setDiamondColor(bluePrint[0]))

	// 2
	drawDiamond(screen, pos_x-(DIAMOND_SIZE*3+DIAMOND_OFFSET), pos_y-(DIAMOND_SIZE*3+DIAMOND_OFFSET), DIAMOND_SIZE, true, setDiamondColor(bluePrint[1]))

	// 3
	drawDiamond(screen, pos_x+(DIAMOND_SIZE*3+DIAMOND_OFFSET), pos_y-(DIAMOND_SIZE*3+DIAMOND_OFFSET), DIAMOND_SIZE, true, setDiamondColor(bluePrint[2]))

	// 4
	drawDiamond(screen, pos_x, pos_y, DIAMOND_SIZE, false, setDiamondColor(bluePrint[3]))

	// 5
	drawDiamond(screen, pos_x-(DIAMOND_SIZE*3+DIAMOND_OFFSET), pos_y+(DIAMOND_SIZE*3+DIAMOND_OFFSET), DIAMOND_SIZE, true, setDiamondColor(bluePrint[4]))

	// 6
	drawDiamond(screen, pos_x+(DIAMOND_SIZE*3+DIAMOND_OFFSET), pos_y+(DIAMOND_SIZE*3+DIAMOND_OFFSET), DIAMOND_SIZE, true, setDiamondColor(bluePrint[5]))

	// 7
	drawDiamond(screen, pos_x, pos_y+(DIAMOND_SIZE*6+DIAMOND_OFFSET*2), DIAMOND_SIZE, false, setDiamondColor(bluePrint[6]))

}

func drawDiamond(screen *ebiten.Image, pos_x, pos_y, size float32, vertical bool, clr color.RGBA) {
	var path vector.Path

	path.MoveTo(pos_x, pos_y)

	if vertical {
		path.LineTo(pos_x, pos_y+(size*3))
		path.LineTo(pos_x+size, pos_y+(size*2))
		path.LineTo(pos_x+size, pos_y-(size*2))
		path.LineTo(pos_x, pos_y-(size*3))
		path.LineTo(pos_x-size, pos_y-(size*2))
		path.LineTo(pos_x-size, pos_y+(size*2))
		path.LineTo(pos_x, pos_y+(size*3))
	} else {
		path.LineTo(pos_x+(size*3), pos_y)
		path.LineTo(pos_x+(size*2), pos_y+size)
		path.LineTo(pos_x-(size*2), pos_y+size)
		path.LineTo(pos_x-(size*3), pos_y)
		path.LineTo(pos_x-(size*2), pos_y-size)
		path.LineTo(pos_x+(size*2), pos_y-size)
		path.LineTo(pos_x+(size*3), pos_y)
	}

	path.Close()

	vector.FillPath(screen, &path, &vector.FillOptions{}, &vector.DrawPathOptions{ColorScale: setColorScale(clr)})

}

func setColorScale(clr color.RGBA) ebiten.ColorScale {
	var cs ebiten.ColorScale
	cs.SetR(float32(clr.R) / 255)
	cs.SetG(float32(clr.G) / 255)
	cs.SetB(float32(clr.B) / 255)
	cs.SetA(float32(clr.A) / 255)
	return cs
}

func setDiamondColor(cond bool) color.RGBA {
	if cond {
		return DIAMOND_CLR_ON
	}

	return DIAMOND_CLR_OFF
}
