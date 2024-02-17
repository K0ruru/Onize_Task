package entity

// User represents a user.
type Users struct {
	ID       string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	No_telp  string `json:"no_telp"`
}

// GetID returns the user ID.
func (u Users) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u Users) GetName() string {
	return u.Name
}
