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

	return result
}

func (thisState State) Configure(creator MessageCreator, queue Queue, workers []Worker) {
	thisState.MessageCreator = creator
	thisState.Queue = queue
	thisState.Workers = workers
	thisState.Name = thisState.GetName()
}

func (thisState State) CheckWorkersBusy() (bool, []bool) {
	var busyValues []bool
	busy := true
	for _, worker := range thisState.Workers {
		busyValues = append(busyValues, worker.IsBusy)
		if !worker.IsBusy { busy = false }
	}

	return busy, busyValues
}

func (thisState State) SetParent(parent State) {
	thisState.Parent = &parent
}

func (thisState State) AddChildren(children []State) {
	for _, child := range children {
		thisState.Children = append(thisState.Children, &child)
	}
}


