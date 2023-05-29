package functions

import "fmt"

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

// make new bill
func NewBill(name string) *Bill {
	b := Bill{
		Name:  name,
		Items: map[string]float64{},
		Tip:   0,
	}
	return &b
}

// format the bill
func (b *Bill) Format() *string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.Items {
		fs += fmt.Sprintf("-%-25v ...$%v\n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("-%-25v ...$%0.2f\n", "tip:", b.Tip)

	// total
	fs += fmt.Sprintf("-%-25v ...$%0.2f\n", "total:", total)

	return &fs
}

// uptade tip
func (b *Bill) UpdateTip(tip float64) {
	b.Tip = tip
}

// add an item to the bill
func (b *Bill) AddItem(itemName string, price float64) {
	b.Items[itemName] = price
}
