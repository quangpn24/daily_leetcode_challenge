package main

func maximumWealth(accounts [][]int) int {
    max := 0
    m, n := len(accounts), len(accounts[0])
    for i := 0; i < m; i++ {
        sum := 0
        for j := 0; j < n; j++ {
            sum += accounts[i][j]
        }
        if max < sum {
            max = sum
        }
    }
    return max
}
func main() {

}
