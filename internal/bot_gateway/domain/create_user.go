package domain

type CreateUserCommand struct {
	UserID   int64
	Username string
	UserFname string
	UserSname string
}
