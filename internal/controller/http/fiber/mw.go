package fiber

type UserRole string

const (
	UserRoleNone  UserRole = ""
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

func (ur UserRole) IsValid() bool {
	switch ur {
	case UserRoleNone, UserRoleAdmin, UserRoleUser:
		return true
	}
	return false
}
