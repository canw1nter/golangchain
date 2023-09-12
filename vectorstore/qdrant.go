package vectorstore

import (
	"context"
	"github.com/pkg/errors"
	pb "github.com/qdrant/go-client/qdrant"
	"golangchain/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type QdrantVectorStoreOption struct {
	Url string
}

type QdrantVectorStore struct {
	collectionClient pb.CollectionsClient
	pointClient      pb.PointsClient
	*QdrantVectorStoreOption
}

func (qd *QdrantVectorStore) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(qd.QdrantVectorStoreOption)
	}
}

func (qd *QdrantVectorStore) Save(data map[string]interface{}) error {
	collectionName, ok := data["collection"].(string)
	if !ok {
		return errors.New("save vector failed: param: collection")
	}
	vectorId, ok := data["vectorId"].(uint64)
	if !ok {
		return errors.New("save vector failed: param: vectorId")
	}
	vector, ok := data["vector"].([]float32)
	if !ok {
		return errors.New("save vector failed: param: vector")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := qd.pointClient.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: collectionName,
		Points: []*pb.PointStruct{
			{
				Id:      &pb.PointId{PointIdOptions: &pb.PointId_Num{Num: vectorId}},
				Vectors: &pb.Vectors{VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: vector}}},
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "save vector failed")
	}

	return nil
}

func (qd *QdrantVectorStore) Search(input string) {
	//TODO implement me
	panic("implement me")
}

func NewQdrantVectorStore(opts ...common.Options) (*QdrantVectorStore, error) {
	vectorStore := &QdrantVectorStore{
		QdrantVectorStoreOption: &QdrantVectorStoreOption{
			Url: "localhost:6334",
		},
	}

	vectorStore.SetOptions(opts...)

	conn, err := grpc.Dial(vectorStore.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "create new qdrant vector store failed")
	}

	vectorStore.pointClient = pb.NewPointsClient(conn)
	vectorStore.collectionClient = pb.NewCollectionsClient(conn)

	return vectorStore, nil
}

func WithUrl(url string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*QdrantVectorStoreOption); ok {
			options.Url = url
		}
	}
}
