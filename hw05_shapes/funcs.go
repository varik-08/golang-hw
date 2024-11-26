package main

import "fmt"

func CalculateArea(s any) (float64, error) {
	switch s := s.(type) {
	case Shape:
		return s.Area(), nil
	default:
		return 0, fmt.Errorf("unsupported shape type: %T", s)
	}
}
