package data

// User struct
type User struct {
		GUID      string `json:"guid"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Gender    string `json:"gender"`
		IPAddress string `json:"ip_address"`
}

// GetUsers func
func GetUsers() []*User {
	return userList
}

var userList = []*User{
	&User{
		GUID: "0e9ffffa-de3e-4378-a1a1-491d59c46425",
		FirstName: "Gabrielle",
		LastName: "Kear",
		Email: "gkear0@trellian.com",
		Gender: "Female",
		IPAddress: "8.79.142.95",
	},
	&User{
		GUID: "70d01dc1-73b6-466d-a6ad-89bc8e2d1cf5",
		FirstName: "Fabian",
		LastName: "Banfill",
		Email: "fbanfill1@woothemes.com",
		Gender: "Male",
		IPAddress: "174.192.211.88",
	},
}