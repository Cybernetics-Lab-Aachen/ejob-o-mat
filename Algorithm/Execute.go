package Algorithm

import (
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/XML"
)

func ExecuteAnswers(answers DB.Answers) (groups []XML.ProductGroup) {

	data := XML.GetData()
	groups = make([]XML.ProductGroup, 6)
	found := 0

	for _, productGroup := range data.ProductsCollection.Products {

		groups[found] = productGroup
		groups[found].Score = `66`
		found++

		if found == 6 {
			return
		}
	}

	return
}
