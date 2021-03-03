package utils

import "os"

const (
	StageStr = "STAGE"
)

func GetStage() string            { return os.Getenv(StageStr) }
func SetStage(stage string) error { return os.Setenv(StageStr, stage) }
