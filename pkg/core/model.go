package core

import "time"

// Model specifies the behavior of a data object.
type Model interface {
	// Prepare the initial state of the data object.
	Prepare(state interface{})

	// Initial state of the data object.
	Init() interface{}

	// Step function for the data object. Returns whether or not the system
	// could take this step with the given inputs and outputs and also
	// returns the new state. This should not mutate the existing state.
	//
	// state must support encoding to and decoding from json.
	Step(state interface{}, input interface{}, output interface{}) (bool, interface{})

	// Equality on states.
	Equal(state1, state2 interface{}) bool

	// Name returns the unique name for the model.
	Name() string
}

// Operation action
const (
	InvokeOperation = "call"
	ReturnOperation = "return"
	InvokeNemesis   = "inject"
	RecoverNemesis  = "recover"
)

// Operation of a data object.
type Operation struct {
	Action string      `json:"action"`
	Proc   int64       `json:"proc"`
	Data   interface{} `json:"data"`
	Time   time.Time   `json:"time"`
}
