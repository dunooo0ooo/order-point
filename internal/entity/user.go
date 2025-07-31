package entity

type Role string

const (
	RoleClient    Role = "client"
	RoleEmployee  Role = "employee"
	RoleModerator Role = "moderator"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     Role
}
