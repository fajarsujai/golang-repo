package task

type Task struct {
	ID          int
	Title       string
	ActionTime  int64
	IsFinished  bool
	CreatedTime int64
	UpdatedTime int64
	Objectives  []Objective
}

type Objective struct {
	ID            int
	ObjectiveName string
	TaskID        int
	IsFinished    bool
}

// type ObjectList struct {
// 	ObjectName string
// 	TaskID     int
// }

// type Objective struct {
// 	ID          int
// 	ObjectName  []string
// 	IsFinished  int
// 	CreatedTime time.Time
// 	UpdatedTime time.Time
// }
