package types

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
)

type types struct{}

func DataTypesDemo() demos.Demo {
	return &types{}
}

func (d *types) Demo() {
	main()
}

func main() {
	fmt.Println(formatter.F().Format("DATA TYPES DEMO"))
	invoice := Invoice{
		ID: 274782,
		InvoiceLineItems: [3]*InvoiceLineItem{
			{
				ID:          1,
				ProductName: "A Product Name for ID 1",
				Processed:   false,
				Quantity:    5,
				PriceExTax:  25.0,
			},
			{
				ID:          5,
				ProductName: "A Product Name for ID 5",
				Processed:   false,
				Quantity:    80,
				PriceExTax:  12.00,
			},
			{
				ID:          83,
				ProductName: "A Product Name for ID 83",
				Processed:   false,
				Quantity:    14,
				PriceExTax:  80.00,
			},
		},
	}
	invoice.ProcessTotals()
	fmt.Println(invoice.String())
	fmt.Println()
}

var taxMap = map[int]float64{
	1:  0.1,
	5:  0,
	83: 0.1,
}

type InvoiceLineItem struct {
	ID          int
	ProductName string
	Processed   bool
	Quantity    int16
	PriceExTax  float64
}

type Invoice struct {
	ID               int
	TotalAmountExGst float64
	TotalAmount      float64
	InvoiceLineItems [3]*InvoiceLineItem
}

func (i *Invoice) ProcessTotals() {
	for _, item := range i.InvoiceLineItems {
		if item.Processed {
			continue
		}
		var taxAmount float64
		if t, ok := taxMap[item.ID]; ok {
			taxAmount = t
		}
		exTax := item.PriceExTax * float64(item.Quantity)
		incTax := exTax * (1 + taxAmount)
		if exTax != 0 {
			i.TotalAmountExGst += exTax
			i.TotalAmount += incTax
		}
		item.Processed = true
	}
}

func (i Invoice) String() string {
	return fmt.Sprintf(`Invoice %d
Total Ex. Tax: %f
Total Inc. Tax: %f`,
		i.ID, i.TotalAmountExGst, i.TotalAmount)
}
