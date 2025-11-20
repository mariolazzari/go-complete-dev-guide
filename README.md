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

```go

```
