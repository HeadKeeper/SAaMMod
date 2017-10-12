package types

import (
	"strconv"
)

type State struct {
	MessageCreator MessageCreator
	Queue Queue
	Workers []Worker

	Parent *State
	Children []*State
	Name string
	AlreadyCreated bool
}

func (thisState State) FullEquals(state State) bool {
	parentEqual := false
	if thisState.Parent == state.Parent {
		parentEqual = true
	}

	childrenNotEqual := true
	if len(thisState.Children) == len(state.Children) {
		for e := range thisState.Children {
			if thisState.Children[e] != (state.Children[e]) {
				childrenNotEqual = false
			}
		}
	}

	return !childrenNotEqual && parentEqual && thisState.Equals(state)
}

func (thisState State) Equals(state State) bool {
	equalWorker := true
	for _, thisWorker := range thisState.Workers {
		for _, worker := range state.Workers {
			if !thisWorker.Equals(worker) {
				equalWorker = true
			}
		}
	}
	equalQueue := false
	if thisState.Queue.Equals(state.Queue) {
		equalQueue = true
	}
	equalCreator := false
	if thisState.MessageCreator.Equals(state.MessageCreator) {
		equalCreator = true
	}

	return equalCreator && equalQueue && equalWorker
}

func (thisState State) GetName() string {
	var result string

	result += strconv.Itoa(thisState.MessageCreator.RemainStepsForNewMessage)

	result += strconv.Itoa(thisState.Queue.AmountOfElementInside)

	for _, thisWorker := range thisState.Workers {
		if thisWorker.IsBusy {
			result += strconv.Itoa(1)
		} else {
			result += strconv.Itoa(0)
		}
	}

	if thisState.AlreadyCreated {
		result += "_C"
	}

	return result
}

func (thisState State) Configure(creator MessageCreator, queue Queue, workers []Worker) {
	thisState.MessageCreator = creator
	thisState.Queue = queue
	thisState.Workers = workers
	thisState.Name = thisState.GetName()
}

func (thisState State) WorkersBusy() bool {
	var busyValues []bool
	busy := true
	for _, worker := range thisState.Workers {
		busyValues = append(busyValues, worker.IsBusy)
		if !worker.IsBusy { busy = false }
	}

	return busy
}

func (thisState State) AnyWorkerFree() bool {
	busy := true
	for _, worker := range thisState.Workers {
		if !worker.IsBusy {
			busy = false
		}
	}

	return busy == false
}

func (thisState State) SetParent(parent State) State {
	thisState.Parent = &parent
	return thisState
}

func (thisState State) AddChildren(children []State) State {
	for _, child := range children {
		thisState.Children = append(thisState.Children, &child)
	}
	return thisState
}

func (thisState State) CreateCopy() State {
	var copiedState State
	copiedState.MessageCreator = thisState.MessageCreator.CreateCopy()
	copiedState.Queue = thisState.Queue.CreateCopy()
	for _, worker := range thisState.Workers {
		copiedState.Workers = append(copiedState.Workers, worker.CreateCopy())
	}
	copiedState.Parent = thisState.Parent
	for _, child := range copiedState.Children {
		copiedState.Children = append(copiedState.Children, child)
	}
	return copiedState
}


