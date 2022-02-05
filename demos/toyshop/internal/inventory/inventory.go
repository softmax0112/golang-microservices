package inventory

type Inventory struct {
	ID            uint   `json:"id"`
	ProductNumber string `json:"productNumber"`
	Quantity      int    `json:"quantity"`
}
