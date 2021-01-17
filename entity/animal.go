package entity

// Animal is
type Animal struct {
	Increment
	Name string
	Timestamps
}

// TableName is
func (Animal) TableName() string {
	return "animals"
}
