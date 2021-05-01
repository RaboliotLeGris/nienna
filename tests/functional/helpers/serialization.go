package helpers

type UserLogin struct {
	Username string `json:"username"`
}

type UserRegister struct {
	Username string `json:"username"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Video struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Uploader    User   `json:"uploader"`
}

type StatusSerialized struct {
	Status string `json:"status"`
}
