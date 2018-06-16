package common

type HoursResponse []DurationSummary

type DurationSummary struct {
	CaseId   int     `json:"case_id"`
	Duration float64 `json:"duration"`
}

type Assignment struct {
	Assignee string
	Team     string
}

type Status struct {
	Name string
}
