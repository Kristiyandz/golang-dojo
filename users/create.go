package users

type user struct {
	id        string
	firstName string
	lastName  string
	email     string
	units     []string
}

func CreateUser() []user {

	createdUsers := make([]user, 0)
	createdUser := user{id: "123", firstName: "John", lastName: "Doe", email: "test@test.com", units: make([]string, 0)}

	s := append(createdUsers, createdUser)
	return s
}
