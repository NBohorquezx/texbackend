package models

type Orderp struct {
	ID           int
	ProductID    int
	Cantidad     int
	PrecioUnit   float64
	PrecioTotal  float64 // Calculado: cantidad * precio_unit
}

