package handlers

import (
	"context"
	"errors"

	"github.com/kurabier/marketplace-api/internal/gen"
)

var errNotImplemented = errors.New("not implemented")

type Handler struct{}

func New() *Handler { return &Handler{} }

var _ gen.StrictServerInterface = (*Handler)(nil)

func (h *Handler) CreateProduct(ctx context.Context, req gen.CreateProductRequestObject) (gen.CreateProductResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handler) ListProducts(ctx context.Context, req gen.ListProductsRequestObject) (gen.ListProductsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handler) GetProduct(ctx context.Context, req gen.GetProductRequestObject) (gen.GetProductResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handler) UpdateProduct(ctx context.Context, req gen.UpdateProductRequestObject) (gen.UpdateProductResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handler) DeleteProduct(ctx context.Context, req gen.DeleteProductRequestObject) (gen.DeleteProductResponseObject, error) {
	return nil, errNotImplemented
}
