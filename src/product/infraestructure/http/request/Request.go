package request

type CreateProductRequest struct{
	Name        string  		`json:"name" validate:"required"`
	Fecha_Adquisicion string  	`json:"fecha_adquisicion" validate:"required"`
}