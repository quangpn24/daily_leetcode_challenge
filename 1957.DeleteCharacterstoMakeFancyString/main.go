package main

func makeFancyString(s string) string {
	arr := []byte{s[0]}

	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == arr[len(arr)-1] {
			cnt++
		} else {
			cnt = 1
		}

		if cnt != 3 {
			arr = append(arr, s[i])
		} else {
			cnt--
		}
	}
	return string(arr)
}
func main() {
	println(makeFancyString("aaabaaaa"))
}
