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

### Multiple return values

```go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

### Byte slices

[io](https://pkg.go.dev/io)
[os](https://pkg.go.dev/os)

[]byte -> asci code array

### Deck to string

```go
func (d deck) toString() string {
	return ""
}
```

### Joining slices

[strings](https://pkg.go.dev/strings)
[join](https://pkg.go.dev/strings#Join)

```go
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

### Save data to file

[WriteFile](https://pkg.go.dev/os#example-WriteFile)

```go
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
```

### Reading from file

[ReadFile](https://pkg.go.dev/os#example-ReadFile)

```go
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}
	// ...
}
```

### Error handling

[split](https://pkg.go.dev/strings#Split)

```go
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}
```

### Shuffling cards

[rand](https://pkg.go.dev/math/rand)

```go
func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
```

### Random number generation

```go
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
```

### Creating go.mod

```sh
go mod init cards
```

### Testing in Go

```go
go test
```

### Writing tests in Go

```go
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

}
```

### Asserting elements in a slice

```go
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[0])
	}
}
```

### Testing file IO

[os](https://pkg.go.dev/os#pkg-functions)

```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
```
