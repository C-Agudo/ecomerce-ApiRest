package domain

// User model structure
type User struct {
	Id int64
	Name string
	Password string
	Role string
}

//Create user command
type CreateUserCMD struct {
	Name string 'json:"name"'
	Password string 'json: "password"'
	Role string 'json:"role"'
}