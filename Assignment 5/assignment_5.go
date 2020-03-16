package assignmentfive

type Case struct {
	Order int
	Name  string
}

var cases = []Case{
	{Order: 1, Name: "one"},
	{Order: 2, Name: "two"},
	{Order: 3, Name: "tri"},
	{Order: 4, Name: "fou"},
	{Order: 5, Name: "fiv"},
}

func rearrange(cases []Case, pastPos int, newPos int) []Case {
	var newCases []Case

	casesMap := map[int]string{}

	for ord := 0; ord < len(cases); ord++ {
		caseData := cases[ord]

		casesMap[caseData.Order] = caseData.Name
	}

	if pastPos < newPos {
		inc := 0
		for ord := 0; ord < len(cases); ord++ {
			caseData := cases[ord]

			if caseData.Order == pastPos {
				inc = 1
			}

			if caseData.Order == newPos {
				newCase := Case{Order: caseData.Order, Name: casesMap[pastPos]}

				newCases = append(newCases, newCase)
				inc = 0
			} else {
				prevCase := cases[ord+inc]
				newCase := Case{Order: caseData.Order, Name: prevCase.Name}
				newCases = append(newCases, newCase)
			}
		}
	} else {
		inc := 0
		for ord := 0; ord < len(cases); ord++ {
			caseData := cases[ord]

			if caseData.Order == newPos {
				newCase := Case{Order: caseData.Order, Name: casesMap[pastPos]}

				newCases = append(newCases, newCase)
				inc = -1
			} else {
				prevCase := cases[ord+inc]
				newCase := Case{Order: caseData.Order, Name: prevCase.Name}
				newCases = append(newCases, newCase)
			}

			if caseData.Order == pastPos {
				inc = 0
			}
		}
	}

	return newCases
}
