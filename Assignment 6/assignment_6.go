package assignmentsix

type Invoice struct {
	To     string
	Amount int
}

func groupInvoiceByPerson(invoices []Invoice) []Invoice {
	newInvoices := []Invoice{}
	personPoss := map[string]int{}

	idx := 0
	for i := 0; i < len(invoices); i++ {
		invoice := invoices[i]

		personPos, exist := personPoss[invoice.To]
		if !exist {
			personPoss[invoice.To] = idx
			idx++

			newInvoices = append(newInvoices, invoice)
		} else {
			newInvoice := &newInvoices[personPos]

			newInvoice.Amount += invoice.Amount
		}
	}

	return newInvoices
}
