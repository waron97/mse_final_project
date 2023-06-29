package app

import (
	"indexer/src/util"
)

func RunTasks() {
	util.GetCrawlPage()
	logger := util.GetLogger()
	logger.Info("Bootstrap", "Starting indexer", nil)
}
