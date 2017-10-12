package types

type Worker struct {
	IsBusy bool
	ProbabilityToPerform float32
}

func (thisWorker Worker) Equals(worker Worker) bool {
	busyEqual := false
	if thisWorker.IsBusy == worker.IsBusy {
		busyEqual = true
	}
	probabilityEqual := false
	if thisWorker.ProbabilityToPerform == worker.ProbabilityToPerform {
		probabilityEqual = true
	}
	return busyEqual && probabilityEqual
}

func (thisWorker Worker) MakeStep(state State, workerNumber int, messageSent bool) []State {
	var nextStates []State

	if thisWorker.IsBusy {
		stateWorkerFree := state.CreateCopy()
		stateWorkerFree.Workers[workerNumber] = Worker {
			ProbabilityToPerform: thisWorker.ProbabilityToPerform,
			IsBusy: false,
		}
		nextStates = append(nextStates, stateWorkerFree)
	}

	if messageSent && !thisWorker.IsBusy {
		stateWorkerBusy := state.CreateCopy()
		stateWorkerBusy.Workers[workerNumber].IsBusy = true
		nextStates = append(nextStates, stateWorkerBusy)
	}

	stateWorker := state.CreateCopy()
	stateWorker.Workers[workerNumber] = thisWorker
	nextStates = append(nextStates, stateWorker)

	return nextStates
}

func (thisWorker Worker) CreateCopy() Worker {
	var copied Worker
	copied.ProbabilityToPerform = thisWorker.ProbabilityToPerform
	copied.IsBusy = thisWorker.IsBusy
	return copied
}