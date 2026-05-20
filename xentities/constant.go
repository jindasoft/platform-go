package xentities

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
)

var StatusMap = map[string]Status{
	string(StatusActive):   StatusActive,
	string(StatusInactive): StatusInactive,
	string(StatusDeleted):  StatusDeleted,
}

func FromStatusCode(code string) (Status, bool) {
	status, ok := StatusMap[code]
	return status, ok
}

func (s Status) IsValid() bool {
	_, ok := StatusMap[string(s)]
	return ok
}

func (s Status) String() string {
	return string(s)
}
