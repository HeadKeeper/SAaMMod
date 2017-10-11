package types

type State struct {
	MessageCreator MessageCreator
	Queue Queue
	Worker Worker

	Parent *State
	Children []*State
}

func (thisState State) Equals(state State) bool {
	equalWorker := false
	if thisState.Worker.Equals(state.Worker) {
		equalWorker = true
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