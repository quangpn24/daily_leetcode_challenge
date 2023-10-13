package main

type MyHashMap struct {
    mMap map[int]int
}

func Constructor() MyHashMap {
    m := make(map[int]int)
    return MyHashMap{mMap: m}
}

func (this *MyHashMap) Put(key int, value int) {
    this.mMap[key] = value
}

func (this *MyHashMap) Get(key int) int {
    if v, ok := this.mMap[key]; ok {
        return v
    }
    return -1
}

func (this *MyHashMap) Remove(key int) {
    delete(this.mMap, key)
}
