package vectorstore

type IVectorStore interface {
	Save(data map[string]interface{}) error
	Search(input string)
}
