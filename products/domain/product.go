package domain

type Product struct {
	Id int64
	Name string
	Price string	
}

type CreateProductCMD struct{
	Name string 'json:"name"'
	Price string 'json:"price"'
}