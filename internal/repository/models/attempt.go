package models

type AttemptStatus int8

const (
	AttemptStatusComplete         AttemptStatus = iota
	AttemptStatusCompilationError AttemptStatus = iota
	AttemptStatusRunningError AttemptStatus = iota
	AttemptStatusOutOfMemory      AttemptStatus = iota
	AttemptStatusTimeout          AttemptStatus = iota
	AttemptStatusOutputOverflow AttemptStatus = iota
)

type TestStatus struct {
	TestId        int64         `json:"test_id"`
	Status        AttemptStatus `json:"status"`
	Output        string        `json:"output"`
	ExecutionTime int64         `json:"execution_time"`
	MemoryUsage   int64         `json:"memory_usage"`
}

type Attempt struct {
	Id         int64
	TaskId     int64
	Status     AttemptStatus
	LanguageId int
	Code       string
	Tests      []TestStatus
}
