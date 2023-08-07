package main

import (
	"fmt"
	"math"
)

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	curr := len(asteroids) - 1
	for i := len(asteroids) - 1; i >= 0; i-- {
		if asteroids[i] > 0 && len(res) > 0 && res[len(res)-1] < 0 {
			asteroids[i+1] = res[len(res)-1]
			res = res[:len(res)-1]
			i++
			curr++
		}
		if i == 0 {
			break
		}
		if asteroids[i] < 0 && asteroids[i-1] > 0 {
			if math.Abs(float64(asteroids[i])) > math.Abs(float64(asteroids[i-1])) {
				asteroids[i-1] = asteroids[i]
			} else if math.Abs(float64(asteroids[i])) == math.Abs(float64(asteroids[i-1])) {
				i--
				curr--
			} else if len(res) > 0 {
				asteroids[i] = res[len(res)-1]
				res = res[:len(res)-1]
				i++
				curr++
			}
		} else {
			res = append(res, asteroids[i])
		}
		curr--
	}
	if curr == 0 {
		res = append(res, asteroids[0])
	}
	//revert arr
	l := 0
	r := len(res) - 1
	for l < r {
		tmp := res[r]
		res[r] = res[l]
		res[l] = tmp
		r--
		l++
	}
	return res
}
func main() {
	asteroids := []int{-3, -8, -4, 5, -2, 8, 10, -10, -7, 4}
	fmt.Println(asteroidCollision(asteroids))
}
