package models

type Transfer struct {
	ID           int64
	UserName     string
	UserDocument string
	Description  string
	Value        float64
	Type         string
}
