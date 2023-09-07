package models

type StatusType string

const (
	Online StatusType = "ONLINE"
	Offline StatusType = "OFFLINE"
	Idle StatusType = "IDLE"
	Busy StatusType = "BUSY"
)

// Represents a payload resetting a user's password on Oprish.
type CreatePasswordRestCode struct {
	// The user's email.
	Email string `json:"email"`
}

// Represents user-related stuff deletion on Oprish.
type PasswordDeleteCredentials struct {
	// The user's password.
	Password string `json:"password"`
}

// Represents a password resetting payload on Oprish.
type RestPassword struct {
	// The password reset code the user got emailed.
	Code int `json:"code"`
	// The user's email.
	Email string `json:"email"`
	// The user's new password.
	Password string `json:"password"`
}

// A user's status.
type Status struct {
	// The status' type.
	Type StatusType `json:"type"`
	// The (optional) status' text.
	Text *string `json:"text"`
}


// Represents the user update payload on Oprish.
type UpdateUser struct {
	// The user's password.
	Password string `json:"password"`
	// The user's new (optional) username.
	Username *string `json:"username"`
	// The user's new (optional) email.
	Email *string `json:"email"`
	// The user's new (optional) password.
	NewPassword *string `json:"new_password"`
}

// Represents the profile update payload on Oprish.
type UpdateProfile struct {
	// The user's new (optional) display name.
	DisplayName *string `json:"display_name"`
	// The user's new (optional) status.
	Status *string `json:"status"`
	// The user's new (optional) status type.
	StatusType *StatusType `json:"status_type"`
	// The user's new (optional) bio.
	Bio *string `json:"bio"`
	// The user's new (optional) avatar.
	Avatar *int `json:"avatar"`
	// The user's new (optional) banner.
	Banner *int `json:"banner"`
}

// Represents the user payload on Oprish.
type User struct {
	// The user's ID.
	Id int `json:"id"`
	// The user's username.
	Username string `json:"username"`
	// The user's (optional) display name.
	DisplayName *string `json:"display_name"`
	// The user's social score credit.
	SocialCredit *int `json:"social_credit"`
	// The user's status.
	Status Status `json:"status"`
	// The user's (optional) bio.
	Bio *string `json:"bio"`
	// The user's (optional) avatar.
	Avatar *int `json:"avatar"`
	// The user's (optional) banner.
	Banner *int `json:"banner"`
	// The user's badges as a bitfield.
	Badges *int `json:"badges"`
	// The user's instance-wide permissions as a bitfield.
	Permissions int `json:"permissions"`
	// The user's email.
	Email *string `json:"email"`
	// The user's verification status.
	Verified bool `json:"verified"`
}

// Represents the user creation payload on Oprish.
type UserCreate struct {
	// The user's name.
	Username string `json:"username"`
	// The user's email.
	Email string `json:"email"`
	// The user's password.
	Password string `json:"password"`
}
