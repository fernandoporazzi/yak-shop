package entity

type LabYak struct {
	Name string `xml:"name,attr" json:"name"`
	Age  string `xml:"age,attr" json:"age"`
	Sex  string `xml:"sex,attr" json:"sex"`
}

type LabYakPayload struct {
	Name          string  `json:"name"`
	Age           float64 `json:"age"`
	Sex           string  `json:"sex"`
	AgeLastShaved float64 `json:"age-last-shaved"`
}
