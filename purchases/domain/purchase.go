package domain

type Purchase struct {
	Id       int64
	Shop     string
	Products string
}

type CreatePurchaseCMD struct {
	Shop     string `json:"shop"`
	Products string `json:"products"`
}
