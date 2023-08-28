package main

import (
	"fmt"


)

type Item struct {
	Name string
	Price float64
	Discount float64
}

func calculatePrice(item Item)float64{
	if item.Discount != 0{
		return ( item.Price - item.Discount )
	}else {
		return item.Price
	}
}

func totalPrice(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

type Describable interface {
	Description() string
}

func (item Item) Description() string {
	if item.Discount != 0 {
		var discountedPrice float64 = item.Price - item.Discount
		return fmt.Sprintf("%s - %.2f TL (Discounted Price: %.2f TL)",item.Name, item.Price, discountedPrice)
	}
	return fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
}


func main() {
	var items = []Item{
	{Name: "item1", Price: 10, Discount: 5},
	{Name: "item2", Price: 20, Discount: 10},
	{Name: "item3", Price: 30, Discount: 20},
}

	for _,item := range items{
		fmt.Println(item.Description())
	}

	var totalPrice float64 = totalPrice(items)
	fmt.Printf("Total price of items: %.2f",totalPrice)







}




