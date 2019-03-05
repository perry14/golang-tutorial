package main

import "github.com/golang/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic {
			pic[i][j] = uint8((j ^ i) * (10 ^ (i * j)))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
