package main

import (
	"billMaker/functions"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func CreateBill() functions.Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := GetInput(reader, "Create new bill name:")

	b := functions.NewBill(name)
	fmt.Println("Created the bill -", name)

	return *b
}

func PromptOptions(b functions.Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := GetInput(reader, "Choose option (a - add item, s - save bill, t - add tip) -- ")

	switch opt {
	case "a":
		name, _ := GetInput(reader, "Item name:")
		price, _ := GetInput(reader, "Item price:")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number!")
			PromptOptions(b)
		}

		b.AddItem(name, p)
		fmt.Println("Item is added. Anything else?")

		PromptOptions(b)
	case "t":
		tip, _ := GetInput(reader, "Enter tip amount ($):")
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number!")
			PromptOptions(b)
		}

		fmt.Println("Thank you for tip <3")
		b.UpdateTip(t)

		PromptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill.")
		fmt.Println(*b.Format())
	default:
		fmt.Println("That was not valid option...")
		PromptOptions(b)
	}
}

func main() {
	myBill := CreateBill()
	PromptOptions(myBill)
}
