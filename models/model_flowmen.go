package models

type FlowmenProduceYaml struct {
	Version int
	Nodes   []FlowmenNode
}

type FlowmenNode struct {
	In          *string
	Name        string
	Type        string
	Description string
	Params      map[string]string
}

const (
	CIOS_WEBSOCKET = "websocket_cios"
	DROP           = "drop"
	PROGRAM        = "program"
	THROUGH        = "through"
	TIMER          = "timer"
	WEBHOOK        = "webhook"
	FLOWMEN_IMAGE  = ""
)
