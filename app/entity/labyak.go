package entity

type LabYak struct {
	Name string `xml:"name,attr" json:"name"`
	Age  string `xml:"age,attr" json:"age"`
	Sex  string `xml:"sex,attr" json:"sex"`
}
