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

		influence["Question1"] = kindCommon(answers["Question1"].Data, productGroup.SharedProperties.VideoContent) * int8(answers["Question1"].Weight)
		influence["Question2"] = kindCommon(answers["Question2"].Data, productGroup.SharedProperties.Tutoring) * int8(answers["Question2"].Weight)
		influence["Question3"] = kindCommon(answers["Question3"].Data, productGroup.SharedProperties.UserComments) * int8(answers["Question3"].Weight)
		influence["Question4"] = kindCommon(answers["Question4"].Data, productGroup.SharedProperties.SynchronousInteraction) * int8(answers["Question4"].Weight)
		influence["Question5"] = kindCommon(answers["Question5"].Data, productGroup.SharedProperties.AsynchronousInteraction) * int8(answers["Question5"].Weight)
		influence["Question6"] = kindAmountAccesses(answers["Question6"].Data, productGroup.SharedProperties.MinAmountAccesses, productGroup.SharedProperties.MaxAmountAccesses) * int8(answers["Question6"].Weight)
		influence["Question7"] = kindCommon(answers["Question7"].Data, productGroup.SharedProperties.Downloads) * int8(answers["Question7"].Weight)
		influence["Question8"] = kindCommon(answers["Question8"].Data, productGroup.SharedProperties.ShowLearningObjectives) * int8(answers["Question8"].Weight)
		influence["Question9"] = kindOperationType(answers["Question9"].Data, productGroup.SharedProperties.Purpose) * int8(answers["Question9"].Weight)
		influence["Question10"] = kindCommon(answers["Question10"].Data, productGroup.SharedProperties.CloudBased) * int8(answers["Question10"].Weight)
		influence["Question11"] = kindCommon(answers["Question11"].Data, productGroup.SharedProperties.Intranet) * int8(answers["Question11"].Weight)
		influence["Question12"] = kindCommon(answers["Question12"].Data, productGroup.SharedProperties.Assessments) * int8(answers["Question12"].Weight)
		influence["Question13"] = kindCommon(answers["Question13"].Data, productGroup.SharedProperties.StudentRoles) * int8(answers["Question13"].Weight)
		influence["Question14"] = kindCommon(answers["Question14"].Data, productGroup.SharedProperties.DisplayEquations) * int8(answers["Question14"].Weight)
		influence["Question15"] = kindCommon(answers["Question15"].Data, productGroup.SharedProperties.WriteEquations) * int8(answers["Question15"].Weight)
		influence["Question16"] = kindCommon(answers["Question16"].Data, productGroup.SharedProperties.TeachingTypePresentation) * int8(answers["Question16"].Weight)
		influence["Question17"] = kindCommon(answers["Question17"].Data, productGroup.SharedProperties.TeachingTypeDevelopment) * int8(answers["Question17"].Weight)
		influence["Question18"] = kindCommon(answers["Question18"].Data, productGroup.SharedProperties.TeachingTypeExplorative) * int8(answers["Question18"].Weight)

		//Total points
		for _, v := range influence {
			groups[n].Points += int(v)
		}

		groups[n].Name = productGroup.InternalName
		groups[n].XMLIndex = n //TODO Use id/internal name instead of XMLIndex, as it might not be consistent when question are added or removed
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
