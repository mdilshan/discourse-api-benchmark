package discourse

type CreateDiscourseUserDto struct {
	IamId    string `json:"iam_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type CreateDiscourseUserBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type CreateDiscourseUserResponse struct {
	Success bool   `json:"success"`
	Active  bool   `json:"active"`
	Message string `json:"message"`
	User_id int    `json:"user_id"`
}

type DiscourseApiError struct {
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	Errors       interface{} `json:"errors"`
	Value        interface{} `json:"value"`
	Is_developer bool        `json:"is_developer"`
}
