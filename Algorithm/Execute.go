package Algorithm

import (
	"sort"
	"strconv"

	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
	"github.com/SommerEngineering/ejob-o-mat/XML"
)

// ExecuteAnswers calculates a recommendation basen on given answers
func ExecuteAnswers(answers map[string]Scheme.Answer) []Scheme.ProductGroup {

	data := XML.GetData()

	numProducts := len(data.ProductsCollection.Products)

	groups := make([]Scheme.ProductGroup, numProducts)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {
		//Calculate points per question for a product
		influence := make(map[string]int8)

		// Calculate influences of each answer
		points := 0
		for _, question := range XML.Questions {
			if question.InternalName[0] == productGroup.InternalName[4] { // Questions type is this progucr group
				v, _ := strconv.Atoi(answers[question.InternalName].Data)
				if v <= 2 {
					influence[question.InternalName] = int8(1)
				} else {
					influence[question.InternalName] = int8(-1)
				}

				points += 5 - v
			}
		}

		groups[n] = Scheme.ProductGroup{
			Name:             productGroup.InternalName,
			Points:           points,
			Percent:          points * 5,
			AnswerInfluences: influence,
		}
	}

	sort.Sort(Scheme.ProductGroupsByPoints(groups))
	return groups
}
