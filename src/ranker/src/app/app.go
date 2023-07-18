package app

import (
	colbert "ranker/src/util/ColbertLateInteraction"
	util "ranker/src/util/core"
)

func RunTasks() {
	util.GetLogger().Info("Bootstrap", "Hello from GO ranker!", nil)
	documents := getMockDocuments()
	query := util.GetEmbeddings("hello", "world")
	colbert.Rank(documents, query)
	// fmt.Println(rankResult)
}
