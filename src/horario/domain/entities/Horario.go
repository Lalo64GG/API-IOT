package entities

type Horario struct{
	ID		int    `json:"id"`
	Minute	string    `json:"minute"`
	Hour    string    `json:"hour"`
}