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

func (thisWorker Worker) MakeStep(state State) []State {
	var nextStates []State

	currentState := State {
		Workers: []Worker { thisWorker },
	}

	nextStates = append(nextStates, currentState)

	if thisWorker.IsBusy {
		stateWillNotPerform := State {
			Workers: []Worker {{
				IsBusy:               true,
				ProbabilityToPerform: thisWorker.ProbabilityToPerform,
			},
		}}
		nextStates = append(nextStates, currentState, stateWillNotPerform)
	}

	return nextStates
}