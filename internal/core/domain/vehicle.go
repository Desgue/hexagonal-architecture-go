package domain

const (
	Available Status = iota
	Reserved
	Sold
)

type Vehicle struct {
	ID          string
	Price       int
	VehicleInfo VehicleInfo
	Status      Status
}

type VehicleInfo struct {
	Make  string
	Model string
	Year  int
	Km    int
}

type Status int
