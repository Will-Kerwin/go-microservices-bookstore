package events

type CreateUserEvent struct {
	ID        string `json:"_id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type UpdateUserEventData struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type UpdateUserEvent struct {
	ID   string              `json:"id"`
	Data UpdateUserEventData `json:"data"`
}
