package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math"

	"golang.org/x/image/math/f64"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	resourses "github.com/PhilipSolo1/DankDoodleJump/assets"
)

const (
	screenWidth  = 640
	screenHeight = 480

	tileSize     = 32
)

var (
	tilesImage *ebiten.Image
)

func init(){
	img, _, err := image.Decode(bytes.NewReader(resourses.Tiles_png))
	if err != nil {log.Fatal(err)}
	tilesImage = ebiten.NewImageFromImage(img)
}

