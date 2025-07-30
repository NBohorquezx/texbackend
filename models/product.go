package models

import "time"

type Product struct {
	ID               int
	Nombre           string
	Descripcion      string
	Precio           float64
	PrecioDescuento  float64
	RutaImg          string
	Stock            int
	CategoriaID      int
	FechaCreacion    time.Time
	FechaActualizacion time.Time
	Publicado        bool
	Destacado        bool
}
