package models

type (
	Job               map[string]MessagingJobs
	MessagingJobs     []MessagingJob
	MessagingJobValue struct {
		Data      string `yaml:"data"`
		Timestamp int64  `yaml:"timestamp"`
	}
	MessagingJob struct {
		Channel string              `yaml:"channel"`
		Loop    int                 `yaml:"loop"`
		Plugin  string              `yaml:"plugin"`
		Value   []MessagingJobValue `yaml:"value,flow"`
	}
)
