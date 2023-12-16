package task

type CronTaskStatus string

const (
	CronTaskKey                          = "cron_task"
	CronTaskStatusKey                    = "cron_task_%s_status"
	CronTaskStatusFailed  CronTaskStatus = "failed"
	CronTaskStatusSuccess CronTaskStatus = "success"
	CronTaskStatusRunning CronTaskStatus = "running"
)
