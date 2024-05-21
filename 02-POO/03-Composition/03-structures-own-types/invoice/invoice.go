package invoice

import (
	"03-Composition/customer"
	invoiceitem "03-Composition/invoiceItem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
