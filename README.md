/*Fungsi yang membalikkan urutan  karakter*/
package main

import (
    
fmt
)

// Fungsi untuk membalik string
func balikString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    str := Halo
Dunia
    fmt.Println(Original:
, str)
    fmt.Println(
Terbalik:
, balikString(str))
} 
