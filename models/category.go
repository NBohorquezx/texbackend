package models

import "time"

type Category struct {
	ID                int
	Nombre            string
	Descripcion       string
	Activo            bool
	FechaCreacion     time.Time
	FechaActualizacion time.Time
}
