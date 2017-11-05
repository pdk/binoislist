package models

//go:generate $GOPATH/src/github.com/pdk/binoislist/make-binois-pointer-list.sh user_list.go models User UserList

// User represents a person who might login.
type User struct {
	Name  string
	Email string
}

// NewUser creates a new user
func NewUser(name, email string) *User {

	return &User{
		Name:  name,
		Email: email,
	}
}

func (user *User) Equals(other *User) bool {
	if user == other {
		return true
	}

	if user == nil || other == nil {
		return false
	}

	return user.Name == other.Name && user.Email == other.Email
}
