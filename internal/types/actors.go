package types

type UserMetadata struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Companies   []string `json:"companies"`
	Department  []string `json:"departments"`
	Role        []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type User struct {
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Companies   []string `json:"companies"`
	Department  []string `json:"departments"`
	Role        []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func (user *User) CreateUser() User {
	
}
