package responses

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeInvalidParam = 20003 // Invalid email
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeInvalidParam: "Invalid email",
}
