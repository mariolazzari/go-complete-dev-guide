# Go: The Complete Developer's Guide

## Intro

## A simple start

### Hello world

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}
```

### Five important questions

```sh
go run main.go
```

### Go packages

- executable: package main
- reusable: package name

### Import statement

[Go package](https://pkg.go.dev/)
[fmt](https://pkg.go.dev/fmt)

## Deeper into Go

### Project overview

### New project folder

```sh
mkdir cards
cd cards
touch main.go
```

### Variables declaration

```go
package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of diamonds"
	fmt.Println(card)
}
```

### Function return type

```go
package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Hearts"
}
```

### Slices and loops

- Array: fixed size
- Slice: can grow or shrink

```go
package main

import "fmt"

func main() {
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades", "Seven of Clubs")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Hearts"
}
```

### OOP vs Go

- custom types
- receiver function

### Custom types

```go
package main

func main() {
	d := deck{"Ace of diamonds", newCard()}
	d = append(d, "Six of spades")
	d.Print()
}

func newCard() string {
	return "Five of Hearts"
}
```

### Receiver function

```go
package main

func main() {
	d := deck{"Ace of diamonds", newCard()}
	d = append(d, "Six of spades")
	d.Print()
}

func newCard() string {
	return "Five of Hearts"
}
```

### Deck creation

```go
package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

### Slice range

```go

```
