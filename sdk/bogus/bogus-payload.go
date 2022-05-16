package bogus

type BogusPayload struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Weight float64 `json:"weight"`
}

func NewBogusPayload(name string, age int, weight float64) *BogusPayload {
	return &BogusPayload{
		Name:   name,
		Age:    age,
		Weight: weight,
	}
}
