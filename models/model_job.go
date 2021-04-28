package models

type (
	Job           map[string]MessagingJobs
	MessagingJobs []MessagingJob
	MessagingJob  struct {
		Value string `yaml:"value,flow"`
	}
)
