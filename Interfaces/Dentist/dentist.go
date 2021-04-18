package dentist

type DaysOfWork struct {
	DayWeek    string
	HourStart  string
	HourFinish string
}

type Dentist struct {
	Name              string
	CRO               string
	IdClinic          string
	Password          string
	Email             string
	Description       string
	Phones            []int
	DaysOfWork        []DaysOfWork
	ScheduleIntervals int32
	ScheduleDuration  int32
	IsActive          bool
}
