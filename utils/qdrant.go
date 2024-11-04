package utils

import (
	"app/common"
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
	"github.com/google/uuid"
)

var client, _ = qdrant.NewClient(&qdrant.Config{
	Host: "localhost",
	Port: 6334,
})

func CreateCollection(collectionName string, vectorSize int) {
	if exists, _ := client.CollectionExists(context.Background(), collectionName); !exists {
		err := client.CreateCollection(context.Background(), &qdrant.CreateCollection{
			CollectionName: collectionName,
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
				Size:     uint64(vectorSize),
				Distance: qdrant.Distance_Cosine,
			}),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Creating collection " + collectionName)
	}
	fmt.Println("Collection " + collectionName + " already exists")
}

func Upsert(collectionName string, chunk common.ChunkFormat) {

	_, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDUUID(uuid.New().String()),
				Vectors: qdrant.NewVectorsDense(chunk.Embedding),
				Payload: qdrant.NewValueMap(map[string]any{"text": chunk.Text}),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Upserted chunk")

}


