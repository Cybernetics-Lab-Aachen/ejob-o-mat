package Algorithm

import (
	"fmt"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/XML"
)

func ExecuteAnswers(answers DB.Answers) (result []DB.ProductGroup) {

	data := XML.GetData()
	groups := make([]DB.ProductGroup, 18)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {

		/*  1 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A1Data, productGroup.SharedProperties.Exam)
		/*  2 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A2Data, productGroup.SharedProperties.Assistant)
		/*  3 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A3Data, productGroup.SharedProperties.UserComments)
		/*  4 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A4Data, productGroup.SharedProperties.AnonymousUsers)
		/*  5 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A5Data, productGroup.SharedProperties.LiveCollaboration)
		/*  6 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A6Data, productGroup.SharedProperties.CommunityCollaboration)
		/*  7 */ groups[n].Points = groups[n].Points + kindAppropriateCountStudents(answers.A7Data, productGroup.SharedProperties.AppropriateCountStudents)
		/*  8 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A8Data, productGroup.SharedProperties.Downloads)
		/*  9 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A9Data, productGroup.SharedProperties.Possibility2DeclareLearningObjectives)
		/* 10 */ groups[n].Points = groups[n].Points + kindOperationType(answers.A10Data, productGroup.SharedProperties.OperationType)
		/* 11 */ groups[n].Points = groups[n].Points + kindCosts(answers.A11Data, productGroup.SharedProperties.Costs, productGroup.SharedProperties.AlsoFreeProducts)
		/* 12 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A12Data, productGroup.SharedProperties.CloudBased)
		/* 13 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A13Data, productGroup.SharedProperties.Intranet)
		/* 14 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A14Data, productGroup.SharedProperties.StandaloneSoftware)
		/* 15 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A15Data, productGroup.SharedProperties.SCROMSupport)
		/* 16 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A16Data, productGroup.SharedProperties.VideoContent)
		/* 17 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A17Data, productGroup.SharedProperties.HighAvailability)
		/* 18 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A18Data, productGroup.SharedProperties.StudentRoles)
		/* 19 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A19Data, productGroup.SharedProperties.TrackedProgress)
		/* 20 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A20Data, productGroup.SharedProperties.DisplayEquations)
		/* 21 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A21Data, productGroup.SharedProperties.WriteEquations)
		/* 22 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A22Data, productGroup.SharedProperties.ContentType)
		/* 23 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A23Data, productGroup.SharedProperties.HomeUse)
		/* 24 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A24Data, productGroup.SharedProperties.TeachingTypePresentation)
		/* 25 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A25Data, productGroup.SharedProperties.TeachingTypeDevelopment)
		/* 26 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A26Data, productGroup.SharedProperties.TeachingTypeExplorative)

		result := (float64(data.ProductsCollection.Products[n].Score) / 26.0) * 100.0
		groups[n].Percent = fmt.Sprintf("%.f", result)
		groups[n].XMLIndex = n
		groups[n].Name = productGroup.InternalName
	}

	// Take the best:
	/*for _, productGroup := range groups {

		//groups[found] = productGroup
	}*/

	result = groups[0:5]
	return
}
