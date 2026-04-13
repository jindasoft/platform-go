type DayOfWeek string

const (
	Sunday    DayOfWeek = "sunday"
	Monday    DayOfWeek = "monday"
	Tuesday   DayOfWeek = "tuesday"
	Wednesday DayOfWeek = "wednesday"
	Thursday  DayOfWeek = "thursday"
	Friday    DayOfWeek = "friday"
	Saturday  DayOfWeek = "saturday"
)

func FromTimeWeekday(w time.Weekday) DayOfWeek {
	days := map[time.Weekday]DayOfWeek{
		time.Sunday:    Sunday,
		time.Monday:    Monday,
		time.Tuesday:   Tuesday,
		time.Wednesday: Wednesday,
		time.Thursday:  Thursday,
		time.Friday:    Friday,
		time.Saturday:  Saturday,
	}
	return days[w]
}
