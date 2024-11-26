package main

import "fmt"

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 8}
	triangle := Triangle{a: 3, b: 4, c: 5}
	trash := struct {
		foo string
	}{foo: "bar"}

	circleArea, err := CalculateArea(&circle)
	if err != nil {
		fmt.Println("Error calculating circle area:", err)
	} else {
		fmt.Printf("Circle area: %.2f\n", circleArea)
	}

	rectangleArea, err := CalculateArea(&rectangle)
	if err != nil {
		fmt.Println("Error calculating rectangle area:", err)
	} else {
		fmt.Printf("Rectangle area: %.2f\n", rectangleArea)
	}

	triangleArea, err := CalculateArea(&triangle)
	if err != nil {
		fmt.Println("Error calculating triangle area:", err)
	} else {
		fmt.Printf("Triangle area: %.2f\n", triangleArea)
	}

	_, err = CalculateArea(&trash)
	if err != nil {
		fmt.Println("Error calculating area for unsupported shape:", err)
	}
}
