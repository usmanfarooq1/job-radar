package query

type Task struct {
	TaskId         string
	TaskType       string
	DelayInSeconds uint32
	SearchKeyword  string
	LocationId     string
	DistanceRadius string
	TaskLocation   string
	CreatedAt      string
	UpdatedAt      string
}
