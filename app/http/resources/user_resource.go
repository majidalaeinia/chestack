package resources

import "github.com/volatiletech/null/v8"

// User Represents the fields of the users table that must be displayed
type User struct {
	ID              uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Email           string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Mobile          string      `boil:"mobile" json:"mobile" toml:"mobile" yaml:"mobile"`
	EmailVerifiedAt null.Time   `boil:"email_verified_at" json:"-" toml:"email_verified_at" yaml:"email_verified_at,omitempty"`
	Password        string      `boil:"password" json:"-" toml:"password" yaml:"password"`
	RememberToken   null.String `boil:"remember_token" json:"-" toml:"remember_token" yaml:"remember_token,omitempty"`
	CreatedAt       null.Time   `boil:"created_at" json:"-" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt       null.Time   `boil:"updated_at" json:"-" toml:"updated_at" yaml:"updated_at,omitempty"`
}
