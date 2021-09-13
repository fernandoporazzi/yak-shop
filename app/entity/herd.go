package entity

type Herd struct {
	LabYaks []LabYak `xml:"labyak"`
}

type HerdPayload struct {
	Herd []LabYakPayload `json:"herd"`
}
