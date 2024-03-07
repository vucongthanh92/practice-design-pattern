package main

import "fmt"

type Item interface {
	GetCost() float32
}

type Product struct {
	Name     string
	Price    float32
	Quantity int
}

func (product Product) GetCost() float32 {
	if product.Quantity == 0 {
		return product.Price
	}
	return product.Price * float32(product.Quantity)
}

type Package struct {
	children  []Item
	ExtraCost float32
}

func (pack Package) GetCost() float32 {
	var totalCost float32
	for _, item := range pack.children {
		totalCost += item.GetCost() + pack.ExtraCost
	}
	return totalCost
}

func CreatePackage() Item {
	return Package{
		children: []Item{
			Product{
				Name:     "Mouse",
				Price:    12,
				Quantity: 2,
			},
			Package{
				children: []Item{
					Product{
						Name:     "Main",
						Price:    20,
						Quantity: 1,
					},
					Package{
						children: []Item{
							Product{
								Name:     "Ram 1",
								Price:    15,
								Quantity: 1,
							},
							Product{
								Name:     "Ram 2",
								Price:    14,
								Quantity: 2,
							},
							Product{
								Name:  "Ram 3",
								Price: 16,
							},
						},
						ExtraCost: 2,
					},
					Product{
						Name:  "CPU",
						Price: 32,
					},
				},
				ExtraCost: 3,
			},
			Package{
				children: []Item{
					Product{
						Name:     "Monitor",
						Price:    40,
						Quantity: 2,
					},
					Product{
						Name:  "Adapter",
						Price: 5,
					},
				},
			},
			Product{
				Name:     "Keyboard",
				Price:    8,
				Quantity: 3,
			},
		},
	}
}

func main() {
	pack1 := CreatePackage()
	fmt.Println(pack1.GetCost())
}
