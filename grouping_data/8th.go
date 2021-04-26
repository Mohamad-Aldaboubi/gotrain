package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolote", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
