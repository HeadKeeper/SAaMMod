package types

type Queue struct {
	Size int
	AmountOfElementInside int
}

func (thisQueue Queue) Equals(queue Queue) bool {
	sizeEqual := false
	if thisQueue.Size == queue.Size {
		sizeEqual = true
	}
	amountInsideEqual := false
	if thisQueue.AmountOfElementInside == queue.AmountOfElementInside {
		amountInsideEqual = true
	}
	return sizeEqual && amountInsideEqual
}

func (thisQueue Queue) CanPush() bool {
	return thisQueue.AmountOfElementInside < thisQueue.Size
}

func (thisQueue Queue) CanPop() bool {
	return thisQueue.AmountOfElementInside > 0
}


func (thisQueue Queue) Push() Queue {
	thisQueue.AmountOfElementInside += 1
	return thisQueue
}

func (thisQueue Queue) Pop() Queue {
	if thisQueue.AmountOfElementInside > 0 {
		thisQueue.AmountOfElementInside -= 1
	}
	return thisQueue
}

func (thisQueue Queue) Clear() Queue {
	thisQueue.AmountOfElementInside = 0
	return thisQueue
}

func (thisQueue Queue) MakeStepPop(state State) []State {
	var newStates []State
	if state.WorkersBusy() {
		return newStates
	}

	if state.Queue.CanPop() {
		newState := state.CreateCopy()
		newState.Queue = newState.Queue.Push()
		newStates = append(newStates, newState)
	}

	return newStates
}

func (thisQueue Queue) MakeStepPush(state State) []State {
	var newStates []State
	if state.AnyWorkerFree() {
		return newStates
	}

	if state.Queue.CanPush() {
		newState := state.CreateCopy()
		newState.Queue = newState.Queue.Push()
		newStates = append(newStates, newState)
	}

	return newStates
}

func (thisQueue Queue) CreateCopy() Queue {
	var copied Queue
	copied.AmountOfElementInside = thisQueue.AmountOfElementInside
	copied.Size = thisQueue.Size
	return copied
}