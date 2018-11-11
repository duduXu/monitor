package tree

type Order struct {
	Price int
	State int
	Title string
}

func (order *Order) GetPrice() (q int, p int) {
	price := 15
	order.Price = price
	return 1, 15
}
