package Algorithm

import (
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
	"sort"
)

func ExecuteAnswers(answers Scheme.Answers) (result Scheme.ProductGroups) {

	data := XML.GetData()
	groups := make(Scheme.ProductGroups, 16)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {

		/*  1 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A1Data, productGroup.SharedProperties.VideoContent) * int(answers.A1Weight))
		/*  2 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A2Data, productGroup.SharedProperties.Tutoring) * int(answers.A2Weight))
		/*  3 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A3Data, productGroup.SharedProperties.UserComments) * int(answers.A3Weight))
		/*  4 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A4Data, productGroup.SharedProperties.SynchronousInteraction) * int(answers.A4Weight))
		/*  5 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A5Data, productGroup.SharedProperties.AsynchronousInteraction) * int(answers.A5Weight))
		/*  6 */ groups[n].Points = groups[n].Points + (kindAmountAccesses(answers.A6Data, productGroup.SharedProperties.MinAmountAccesses, productGroup.SharedProperties.MaxAmountAccesses) * int(answers.A6Weight))
		/*  7 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A7Data, productGroup.SharedProperties.Downloads) * int(answers.A7Weight))
		/*  8 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A8Data, productGroup.SharedProperties.ShowLearningObjectives) * int(answers.A8Weight))
		/*  9 */ groups[n].Points = groups[n].Points + (kindOperationType(answers.A9Data, productGroup.SharedProperties.Purpose) * int(answers.A9Weight))
		/* 10 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A10Data, productGroup.SharedProperties.CloudBased) * int(answers.A10Weight))
		/* 11 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A11Data, productGroup.SharedProperties.Intranet) * int(answers.A11Weight))
		/* 12 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A12Data, productGroup.SharedProperties.Assessments) * int(answers.A12Weight))
		/* 13 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A13Data, productGroup.SharedProperties.StudentRoles) * int(answers.A13Weight))
		/* 14 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A14Data, productGroup.SharedProperties.DisplayEquations) * int(answers.A14Weight))
		/* 15 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A15Data, productGroup.SharedProperties.WriteEquations) * int(answers.A15Weight))
		/* 16 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A16Data, productGroup.SharedProperties.TeachingTypePresentation) * int(answers.A16Weight))
		/* 17 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A17Data, productGroup.SharedProperties.TeachingTypeDevelopment) * int(answers.A17Weight))
		/* 18 */ groups[n].Points = groups[n].Points + (kindCommon(answers.A18Data, productGroup.SharedProperties.TeachingTypeExplorative) * int(answers.A18Weight))

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
