package base

// UUID is
type UUID struct {
	ID string `gorm:"primaryKey; type:char(36)" json:"id"`
}
