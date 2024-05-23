package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := [][]uint8{}
	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			pixel := uint8((i * j) / 2)
			row = append(row, pixel)
		}
		result = append(result, row)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
