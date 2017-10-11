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