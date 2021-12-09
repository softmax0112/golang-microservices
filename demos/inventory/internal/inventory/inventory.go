package inventory

// This should actually be pulled into a separate package
// since used from multiple locations
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
