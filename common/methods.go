package common

func (r HoursResponse) Increment(summary DurationSummary) HoursResponse {
	for idx, stored := range r {
		if summary.CaseId == stored.CaseId {
			r[idx].Duration += summary.Duration
			return r
		}
	}
	return append(r, summary)
}
