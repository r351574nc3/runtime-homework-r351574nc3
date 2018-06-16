package common

type HoursResponse []DurationSummary

type DurationSummary struct {
	CaseId   int
	Duration int
}

type Assignment struct {
	Assignee string
	Team     string
}

type Status struct {
	Name string
}
