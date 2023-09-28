package main

func findTheDifference(s string, t string) byte {
    mMap := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        mMap[s[i]]++
    }

    for i := 0; i < len(t); i++ {
        if v, ok := mMap[t[i]]; ok && v != 0 {
            mMap[t[i]]--
        } else {
            return t[i]
        }
    }
    return 0
}
func main() {

}
