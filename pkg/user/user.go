package user

import (
	"errors"
)

// User represents a user entity in the database.
type User struct {
	ID       int    `db:"User_id"`
	Name     string `db:"User_name"`
	Email    string `db:"User_Email"`
	Password string `db:"User_Password"`
	Role     string `db:"User_Role"`
	About    string `db:"User_About"`
}

// NewUser creates a new User instance.
func NewUser(id int, name, email, password, role, about string) (*User, error) {
	if name == "" {
		return nil, errors.New("Name cannot be blank")
	}
	if len(about) > 500 {
		return nil, errors.New("About cannot exceed 500 characters")
	}
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
		About:    about,
	}, nil
}

// SetName sets the name of the user with validation.
func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New("Name cannot be blank")
	}
	u.Name = name
	return nil
}

// SetAbout sets the 'about' information of the user with validation.
func (u *User) SetAbout(about string) error {
	if len(about) > 500 {
		return errors.New("About cannot exceed 500 characters")
	}
	u.About = about
	return nil
}

// GetID returns the ID of the user.
func (u *User) GetID() int {
	return u.ID
}

// SetID sets the ID of the user.
func (u *User) SetID(id int) {
	u.ID = id
}

// GetName returns the name of the user.
func (u *User) GetName() string {
	return u.Name
}

// SetEmail sets the email of the user.
func (u *User) SetEmail(email string) {
	u.Email = email
}

// GetEmail returns the email of the user.
func (u *User) GetEmail() string {
	return u.Email
}

// GetPassword returns the password of the user.
func (u *User) GetPassword() string {
	return u.Password
}

// SetPassword sets the password of the user.
func (u *User) SetPassword(password string) {
	u.Password = password
}

// GetRole returns the role of the user.
func (u *User) GetRole() string {
	return u.Role
}

// SetRole sets the role of the user.
func (u *User) SetRole(role string) {
	u.Role = role
}

// GetAbout returns the 'about' information of the user.
func (u *User) GetAbout() string {
	return u.About
}
