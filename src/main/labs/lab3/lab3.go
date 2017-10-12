package lab3

import (
	"main/types"
	"fmt"
)

const (
	WORKERS_TYPE__PARALLEL = "PARALLEL"
	WORKERS_TYPE__INLINE = "INLINE"
)

var (
	testInitialState = types.State {
		Queue: types.Queue {
			Size: 1,
			AmountOfElementInside: 0,
		},
		MessageCreator: types.MessageCreator {
			ProbabilityToCreateMessage: 1,
			StepsForNewMessage: 2,
			RemainStepsForNewMessage: 2,
		},
		Workers: []types.Worker {
			{
				ProbabilityToPerform: 1,
				IsBusy: false,
			},
			{
				ProbabilityToPerform: 1,
				IsBusy: false,
			},
		},
	}

	initialState types.State = testInitialState.CreateCopy()
	stepsAmount int = 4
	workersType string = WORKERS_TYPE__PARALLEL
)

func Setup() {
	fmt.Println("Enter probability and steps amount for message creator :")
	fmt.Printf("\tEnter probability (0.1 - 1): ")
	fmt.Scanf("%f", &initialState.MessageCreator.ProbabilityToCreateMessage)
	fmt.Printf("\tEnter steps: ")
	fmt.Scanf("%d", &initialState.MessageCreator.StepsForNewMessage)
	initialState.MessageCreator.RemainStepsForNewMessage = initialState.MessageCreator.StepsForNewMessage
	fmt.Println()

	fmt.Printf("Enter size of queue: ")
	fmt.Scanf("%d", &initialState.Queue.Size)
	initialState.Queue.AmountOfElementInside = 0
	fmt.Println()

	fmt.Printf("Enter amount of workers ( > 0): ")
	var workersSize int
	fmt.Scanf("%d", &workersSize)
	for index := 0; index < workersSize; index ++ {
		var worker types.Worker
		fmt.Printf("\tEnter probability for worker %d (0.1 - 1): ", index + 1)
		fmt.Scanf("%f", &worker.ProbabilityToPerform)
		worker.IsBusy = false
		initialState.Workers = append(initialState.Workers, worker)
	}
	fmt.Println()

	fmt.Printf("Enter workers type (" + WORKERS_TYPE__PARALLEL + ", " + WORKERS_TYPE__INLINE + "): ")
	fmt.Scanf("%v", &workersType)
	fmt.Println()

	fmt.Printf("Enter amount of steps: ")
	fmt.Scanf("%d", &stepsAmount)
	fmt.Println()

	fmt.Println("Setup is complete")
	fmt.Println()
}

func Perform() {

	var states []types.State

	states = append(states, initialState)

	for step := 0; step < stepsAmount; step++ {
		fmt.Printf("\nSTART {\n\n")
		newStates := NextStep(states)

		//newStates = removeFullEquals(newStates)
		//newStates = removeImpossibleStates(newStates)

		states = append(states, newStates...)
		fmt.Printf("}\n")
	}

	//states = filterStates(states)
	showStates(states)
}

func NextStep(previousStates []types.State) []types.State {
	var newStates []types.State

	for _, state := range previousStates {
		switch state.MessageCreator.WillSendMessageNextStep() {
		case types.MESSAGE_CREATOR__SEND:
			newStates = append(newStates, analyzeWillSend(state)...)
			break
		case types.MESSAGE_CREATOR__NOT_SEND:
			newStates = append(newStates, analyzeWillNotSend(state)...)
			break
		case types.MESSAGE_CREATOR__PROBABILITY_SEND:
			newStates = append(newStates, analyzeSendWithProbability(state)...)
			break
		default:
			break
		}
	}

	return newStates
}


func analyzeWillSend(previousState types.State) []types.State  {
	newStateBase := previousState.CreateCopy()
	newStateBase = newStateBase.SetParent(previousState)

	var newStates []types.State

	newStateBase.MessageCreator = newStateBase.MessageCreator.SendMessage()
	newStatesWithWorkers := append(newStates, sendMessageToWorkers(newStateBase, true)...)
	for _, newState := range newStatesWithWorkers {
		newStates = append(newStates, sendMessageToQueue(newState)...)
	}

	if len(newStates) == 0 { newStates = append(newStates, newStatesWithWorkers...) }
	previousState = previousState.AddChildren(newStates)
	return newStates
}

func analyzeWillNotSend(previousState types.State) []types.State  {
	newStateBase := previousState.CreateCopy()
	newStateBase = newStateBase.SetParent(previousState)

	var newStates []types.State

	newStateBase.MessageCreator = newStateBase.MessageCreator.MakeStep()

	newStatesWithWorkers := append(newStates, sendMessageToWorkers(newStateBase, false)...)
	for _, newState := range newStatesWithWorkers {
		newStates = append(newStates, popMessageFromQueue(newState)...)
	}

	if len(newStates) == 0 { newStates = append(newStates, newStateBase) }
	previousState = previousState.AddChildren(newStates)
	return newStates
}

func analyzeSendWithProbability(previousState types.State) []types.State  {
	return []types.State { }
}


func sendMessageToQueue(state types.State) []types.State {
	return append([]types.State {}, state.Queue.MakeStepPush(state)...)
}

func popMessageFromQueue(state types.State) []types.State {
	return append([]types.State {}, state.Queue.MakeStepPop(state)...)
}

func sendMessageToWorkers(state types.State, messageSent bool) []types.State {
	var newStates []types.State

	switch workersType {
	case WORKERS_TYPE__PARALLEL:
		for workerNumber, worker := range state.Workers {
			newStates = append(newStates, worker.MakeStep(state, workerNumber, messageSent)...)
		}
		break
	case WORKERS_TYPE__INLINE:
		break
	default:
		break
	}

	return newStates
}


func showStates(states []types.State) {
	for _, state := range states {
		fmt.Printf("\tSTATE " + state.GetName() + "\n")
	}
}

func filterStates(states []types.State) []types.State {
	/*for idx := range states {
		for idx2 := range states {
			if states[idx].GetName() == states[idx2].GetName() {
				states[idx2].AlreadyCreated = true
			}
		}
	}*/
	states = removeFullEquals(states)
	states = markDuplicates(states)
	return states
}

func removeFullEquals(elements []types.State) []types.State {
	var values []types.State
	for i := 0; i < len(elements); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if elements[v].FullEquals(elements[i]) {
				exists = true
				break
			}
		}
		if !exists {
			values = append(values, elements[i])
		}
	}
	return values
}

func markDuplicates(elements []types.State) []types.State {
	for i := 0; i < len(elements); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if elements[v].Equals(elements[i]) {
				exists = true
				break
			}
		}
		if exists {
			elements[i].AlreadyCreated = true
		}
	}
	return elements
}

func removeImpossibleStates(states []types.State) []types.State {
	var newStates []types.State
	for _, state := range states {
		impossible := false
		if state.AnyWorkerFree() && state.Queue.CanPop() { impossible = true }

		if state.MessageCreator.Sent() && state.WorkersBusy() && state.Queue.CanPush() { impossible = true }

		if !impossible {
			newStates = append(newStates, state)
		}
	}
	return newStates
}
