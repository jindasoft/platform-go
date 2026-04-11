package xenums

type UserSource string

const (
	UserSourceApp     UserSource = "app"
	UserSourceWeb     UserSource = "web"
	UserSourceDefault UserSource = UserSourceWeb
)

func (u UserSource) IsValid() bool {
	switch u {
	case UserSourceApp, UserSourceWeb:
		return true
	}
	return false
}

func (u UserSource) IsValidOrDefault() UserSource {
	if u.IsValid() {
		return u
	}
	return UserSourceDefault
}

func (u UserSource) String() string {
	return string(u)
}

func (u UserSource) IsApp() bool {
	return u == UserSourceApp
}

func (u UserSource) IsWeb() bool {
	return u == UserSourceWeb
}
