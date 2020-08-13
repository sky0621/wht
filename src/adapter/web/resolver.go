package web

import (
	"github.com/sky0621/wht/adapter/storage"
	"github.com/sky0621/wht/application/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	wht usecase.Wht

	gcsClient storage.CloudStorageClient
}

func NewResolver(wht usecase.Wht, gcsClient storage.CloudStorageClient) *Resolver {
	return &Resolver{wht: wht, gcsClient: gcsClient}
}
