package vector

import "context"

type Document struct {
	ID        string            `json:"id"`
	Content   string            `json:"content"`
	Metadata  map[string]string `json:"metadata"`
	Embedding []float32         `json:"embedding"`
}

type Result struct {
	Document
	Score float32 `json:"score"`
}

type Store interface {
	Upsert(ctx context.Context, docs []Document) error
	Query(ctx context.Context, embedding []float32, topK int) ([]Result, error)
	Delete(ctx context.Context, ids []string) error
	DeleteBySource(ctx context.Context, source string) error
	Close() error
}
