package Algorithm

import (
	"fmt"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/XML"
)

func ExecuteAnswers(answers DB.Answers) (groups []XML.ProductGroup) {

	data := XML.GetData()
	groups = make([]XML.ProductGroup, 6)
	found := 0

	// Algorithm:

	for n, productGroup := range data.ProductsCollection.Products {

		/*  1 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A1Data, productGroup.SharedProperties.Exam)
		/*  2 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A2Data, productGroup.SharedProperties.Assistant)
		/*  3 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A3Data, productGroup.SharedProperties.UserComments)
		/*  4 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A4Data, productGroup.SharedProperties.AnonymousUsers)
		/*  5 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A5Data, productGroup.SharedProperties.LiveCollaboration)
		/*  6 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A6Data, productGroup.SharedProperties.CommunityCollaboration)
		/*  7 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindAppropriateCountStudents(answers.A7Data, productGroup.SharedProperties.AppropriateCountStudents)
		/*  8 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A8Data, productGroup.SharedProperties.Downloads)
		/*  9 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A9Data, productGroup.SharedProperties.Possibility2DeclareLearningObjectives)
		/* 10 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindOperationType(answers.A10Data, productGroup.SharedProperties.OperationType)
		/* 11 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindCosts(answers.A11Data, productGroup.SharedProperties.Costs, productGroup.SharedProperties.AlsoFreeProducts)
		/* 12 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A12Data, productGroup.SharedProperties.CloudBased)
		/* 13 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A13Data, productGroup.SharedProperties.Intranet)
		/* 14 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A14Data, productGroup.SharedProperties.StandaloneSoftware)
		/* 15 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A15Data, productGroup.SharedProperties.SCROMSupport)
		/* 16 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A16Data, productGroup.SharedProperties.VideoContent)
		/* 17 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A17Data, productGroup.SharedProperties.HighAvailability)
		/* 18 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A18Data, productGroup.SharedProperties.StudentRoles)
		/* 19 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPresence(answers.A19Data, productGroup.SharedProperties.TrackedProgress)
		/* 20 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A20Data, productGroup.SharedProperties.DisplayEquations)
		/* 21 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A21Data, productGroup.SharedProperties.WriteEquations)
		/* 22 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalYesNo(answers.A22Data, productGroup.SharedProperties.ContentType)
		/* 23 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A23Data, productGroup.SharedProperties.HomeUse)
		/* 24 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A24Data, productGroup.SharedProperties.TeachingTypePresentation)
		/* 25 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A25Data, productGroup.SharedProperties.TeachingTypeDevelopment)
		/* 26 */ data.ProductsCollection.Products[n].Score = data.ProductsCollection.Products[n].Score + kindConditionalPossibility(answers.A26Data, productGroup.SharedProperties.TeachingTypeExplorative)

		result := (float64(productGroup.Score) / 26.0) * 100.0
		data.ProductsCollection.Products[n].ScoreText = fmt.Sprintf("%.f", result)
	}

	// Take the best:
	for _, productGroup := range data.ProductsCollection.Products {

		groups[found] = productGroup
		found++

		if found == 6 {
			return
		}
	}

	return
}
