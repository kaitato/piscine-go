package main

import (
	"fmt"
)

func main() {
	fmt.Println(MaxWordCountN(`
	Orange Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basetball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3))
}
