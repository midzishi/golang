package main

import "fmt"

func main() {
	text := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	t := 0
	for i := 0; i < len(text); i++ {
		c := text[i] -'A'
		if t >= 6 {
			t = 0
		}
		c = (c - (keyword[t]-'A') + 26) % 26 + 'A'

		t ++
		fmt.Printf("%c", c)
	}
}
