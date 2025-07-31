package models

type Product struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Descripcion   string  `json:"descripcion,omitempty"`
	Precio        string  `json:"precio"`
	PrecioDesc    *string `json:"precio_descuento,omitempty"` 
	Imagen        string  `json:"imagen,omitempty"`
	Stock         int     `json:"stock"`
	FechaCreacion string  `json:"fecha_creacion,omitempty"`
	FechaActual   string  `json:"fecha_actualizacion,omitempty"`
	Publicar      bool    `json:"publicado"`
	Destacado     bool    `json:"destacado"`
}
