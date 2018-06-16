package events

func (bt ByTimestamp) Len() int {
	return len(bt)
}

func (bt ByTimestamp) Swap(i, j int) {
	bt[i], bt[j] = bt[j], bt[i]
}

func (bt ByTimestamp) Less(i, j int) bool {
	return bt[i].Timestamp.Before(bt[j].Timestamp)
}
