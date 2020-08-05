package web

import (
	"github.com/sky0621/wht/adapter/store"
	"github.com/sky0621/wht/application/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	wht usecase.Wht

	gcsClient store.CloudStorageClient
}

func NewResolver(wht usecase.Wht, gcsClient store.CloudStorageClient) *Resolver {
	return &Resolver{wht: wht, gcsClient: gcsClient}
}
