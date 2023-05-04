package dto

type CreateProductDTO struct {
	ProductName        string   `json:"pro_name" validate:"required"`
	ProductDescription string   `json:"pro_description" validate:"required"`
	TotalProduct       int64    `json:"total_pro" validate:"required"`
	Price              float64  `json:"price"  validate:"required"`
	Colors             []string `json:"colors" validate:"required"`
	SellerId           string   `json:"seller_id" validate:"required"`
}

type UpdateProductDTO struct {
	Id      string `json:"id" validate:"required" `
	Product CreateProductDTO
}

type CreateCategoriesDTO struct {
	CategoryType        string `json:"cat_type" validate:"required"`
	CategoryDescription string `json:"cat_desc" validate:"required"`
}

type UpdateCategoriesDTO struct {
	Id                  string `json:"id" validate:"required"`
	CategoryType        string `json:"cat_type" validate:"required"`
	CategoryDescription string `json:"cat_desc" validate:"required"`
}

type CreateProductsSellingDTO struct {
	ProductId  string `json:"prod_id" validate:"required"`
	CategoryId string `json:"cat_id" validate:"required"`
}
