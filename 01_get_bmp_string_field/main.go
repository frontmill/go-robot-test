package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	var xPos1, xPos2, yPos1, yPos2 int

	ok := robotgo.AddEvents("1", "ctrl", "shift")
	if ok {
		xPos1, yPos1 = robotgo.GetMousePos()
		fmt.Println("x1 & y1 captured! Move to x2, y2")
	}

	ok = robotgo.AddEvents("2", "ctrl", "shift")
	if ok {
		xPos2, yPos2 = robotgo.GetMousePos()
		fmt.Println(xPos2, yPos2)
	}

	// start hook
	s := robotgo.Start()
	// end hook
	defer robotgo.End()

	for ev := range s {
		fmt.Println(ev)

		if noneEmpty(xPos1, yPos1, xPos2, yPos2){

			bmp := getBitmapField(xPos1, yPos1, xPos2, yPos2)

			fmt.Println(bmp)

			return
		}

	}

}

func noneEmpty(xInt ...int) bool {
	for _, i := range xInt {
		if i == 0 {
			return false
		}
	}
	return true
}

//Gets the bitmap-string from within a certain box (x1,y1,x2,y2)
func getBitmapField(x1, y1, x2, y2 int) string {

	outX1, outY1, w, h := coordsToCoordsWH(x1, y1, x2, y2)

	screen := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(screen)

	fmt.Println(screen)

	portion := robotgo.GetPortion(screen, outX1, outY1, w, h)
	defer robotgo.FreeBitmap(portion)

	bmpStr := robotgo.TostringBitmap(portion)

	return bmpStr
}

//converts x1,y1,x2,y2 to x1,y1,width,height
func coordsToCoordsWH(x1, y1, x2, y2 int) (int, int, int, int){

	var outX1, outY1, outX2, outY2 int

	//If x1 > x2: switch them around
	if x1 > x2 {
		outX1 = x2
		outX2 = x1
	} else {
		outX1 = x1
		outX2 = x2
	}

	//If y1 > y2: switch them around
	if y1 > y2 {
		outY1 = y2
		outY2 = y1
	} else {
		outY1 = y1
		outY2 = y2
	}

	xDiff := outX2 - outX1
	yDiff := outY2 - outY1

	return outX1, outY1, xDiff, yDiff

}
