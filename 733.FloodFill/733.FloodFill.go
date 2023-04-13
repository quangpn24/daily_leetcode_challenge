package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	preColor := image[sr][sc]
	if image[sr][sc] == color {
		return image
	}
	image[sr][sc] = color
	if sr+1 < len(image) && image[sr+1][sc] == preColor {
		floodFill(image, sr+1, sc, color)
	}
	if sr-1 >= 0 && image[sr-1][sc] == preColor {
		floodFill(image, sr-1, sc, color)
	}
	if sc+1 < len(image[sr]) && image[sr][sc+1] == preColor {
		floodFill(image, sr, sc+1, color)
	}
	if sc-1 >= 0 && image[sr][sc-1] == preColor {
		floodFill(image, sr, sc-1, color)
	}
	return image
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
}
