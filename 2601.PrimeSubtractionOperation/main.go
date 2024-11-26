package main

func primeSubOperation(nums []int) bool {
	primes := genPrimes(1000)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			continue
		}

		for j := 2; j < len(primes); j++ {
			if nums[i]-nums[i+1] < primes[j] {
				nums[i] -= primes[j]
				break
			}
		}
		if nums[i] <= 0 || nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func genPrimes(maxNum int) []int {
	primesCheck := make([]bool, maxNum+1)

	for i := range primesCheck {
		primesCheck[i] = true
	}

	for i := 2; i <= maxNum; i++ {
		for j := i * i; j <= maxNum; j += i {
			primesCheck[j] = false
		}
	}

	primes := make([]int, 0)
	for i, b := range primesCheck {
		if b {
			primes = append(primes, i)
		}
	}

	return primes
}
func main() {
	println(primeSubOperation([]int{18, 12, 14, 6}))
}
