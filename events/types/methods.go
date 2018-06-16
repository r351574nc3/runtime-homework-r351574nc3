package types

func (i *SearchableDurationIndex) New() *SearchableDurationIndex {
	if i == nil {
		retval := new(SearchableDurationIndex)
		retval.Items = make([]StateEventSummary, 10)
		return retval
	}
	return i
}

func (i *SearchableDurationIndex) Filter(f DurationIndexSearchFilter) []StateEventSummary {
	var retval []StateEventSummary = make([]StateEventSummary, 10)
	for _, summary := range i.Items {
		if f(summary) {
			retval = append(retval, summary)
		}
	}
	return retval
}

// AssignTo copies values from the source Event to the Target event as long as the fields are not
// populated. Returns true if all fields are populated
func (source *Event) AssignTo(target *Event) bool {
	if source.Id != target.Id {
		return false
	}

	if source.Assignee != "" && target.Assignee == "" {
		target.Assignee = source.Assignee
	}

	if source.Team != "" && target.Team == "" {
		target.Team = source.Team
	}

	if source.State != nil && source.State.From != "" && target.State.From == "" {
		target.State.From = source.State.From
	}

	if source.State != nil && source.State.To != "" && target.State.To == "" {
		target.State.To = source.State.To
	}

	if target.Assignee != "" && target.Team != "" && target.State.From != "" && target.State.To != "" {
		return true
	}
	return true
}
