package main

import (
	"fmt"
)

type Shawarma interface {
	getDescription() string
	cost() float32
}

//======================================================

type NormSh struct{}
func (i NormSh) cost() float32 {
	return 4.99
}
func (i NormSh) getDescription() string {
	return "Iphone 14: 4.99$\n"
}

//======================================================

type GigaSh struct{}
func (ip GigaSh) cost() float32 {
	return 7.99
}
func (ip GigaSh) getDescription() string {
	return "Iphone 14 Pro: 7.99$\n"
}

//======================================================

type Cheese struct {
	shw Shawarma
}
func (ap Cheese) cost() float32 {
	p := ap.shw.cost()
	return p + 0.50
}
func (ap Cheese) getDescription() string {
	return ap.shw.getDescription() + "\t+ Cheese: 0.50$\n"
}

//======================================================

type Onions struct {
	shw Shawarma
}
func (c Onions) cost() float32 {
	p := c.shw.cost()
	return p + 0.50
}
func (c Onions) getDescription() string {
	return c.shw.getDescription() + "\t+ Onions: 0.50$\n"
}

//======================================================

type fries struct {
	shw Shawarma
}
func (aw fries) cost() float32 {
	p := aw.shw.cost()
	return p + 0.99
}
func (aw fries) getDescription() string {
	return aw.shw.getDescription() + "\t+ fries: 0.99$\n"
}

//======================================================

type Chips struct {
	shw Shawarma
}
func (c Chips) cost() float32 {
	p := c.shw.cost()
	return p + 0.99
}
func (c Chips) getDescription() string {
	return c.shw.getDescription() + "\t+ Chips: 0.99$\n"
}

//======================================================

func mainMenu() {
	fmt.Println("Welcome to ShawarmaOneLove! \n1 - create a new order\n2 - exit")
	var id int
	fmt.Scanf("%d\n", &id)
	switch id {
	case 1:
		chooseBase()
	default:
		fmt.Println("Goodbye!")
	}
}

//======================================================

func subMenu(b Shawarma) {
	fmt.Println("Welcome to ShawarmaOneLove! \n1 - add something\n2 - back to menu")
	var id int
	fmt.Scanf("%d\n", &id)
	switch id {
	case 1:
		chooseComponents(b)
	case 2:
		mainMenu()
	default:
		fmt.Println("Goodbye!")
	}
}

//======================================================

func chooseBase() {
	fmt.Println("List of Shawarma! \n1 - Normal Size\n2 - Mega size")
	var id int
	fmt.Scanf("%d\n", &id)
	var base Shawarma
	switch id {
	case 1:
		base = NormSh{}
	case 2:
		base = GigaSh{}
	showOrder(base)
	subMenu(base)
}

//======================================================

func chooseComponents(base Shawarma) {
	fmt.Println("List of supplements! \n1 - Cheese\n2 - Onions\n3 - Fries\n4 - Chips")
	var id int
	fmt.Scanf("%d\n", &id)
	var baseWithSupplements Shawarma
	switch id {
	case 1:
		baseWithSupplements = Cheese{base}
	case 2:
		baseWithSupplements = Onions{base}
	case 3:
		baseWithSupplements = fries{base}
	case 4:
		baseWithSupplements = Chips{base}
	}
	showOrder(baseWithSupplements)
	subMenu(baseWithSupplements)
}

//======================================================

func showOrder(b Shawarma) {
	fmt.Println("\n Your order")
	fmt.Println(b.getDescription())
	fmt.Printf("Total cost: %v$\n", fmt.Sprintf("%.2f", b.cost()))
}

func main() {
	fmt.Println("Welcome to ShawarmaOneLove!")
	mainMenu()
}