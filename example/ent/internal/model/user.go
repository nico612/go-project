package model

type User struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
	Car  []*Car
}

type Car struct {
	Model        string `json:"model"`
	RegisteredAt string `json:"registeredAt"`
}
