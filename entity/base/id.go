package base

// Increment is
type Increment struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

// UUID is
type UUID struct {
	ID string `gorm:"primaryKey; type:varchar(36)" json:"id"`
}
