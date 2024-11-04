package main

import (
	"app/utils"
	"app/common"
)

const FILES_DIR = "samples/"

const COLLECTION_NAME = "chunks"

func main(){
	jsonResps := utils.Partition(FILES_DIR + "test.txt")

	var chunks [](common.ChunkFormat)
	for i := 0; i < len(jsonResps); i++ {
		t := jsonResps[i]
		chunks = append(chunks, common.ChunkFormat{
			ElementID: t.ElementID,
			Text: t.Text,
			Embedding: utils.Embed(t.Text),
		})
	}

	utils.CreateCollection(COLLECTION_NAME, 384)

	for i := 0; i < len(chunks); i++ {
		utils.Upsert(COLLECTION_NAME, chunks[i])
	}
}