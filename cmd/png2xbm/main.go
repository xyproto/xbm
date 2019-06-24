package main

import (
	"github.com/xyproto/xbm"
	"image/png"
	"os"
)

func main() {
	inputFileName := "input.png"
	outputFileName := "output.xbm"

	f, err := os.Open(inputFileName)
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = xbm.Encode(f, m)
	if err != nil {
		panic(err)
	}
}
