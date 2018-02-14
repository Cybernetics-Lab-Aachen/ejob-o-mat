package Algorithm

import (
	"sort"

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
		groups[n].AnswerInfluences = make(map[string]int8)
		influence := groups[n].AnswerInfluences

		// Calculate influences of each answer
		for _, question := range XML.Questions {
			v := 0 // TODO

			influence[question.InternalName] = int8(v)
			groups[n].Points += int(v)
		}

		groups[n].Name = productGroup.InternalName
	}

	//
	// Re-align the results to respect the range from 0-100%:
	//
	sort.Sort(Scheme.ProductGroupsByPoints(groups))
	worstPoints := groups[len(groups)-1].Points
	bestPointsNormal := 0.0
	for _, answer := range answers {
		bestPointsNormal += float64(answer.Weight)
	}

	for n := range groups {
		if worstPoints < 0 {
			groups[n].Points -= worstPoints // Normalize points so that worst has zero points
			// percent = (points - worst) / (maxPoints - worst) * 100
			percent := float64(groups[n].Points) / (bestPointsNormal - float64(worstPoints)) * 100.0
			groups[n].Percent = int(percent)
		} else {
			// percent = points / maxPoints * 100
			percent := float64(groups[n].Points) / bestPointsNormal * 100.0
			groups[n].Percent = int(percent)
		}
	}

	return groups
}
