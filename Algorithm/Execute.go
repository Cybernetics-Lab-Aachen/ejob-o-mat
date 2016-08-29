package Algorithm

import (
	"sort"

	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
)

func ExecuteAnswers(answers Scheme.Answers) (result Scheme.ProductGroups, influence [][]int) {

	data := XML.GetData()

	numProducts := len(data.ProductsCollection.Products)

	groups := make(Scheme.ProductGroups, numProducts)

	influence = make([][]int, numProducts)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {
		//Calculate points per question for a product
		influence[n] = make([]int, 18)

		/*  1 */ influence[n][0] = kindCommon(answers.A1Data, productGroup.SharedProperties.VideoContent) * int(answers.A1Weight)
		/*  2 */ influence[n][1] = kindCommon(answers.A2Data, productGroup.SharedProperties.Tutoring) * int(answers.A2Weight)
		/*  3 */ influence[n][2] = kindCommon(answers.A3Data, productGroup.SharedProperties.UserComments) * int(answers.A3Weight)
		/*  4 */ influence[n][3] = kindCommon(answers.A4Data, productGroup.SharedProperties.SynchronousInteraction) * int(answers.A4Weight)
		/*  5 */ influence[n][4] = kindCommon(answers.A5Data, productGroup.SharedProperties.AsynchronousInteraction) * int(answers.A5Weight)
		/*  6 */ influence[n][5] = kindAmountAccesses(answers.A6Data, productGroup.SharedProperties.MinAmountAccesses, productGroup.SharedProperties.MaxAmountAccesses) * int(answers.A6Weight)
		/*  7 */ influence[n][6] = kindCommon(answers.A7Data, productGroup.SharedProperties.Downloads) * int(answers.A7Weight)
		/*  8 */ influence[n][7] = kindCommon(answers.A8Data, productGroup.SharedProperties.ShowLearningObjectives) * int(answers.A8Weight)
		/*  9 */ influence[n][8] = kindOperationType(answers.A9Data, productGroup.SharedProperties.Purpose) * int(answers.A9Weight)
		/* 10 */ influence[n][9] = kindCommon(answers.A10Data, productGroup.SharedProperties.CloudBased) * int(answers.A10Weight)
		/* 11 */ influence[n][10] = kindCommon(answers.A11Data, productGroup.SharedProperties.Intranet) * int(answers.A11Weight)
		/* 12 */ influence[n][11] = kindCommon(answers.A12Data, productGroup.SharedProperties.Assessments) * int(answers.A12Weight)
		/* 13 */ influence[n][12] = kindCommon(answers.A13Data, productGroup.SharedProperties.StudentRoles) * int(answers.A13Weight)
		/* 14 */ influence[n][13] = kindCommon(answers.A14Data, productGroup.SharedProperties.DisplayEquations) * int(answers.A14Weight)
		/* 15 */ influence[n][14] = kindCommon(answers.A15Data, productGroup.SharedProperties.WriteEquations) * int(answers.A15Weight)
		/* 16 */ influence[n][15] = kindCommon(answers.A16Data, productGroup.SharedProperties.TeachingTypePresentation) * int(answers.A16Weight)
		/* 17 */ influence[n][16] = kindCommon(answers.A17Data, productGroup.SharedProperties.TeachingTypeDevelopment) * int(answers.A17Weight)
		/* 18 */ influence[n][17] = kindCommon(answers.A18Data, productGroup.SharedProperties.TeachingTypeExplorative) * int(answers.A18Weight)

		//Total points
		//for i := range influence[n] {
		for _, v := range influence[n] {
			groups[n].Points += v
		}

		groups[n].Name = productGroup.InternalName
		groups[n].XMLIndex = n
	}

	//
	// Re-align the results to respect the range from 0-100%:
	//
	sort.Sort(groups)
	worstPoints := groups[len(groups)-1].Points
	correctionPoints := worstPoints * -1
	bestPointsNormal := float64(int(answers.A1Weight) + int(answers.A2Weight) + int(answers.A3Weight) + int(answers.A4Weight) + int(answers.A5Weight) + int(answers.A6Weight) + int(answers.A7Weight) + int(answers.A8Weight) + int(answers.A9Weight) + int(answers.A10Weight) + int(answers.A11Weight) + int(answers.A12Weight) + int(answers.A13Weight) + int(answers.A14Weight) + int(answers.A15Weight) + int(answers.A16Weight) + int(answers.A17Weight) + int(answers.A18Weight))
	bestPointsCorrected := bestPointsNormal + float64(correctionPoints)

	for n, _ := range groups {
		if worstPoints < 0 {
			groups[n].Points += correctionPoints
			result := (float64(groups[n].Points) / bestPointsCorrected) * 100.0
			groups[n].Percent = int(result)
		} else {
			result := (float64(groups[n].Points) / bestPointsNormal) * 100.0
			groups[n].Percent = int(result)
		}
	}

	result = groups //[0-6]
	return
}
