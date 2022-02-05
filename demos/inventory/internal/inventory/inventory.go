package inventory

type Inventory struct {
	ID            uint   `json:"id"`
	ProductNumber string `json:"productNumber"`
	Quantity      int    `json:"quantity"`
}

type InventoryDatastore interface {
	CreateInventory(inventory *Inventory) error
	UpdateInventory(id string, inventory Inventory) error
	GetInventory(productNumber string) (Inventory, error)
}
