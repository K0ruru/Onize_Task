package entity

// User represents a user.
type User struct {
	ID         string
	Name       string
	Passphrase string
	Email      string
	No_telp    string
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetName() string {
	return u.Name
}

func (u User) GetPassphrase() string {
	return u.Passphrase
}

func (u User) GetEmail() string {
	return u.Email
}
