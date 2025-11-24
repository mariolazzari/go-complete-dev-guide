package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	colors2 := make(map[string]string)
	colors2["black"] = "#000000"
	colors2["yellow"] = "#ffff00"
	fmt.Println(colors2)

	delete(colors2, "black")

	for k, v := range colors {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for c, hex := range c {
		fmt.Printf("Color: %s Hex: %s\n", c, hex)
	}
}
