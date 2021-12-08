package toys

type Toy struct {
	ID            uint    `json:"id"`
	ProductNumber string  `json:"productNumber"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	UnitCost      float32 `json:"unitCost"`
}

type ToysDatastore interface {
	CreateToy(toy *Toy) error
	GetAllToys() ([]Toy, error)
	UpdateToy(id string, toy Toy) error
	DeleteToy(id string) error
	GetToy(id string) (Toy, error)
}
