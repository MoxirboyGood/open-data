package domain
type Drug struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Description string `json:"description"`
	Receipt string `json:"receipt"`
	Photo []string `json:"photo"`
}
type DrugSearch struct{
	Id string `json:"id"`
	Name string `json:"name"`
}