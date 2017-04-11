package Algorithm

import (
	"sort"

	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
)

// ExecuteAnswers calculates a recommendation basen on given answers
func ExecuteAnswers(answers Scheme.Answers) []Scheme.ProductGroup {

	data := XML.GetData()

	numProducts := len(data.ProductsCollection.Products)

	groups := make([]Scheme.ProductGroup, numProducts)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {
		//Calculate points per question for a product
		groups[n].AnswerInfluences = make(map[string]int8)
		influence := groups[n].AnswerInfluences

		// Calculate influences of each answer
		/*  1 */ influence["Question1"] = kindCommon(answers.A1Data, productGroup.SharedProperties.VideoContent) * int8(answers.A1Weight)
		/*  2 */ influence["Question2"] = kindCommon(answers.A2Data, productGroup.SharedProperties.Tutoring) * int8(answers.A2Weight)
		/*  3 */ influence["Question3"] = kindCommon(answers.A3Data, productGroup.SharedProperties.UserComments) * int8(answers.A3Weight)
		/*  4 */ influence["Question4"] = kindCommon(answers.A4Data, productGroup.SharedProperties.SynchronousInteraction) * int8(answers.A4Weight)
		/*  5 */ influence["Question5"] = kindCommon(answers.A5Data, productGroup.SharedProperties.AsynchronousInteraction) * int8(answers.A5Weight)
		/*  6 */ influence["Question6"] = kindAmountAccesses(answers.A6Data, productGroup.SharedProperties.MinAmountAccesses, productGroup.SharedProperties.MaxAmountAccesses) * int8(answers.A6Weight)
		/*  7 */ influence["Question7"] = kindCommon(answers.A7Data, productGroup.SharedProperties.Downloads) * int8(answers.A7Weight)
		/*  8 */ influence["Question8"] = kindCommon(answers.A8Data, productGroup.SharedProperties.ShowLearningObjectives) * int8(answers.A8Weight)
		/*  9 */ influence["Question9"] = kindOperationType(answers.A9Data, productGroup.SharedProperties.Purpose) * int8(answers.A9Weight)
		/* 10 */ influence["Question10"] = kindCommon(answers.A10Data, productGroup.SharedProperties.CloudBased) * int8(answers.A10Weight)
		/* 11 */ influence["Question11"] = kindCommon(answers.A11Data, productGroup.SharedProperties.Intranet) * int8(answers.A11Weight)
		/* 12 */ influence["Question12"] = kindCommon(answers.A12Data, productGroup.SharedProperties.Assessments) * int8(answers.A12Weight)
		/* 13 */ influence["Question13"] = kindCommon(answers.A13Data, productGroup.SharedProperties.StudentRoles) * int8(answers.A13Weight)
		/* 14 */ influence["Question14"] = kindCommon(answers.A14Data, productGroup.SharedProperties.DisplayEquations) * int8(answers.A14Weight)
		/* 15 */ influence["Question15"] = kindCommon(answers.A15Data, productGroup.SharedProperties.WriteEquations) * int8(answers.A15Weight)
		/* 16 */ influence["Question16"] = kindCommon(answers.A16Data, productGroup.SharedProperties.TeachingTypePresentation) * int8(answers.A16Weight)
		/* 17 */ influence["Question17"] = kindCommon(answers.A17Data, productGroup.SharedProperties.TeachingTypeDevelopment) * int8(answers.A17Weight)
		/* 18 */ influence["Question18"] = kindCommon(answers.A18Data, productGroup.SharedProperties.TeachingTypeExplorative) * int8(answers.A18Weight)

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
	bestPointsNormal := float64(int(answers.A1Weight) + int(answers.A2Weight) + int(answers.A3Weight) + int(answers.A4Weight) + int(answers.A5Weight) + int(answers.A6Weight) + int(answers.A7Weight) + int(answers.A8Weight) + int(answers.A9Weight) + int(answers.A10Weight) + int(answers.A11Weight) + int(answers.A12Weight) + int(answers.A13Weight) + int(answers.A14Weight) + int(answers.A15Weight) + int(answers.A16Weight) + int(answers.A17Weight) + int(answers.A18Weight))

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
