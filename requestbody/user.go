package requestbody

type CreateUser struct {
	Name        string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateUser struct {
	Name        string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
}
