package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Clone() Point {
	return Point{
		X: p.X,
		Y: p.Y,
	}
}

type Node struct {
	Value    Point
	Children []Point
}

func (n *Node) Clone() Node {
	children := make([]Point, len(n.Children))
	for i := range children {
		children[i] = n.Children[i]
	}
	return Node{
		Value:    n.Value.Clone(),
		Children: children,
	}
}
func main() {
	p1 := Point{1, 2}
	p2 := p1.Clone()

	p1.X = 5
	fmt.Println(p1, p2)

	n1 := Node{
		Value: p1,
		Children: []Point{
			{1, 2},
			{2, 3},
		},
	}

	n2 := n1.Clone()
	n1.Children[0].X = 10
	fmt.Println(n1.Children[0].X, n2.Children[0].X)
}

type Color struct {
	red, green, blue uint8
}

func (c *Color) ColorWithRed(r uint8) Color {
	return Color{red: r, green: c.green, blue: c.blue}
}
