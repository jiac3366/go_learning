package main

import "fmt"

type Land interface {
	getName() string
}

type Sky interface {
	getModel() string
}

type Human struct {
	FirstName, LastName string
}

func (h *Human) getName() string {
	return h.FirstName + h.LastName
}

type Car struct {
	Factory, Model string
}

func (c *Car) getName() string {
	return c.Factory + c.Model
}

type Plane struct {
	Color, Model string
}

func (p *Plane) getModel() string {
	return p.Color + p.Model
}

func main() {
	interface_list := []Land{}
	h := new(Human)
	h.FirstName = "fist"
	h.LastName = "last"
	interface_list = append(interface_list, h)

	// c := new(Car)
	// c.Factory = "fa"
	// c.Model = "mo"
	p := new(Plane)
	p.Color = "fa2222"
	p.Model = "mo2222"
	interface_list = append(interface_list, p)
	fmt.Println("j")
	// for _, val := range interface_list {
	// 	fmt.Println(val.getName())
	// }

}
