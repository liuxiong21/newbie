package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"strings"
	"path/filepath"
	"image/jpeg"
	"image/png"
	"runtime"
	"os"
)

var saneLength, saneRadius, saneSides func(int) int

func init() {
	saneLength = makeBoundedIntFunc(1, 4096)
	saneRadius = makeBoundedIntFunc(1, 1024)
	saneSides = makeBoundedIntFunc(3, 60)
}


func makeBoundedIntFunc(min, max int) func(int) int {
	return func(x int) int {
		valid := x
		if x < min {
			valid = min
		} else if x > max {
			valid = max
		}
		if valid != x {
			log.Printf("%s(): replaced %d with %d\n", caller(1), x, valid)
		}
		return valid
	}
}

func caller(steps int) string {
	name := "?"
	if pc, _, _, ok := runtime.Caller(steps); ok {
		name = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return name
}

type Shaper interface {
	Fill() color.Color
	SetFill(fill color.Color)
	Draw(img draw.Image, x, y int) error
}

type CircularShaper interface {
	Shaper
	Radius() int
	SetRadius(radius int)
}

type RegularPolygonalShaper interface {
	CircularShaper
	Sides() int
	SetSides(side int)
}

type Shape struct {
	fill color.Color
}

func NewShape(fill color.Color) *Shape {
	if fill == nil {
		fill = color.Black
	}
	return &Shape{fill}
}

func (shape Shape) Fill() color.Color {
	return shape.fill
}

func (shape *Shape) SetFill(fill color.Color) {
	if fill == nil {
		fill = color.Black
	}
	shape.fill = fill
}

type Circle struct {
	Shape
	radius int
}

func (circle Circle) Radius() int {
	return circle.radius
}

func (circle *Circle) SetRadius(radius int) {
	circle.radius = saneRadius(radius)
}

func NewCircle(fill color.Color, radius int) *Circle {
	return &Circle{*NewShape(fill), saneRadius(radius)}
}

func (circle *Circle) Draw(img draw.Image, x, y int) error {
	if err := checkBounds(img, x, y); err != nil {
		return err
	}
	fill, radius := circle.fill, circle.radius
	x0, y0 := x, y
	f := 1 - radius
	ddF_x, ddF_y := 1, -2*radius
	x, y = 0, radius

	img.Set(x0, y0+radius, fill)
	img.Set(x0, y0-radius, fill)
	img.Set(x0+radius, y0, fill)
	img.Set(x0-radius, y0, fill)

	for x < y {
		if f >= 0 {
			y--
			ddF_y += 2
			f += ddF_y
		}
		x++
		ddF_x += 2
		f += ddF_x
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
		img.Set(x0+y, y0+x, fill)
		img.Set(x0-y, y0+x, fill)
		img.Set(x0+y, y0-x, fill)
		img.Set(x0-y, y0-x, fill)
	}
	return nil
}

func (circle *Circle) String() string {
	return fmt.Sprintf("circle(fill=%v, radius=%d)", circle.fill,
		circle.radius)
}

func checkBounds(img image.Image, x, y int) error {
	if !image.Rect(x, y, x, y).In(img.Bounds()) {
		return fmt.Errorf("%s(): point (%d, %d) is outside the image\n",
			caller(1), x, y)
	}
	return nil
}

type RegularPolygon struct {
	*Circle
	sides int
}

func NewRegularPolygon(fill color.Color, radius,
	sides int) *RegularPolygon {
	return &RegularPolygon{NewCircle(fill, radius), saneSides(sides)}
}

func (polygon *RegularPolygon) Sides() int {
	return polygon.sides
}

func (polygon *RegularPolygon) SetSides(sides int) {
	polygon.sides = saneSides(sides)
}

func (polygon *RegularPolygon) Draw(img draw.Image, x, y int) error {

	if err := checkBounds(img, x, y); err != nil {
		return err
	}
	points := getPoints(x, y, polygon.sides, float64(polygon.Radius()))
	for i := 0; i < polygon.sides; i++ { // Draw lines between the apexes
		drawLine(img, points[i], points[i+1], polygon.Fill())
	}
	return nil
}

func getPoints(x, y, sides int, radius float64) []image.Point {
	points := make([]image.Point, sides+1)
	// Compute the shape's apexes (thanks to Jasmin Blanchette)
	fullCircle := 2 * math.Pi
	x0, y0 := float64(x), float64(y)
	for i := 0; i < sides; i++ {
		θ := float64(float64(i) * fullCircle / float64(sides))
		x1 := x0 + (radius * math.Sin(θ))
		y1 := y0 + (radius * math.Cos(θ))
		points[i] = image.Pt(int(x1), int(y1))
	}
	points[sides] = points[0] // close the shape
	return points
}

func drawLine(img draw.Image, start, end image.Point,
	fill color.Color) {
	x0, x1 := start.X, end.X
	y0, y1 := start.Y, end.Y
	Δx := math.Abs(float64(x1 - x0))
	Δy := math.Abs(float64(y1 - y0))
	if Δx >= Δy { // shallow slope
		if x0 > x1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		y := y0
		yStep := 1
		if y0 > y1 {
			yStep = -1
		}
		remainder := float64(int(Δx/2)) - Δx
		for x := x0; x <= x1; x++ {
			img.Set(x, y, fill)
			remainder += Δy
			if remainder >= 0.0 {
				remainder -= Δx
				y += yStep
			}
		}
	} else { // steep slope
		if y0 > y1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		x := x0
		xStep := 1
		if x0 > x1 {
			xStep = -1
		}
		remainder := float64(int(Δy/2)) - Δy
		for y := y0; y <= y1; y++ {
			img.Set(x, y, fill)
			remainder += Δx
			if remainder >= 0.0 {
				remainder -= Δy
				x += xStep
			}
		}
	}
}

func (polygon *RegularPolygon) String() string {
	return fmt.Sprintf("polygon(fill=%v, radius=%d, sides=%d)",
		polygon.Fill(), polygon.Radius(), polygon.sides)
}

type Option struct {
	Fill   color.Color
	Radius int
}

func New(shape string, option Option) (Shaper, error) {
	sidesForShape := map[string]int{"triangle": 3, "square": 4,
		"pentagon": 5, "hexagon": 6, "heptagon": 7, "octagon": 8,
		"enneagon": 9, "nonagon": 9, "decagon": 10}
	if sides, found := sidesForShape[shape]; found {
		return NewRegularPolygon(option.Fill, option.Radius, sides), nil
	}
	if shape != "circle" {
		return nil, fmt.Errorf("shapes.New(): invalid shape '%s'", shape)
	}
	return NewCircle(option.Fill, option.Radius), nil
}

func FilledImage(width, height int, fill color.Color) draw.Image {
	if fill == nil { // We silently treat a nil color as black
		fill = color.Black
	}
	width = saneLength(width)
	height = saneLength(height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{fill}, image.ZP, draw.Src)
	return img
}

func DrawShapes(img draw.Image, x, y int, shapes ...Shaper) error {
	for _, shape := range shapes {
		if err := shape.Draw(img, x, y); err != nil {
			return err
		}
		// Thicker so that it shows up better in screenshots
		if err := shape.Draw(img, x+1, y); err != nil {
			return err
		}
		if err := shape.Draw(img, x, y+1); err != nil {
			return err
		}
	}
	return nil
}

func sanityCheck(name string, shape Shaper) {
	fmt.Print("name=", name, " ")
	fmt.Print("fill=", shape.Fill(), " ")
	if shape, ok := shape.(CircularShaper); ok {
		fmt.Print("radius=", shape.Radius(), " ")
		if shape, ok := shape.(RegularPolygonalShaper); ok {
			fmt.Print("sides=", shape.Sides(), " ")
		}
	}
	fmt.Println()
}

func showShapeDetails(shape Shaper) {
	fmt.Print("fill=", shape.Fill(), " ")        // All shapes have a fill color
	if shape, ok := shape.(CircularShaper); ok { // shadow variable
		fmt.Print("radius=", shape.Radius(), " ")
		if shape, ok := shape.(RegularPolygonalShaper); ok { //shadow
			fmt.Print("sides=", shape.Sides(), " ")
		}
	}
	fmt.Println()
}

func printShape1() {
	log.SetFlags(0)
	const width, height = 400, 200
	img := FilledImage(width, height,
		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	x, y := width/4, height/2

	red := color.RGBA{0xFF, 0, 0, 0xFF}
	blue := color.RGBA{0, 0, 0xFF, 0xFF}
	// Purely for testing New() vs. New*()
	if len(os.Args) == 1 {
		fmt.Println("Using NewCircle() & NewRegularPolygon()")
		circle := NewCircle(blue, 90)
		circle.SetFill(red) // Uses the aggregated shape.SetFill method
		octagon := NewRegularPolygon(red, 75, 8)
		octagon.SetFill(blue) // Uses the aggregated circle.shape.SetFill
		polygon := NewRegularPolygon(image.Black, 65, 4)
		if err := DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheck("circle", circle)
		sanityCheck("octagon", octagon)
		sanityCheck("polygon", polygon)
	} else {
		fmt.Println("Using New()")
		
		if _, err := New("Misshapen", Option{blue, 5}); err == nil {
			fmt.Println("unexpectedly got a non-nil invalid shape!")
		}
		circle, _ := New("circle", Option{blue, 5})
		circle.SetFill(red)
		circle.(CircularShaper).SetRadius(90)
		octagon, _ := New("octagon", Option{red, 10})
		octagon.SetFill(blue)
		
		if octagon, ok := octagon.(RegularPolygonalShaper); ok {
			octagon.SetRadius(75)
		}
		polygon, _ := New("square", Option{Radius: 65})
		if err := DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheck("circle", circle)
		sanityCheck("octagon", octagon)
		sanityCheck("polygon", polygon)
	}
	polygon := NewRegularPolygon(color.RGBA{0, 0x7F, 0, 0xFF}, 65, 4)
	showShapeDetails(polygon)
	y = 30
	for i, radius := range []int{60, 55, 50, 45, 40} {
		polygon.SetRadius(radius)
		polygon.SetSides(i + 5)
		x += radius
		y += height / 8
		if err := DrawShapes(img, x, y, polygon); err != nil {
			fmt.Println(err)
		}
	}

	filename := "~/Downloads/shapes.png"
	if err := SaveImage(img, filename); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Saved", filename)
	}
	fmt.Println("OK")
	img = FilledImage(width, height, image.White)
	x, y = width/3, height/4
}

func SaveImage(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    switch strings.ToLower(filepath.Ext(filename)) {
    case ".jpg", ".jpeg":
        return jpeg.Encode(file, img, nil)
    case ".png":
        return png.Encode(file, img)
    }
    return fmt.Errorf("shapes.SaveImage(): '%s' has an unrecognized "+
        "suffix", filename)
}
