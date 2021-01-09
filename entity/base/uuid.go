package base

// UUID is
type UUID struct {
	ID string `gorm:"primaryKey; type:varchar(36)" json:"id"`
}
