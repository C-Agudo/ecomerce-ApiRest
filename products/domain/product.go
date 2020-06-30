package domain

type Product struct {
	Id    int64
	Name  string
	Price string
}

type CreateProductCMD struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type UpdateProductCMD struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
