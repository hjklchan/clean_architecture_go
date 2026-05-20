package user_status

type UserStatus int

func DefaultUserStatus() UserStatus {
	return PendingVerification
}

func (us UserStatus) EqualTo(other UserStatus) bool {
	return int(us) == int(other)
}

func (us UserStatus) IsActive() bool {
	return us == Active
}

func (us UserStatus) IsSuspended() bool {
	return us == Suspended
}

func (us UserStatus) DisplayName() string {
	switch us {
	case PendingVerification:
		return "pending verification"
	case Active:
		return "active"
	case Suspended:
		return "suspended"
	case Frozen:
		return "frozen"
	default:
		return "unknown"
	}
}

func (us UserStatus) Available() bool {
	return us == Active
}
