package domain

const (
	Available int = iota
	Reserved
	Sold
)

type Vehicle struct {
	ID          string      `json:"id"`
	Price       int         `json:"price"`
	VehicleInfo VehicleInfo `json:"vehicleInfo"`
	Status      int         `json:"status"`
}

type VehicleInfo struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Km    int    `json:"km"`
}
