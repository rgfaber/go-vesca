package model

type Details struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewDetails(name string, description string) *Details {
	return &Details{
		Name:        name,
		Description: description,
	}
}
