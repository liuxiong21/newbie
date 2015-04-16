package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit.\n"

func printPrompt() {
	if runtime.GOOS == "windows" {
		fmt.Printf(prompt, "Ctrl-Z,Entry")
	} else {
		fmt.Printf(prompt, "Ctrl-D")
	}
}

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

func Polar2cartesian() {
	printPrompt()
	questions := make(chan polar, 10)
	defer close(questions)
	answers := createResover(questions)
	defer close(answers)

	interact(questions, answers)
}

func createResover(questions chan polar) chan cartesian {
	answers := make(chan cartesian, 10)
	go func() {
		for {
			polarCoord := <-questions
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Radius and angle: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		var radius, θ float64
		fmt.Sscanf(input, "%f %f", &radius, &θ)
		questions <- polar{radius, θ}
		answer := <-answers
		fmt.Printf("Radius=%.02f,θ=%.02f\n", answer.x, answer.y)
	}

}
