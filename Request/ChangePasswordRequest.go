package Request

type ChangePasswordRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
}
