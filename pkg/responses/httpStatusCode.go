package responses

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeInvalidParam = 20003 // Invalid email
	ErrCodeInvalidToken = 30001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeInvalidParam: "Invalid email",
	ErrCodeInvalidToken: "Invalid token",
}
