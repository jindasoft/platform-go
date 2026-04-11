package xenums

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleTraveler   Role = "traveler"
	RoleInfluencer Role = "influencer"
	RoleDefault    Role = RoleTraveler
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleTraveler, RoleInfluencer:
		return true
	}
	return false
}

func (r Role) IsValidOrDefault() Role {
	if r.IsValid() {
		return r
	}
	return RoleDefault
}

func (r Role) String() string {
	return string(r)
}

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsTraveler() bool {
	return r == RoleTraveler
}

func (r Role) IsInfluencer() bool {
	return r == RoleInfluencer
}
