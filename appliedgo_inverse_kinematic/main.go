package main

import (
	"fmt"
	"math"
)

const (
	len1 = 10.0
	len2 = 10.0
)

func lawOfCosines(a, b, c float64) (C float64) {
	return math.Acos((a*a + b*b - c*c) / (2 * a * b))
}

func distance(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func angles(x, y float64) (A1, A2 float64) {
	dist := distance(x, y)
	D1 := math.Atan2(y, x)
	D2 := lawOfCosines(dist, len1, len2)
	A1 = D1 + D2
	A2 = lawOfCosines(len1, len2, dist)

	return A1, A2
}

func deg(rad float64) float64 {
	return rad * 180 / math.Pi
}

func main() {
	fmt.Println("Lets do some tests. First move to (5,5):")
	x, y := 5.0, 5.0
	a1, a2 := angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))
	fmt.Println("if y is 0 and x = sqrt(10^2 + 10^2) then alpha should become 45 degrees and beta should become 90 degrees.")
	x, y = math.Sqrt(200), 0
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))

	fmt.Println("Now let's try moving to (1, 19).")
	x, y = 1, 19
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))

	fmt.Println("n extreme case: (20, 0). The arm needs to stretch along the y axis.")
	x, y = 20, 0
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))

	fmt.Println("And(0,20).")
	x, y = 0, 20
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))

	fmt.Println("Moving to (0,0) technically works if the arm segments have the same length, and if the arm does not block itself. Still the result looks a bit weird!?")
	x, y = 0, 0
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))

	fmt.Println("What happens if the target point is outside the reach? Like (20,20).")
	x, y = 20, 20
	a1, a2 = angles(x, y)
	fmt.Printf("x=%v, y=%v: A1=%v (%v degrees), A2=%v (%v degrees)\n", x, y, a1, deg(a1), a2, deg(a2))
}
