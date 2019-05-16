package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main()  {
	out := robotgo.OpenBitmap("test.png")
	defer robotgo.FreeBitmap(out)
	fmt.Println(out)

	robotgo.BitmapClick(out)

}


/*func clickBitmap(bmp robotgo.CBitmap) {
	fx, fy := robotgo.FindCBitmap(bmp)
}*/