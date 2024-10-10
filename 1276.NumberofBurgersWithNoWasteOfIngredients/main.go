package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// x is the number of Jumbo Burger,
	// y is the number of Small Burger
	//=>  4x + 2y = tomatoSlices && x + y = cheeseSlices <=> 2x + y = tomatoSlices / 2 && x + y = cheeseSlices
	// => x = tomatoSlices / 2 - cheeseSlices
	// => y = cheeseSlices - x

	if tomatoSlices%2 != 0 {
		return []int{}
	}

	x := tomatoSlices/2 - cheeseSlices
	y := cheeseSlices - x
	if x < 0 || y < 0 {
		return []int{}
	}

	return []int{x, y}
}

func main() {
	arr := numOfBurgers(16, 7)
	for _, v := range arr {
		println(v)
	}
}
