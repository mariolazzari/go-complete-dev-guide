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

### Challenge: even or odd?

```go
package main

import "fmt"

func main() {
	nums := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
```

## Structs

### Structs in Go

Collection of related properties

### Struct definition

```go
package main

type person struct {
	firstName string
	lastName  string
}

func main() {

}
```

### Struct declaration

```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mario := person{"Mario", "Lazzari"}
	fmt.Println(mario)
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	fmt.Println(maria)

}
```

### Updating struct

Zero values

| Type   | Value |
| ------ | ----- |
| string | ""    |
| int    | 0     |
| float  | 0     |
| bool   | false |

```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mario := person{"Mario", "Lazzari"}
	fmt.Println(mario)
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	fmt.Println(maria)

	var alex person
	fmt.Printf("alxe defaults: %+v\n", alex)
	alex.firstName = "Alex"
	alex.lastName = "Lifeson"
	fmt.Printf("alex after assignment: %+v\n", alex)

}
```

### Embedding structs

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	mario := person{"Mario", "Lazzari", contactInfo{"mario.lazzari@gmail.com", 38066}}
	fmt.Println(mario)
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	fmt.Println(maria)

	var alex person
	fmt.Printf("alxe defaults: %+v\n", alex)
	alex.firstName = "Alex"
	alex.lastName = "Lifeson"
	fmt.Printf("alex after assignment: %+v\n", alex)

	mary := person{
		firstName: "Mary",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "mary@mail.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("mary: %+v\n", mary)
}
```

### Structs with receiver functions

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%s after assignment: %+v\n", p.firstName, p)
}

func main() {
	mario := person{"Mario", "Lazzari", contactInfo{"mario.lazzari@gmail.com", 38066}}
	mario.print()
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	maria.print()

	var alex person
	fmt.Printf("alxe defaults: %+v\n", alex)
	alex.firstName = "Alex"
	alex.lastName = "Lifeson"
	alex.print()

	mary := person{
		firstName: "Mary",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "mary@mail.com",
			zipCode: 12345,
		},
	}
	mary.print()

	mary.updateName("Marianne")
	mary.print()
}
```

### Pass by value

```go
func (p person) updateName(name string) {
	p.firstName = name
}
```

### Structs with pointers shortcut

```go
func (p *person) updateName(name string) {
	p.firstName = name
}
```

### Reference vsvalu types

| Value types | Reference types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

## Maps

### What is a map

A collaction of key/value pairs

### Manipulating maps

```go
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
}
```

### Iterating maps

```go
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
```

### Maps vs structs

#### Maps

- all keys of same type
- all values of same type
- keys are indexed
- reference type

#### Structs

- values can be of different types
- keys not indexed
- value type

## Interfaces

### Without interfaces

```go
package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	eb.printGreeting()
	sb.printGreeting()
}

func (englishBot) getGreeting() string {
	// custom logic
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// custom logic
	return "Hola"
}

func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) printGreeting() {
	fmt.Println(sb.getGreeting())
}
```

### Interfaces in practices

```go
package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// custom logic
	return "Hola"
}
```

### Rules on interfaces

- [The bigger the interface, the weaker the abstraction](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)
- [interface{} says nothing](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=456s)

### Interface notes

- not generic type
- implicit
- contract
- understand how to read stdlib docs

### http package

[http](https://pkg.go.dev/net/http)

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error getting google.com", err)
		os.Exit(1)
	}


fmt.Println(resp)

}
```

### Shape challenge

```go
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	b float64
	h float64
}

type square struct {
	l float64
}

func main() {
	s := square{2}
	t := triangle{2, 3}

	printArea(&s)
	printArea(&t)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (t *triangle) getArea() float64 {
	return t.b * t.h / 2
}

func (s *square) getArea() float64 {
	return s.l * s.l
}
```

### Challenge hard mode

```go
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
```

## Channels and Goroutines

### Serial link checking

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		return
	}
	fmt.Printf("%s is up\n", link)
}
```

```sh
time go run main.go
https://google.com is up
https://yahoo.com is up
https://bing.com is up
https://mariolazzari.it is up
go run main.go  0.65s user 0.57s system 39% cpu 3.077 total
```

### Goroutines

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	for _, link := range links {
		// run function in a new goroutine
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		return
	}
	fmt.Printf("%s is up\n", link)
}
```

### Channels

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		// run function in a new goroutine
		go checkLink(link, ch)
	}

	fmt.Println(<-ch)
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- "down"
		return
	}

	ch <- "up"
	fmt.Printf("%s is up\n", link)

	ch <- "up"
	ch <- "up"
	ch <- "up"
	ch <- "up"
}
```

### Blocking channel

```go
func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- "down"
		return
	}

	ch <- "up"
	fmt.Printf("%s is up\n", link)
}
```

### Receiving messages

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		// run function in a new goroutine
		go checkLink(link, ch)
	}

	for range links {
		fmt.Println(<-ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- "down"
		return
	}

	ch <- "up"
	fmt.Printf("%s\n", link)
}
```

### Repeating routines

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for {
		go checkLink(<-ch, ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
```

### Alternative loop syntax

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go checkLink(l, ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
```

### Sleeping a routine

[time](https://pkg.go.dev/time)
[sleep](https://pkg.go.dev/time#Sleep)

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go checkLink(l, ch)
	}
}

func checkLink(link string, ch chan string) {

_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
```

### Function liteals

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func() {
			time.Sleep(time.Second * 5)
			checkLink(l, ch)
		}()
	}
}

func checkLink(link string, ch chan string) {
	time.Sleep(time.Second * 5)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
```

### Channels gotchas

-
