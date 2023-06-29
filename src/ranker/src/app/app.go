package app

import (
	"ranker/src/util"
)

func RunTasks() {
	util.GetCrawlPage()
	util.GetLogger().Info("Bootstrap", "Hello from GO ranker!", nil)
}
