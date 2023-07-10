package main

import (
	"indexer/internal/util"
)

func main() {
	//buildFlag := flag.Bool("b", false, "Build the index")
	//flag.Parse()
	//
	//buildValue := *buildFlag

	logger := util.GetLogger()
	logger.Info("Bootstrap", "Starting indexer", nil)

	indexer := util.New("./offline-index", logger)
	//fmt.Println(indexer.Timestamp)
	////indexer.Store()
	//fmt.Println("Finished")
	//
	indexer.BuildIndex(20, 3)

}
