package main

import "fmt"

func predictPartyVictory(senate string) string {
    radiants := []int{} // contains position (index) of 'R'
    dires := []int{}    // contains position (index) of 'D'
    n := len(senate)
    for i, s := range senate {
        if s == 'R' {
            radiants = append(radiants, i)
        } else {
            dires = append(dires, i)
        }
    }

    for len(radiants) > 0 && len(dires) > 0 {
        r := radiants[0]
        d := dires[0]
        if r < d {
            radiants = append(radiants, n+r)
        } else {
            dires = append(dires, n+d)
        }
        radiants = radiants[1:]
        dires = dires[1:]
    }
    if len(radiants) != 0 {
        return "Radiant"
    }
    return "Dire"
}
func main() {
    senate := "RDD"
    fmt.Println(predictPartyVictory(senate))
}
