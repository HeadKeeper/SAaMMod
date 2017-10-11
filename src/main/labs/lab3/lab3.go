package lab3

import (
	"main/types"
	"fmt"
)

var (
	initialState types.State
	stepsAmount int
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
		states = append(states, NextStep(states)...)
	}

	states = filterStates(states)
	showStates(states)
}

func NextStep(previousStates []types.State) []types.State {
	var newStates []types.State

	for _, state := range previousStates {
		switch state.MessageCreator.WillSendMessageNextStep() {
		case 1:
			fmt.Println("1 ", state.GetName())
			newStates = append(newStates, analyzeWillSend(state)...)
			break
		case 0:
			fmt.Println("0 ", state.GetName())
			newStates = append(newStates, analyzeWillNotSend(state)...)
			break
		case -1:
			fmt.Println("-1 ", state.GetName())
			newStates = append(newStates, analyzeSendWithProbability(state)...)
			break
		default:
			break
		}
	}

	return newStates
}

func analyzeWillSend(previousState types.State) []types.State  {
	newStateBase := previousState
	newStateBase.SetParent(previousState)

	var newStates []types.State

	newStateBase.MessageCreator.SendMessage()
	newStates = append(newStates, sendMessageToQueue(newStateBase)...)
	newStates = append(newStates, sendMessageToWorkers(newStateBase)...)

	previousState.AddChildren(newStates)
	return newStates
}

func analyzeWillNotSend(previousState types.State) []types.State  {
	newStateBase := previousState
	newStateBase.SetParent(previousState)

	var newStates []types.State

	newStateBase.MessageCreator.MakeStep()

	newStatesWithQueue := append(newStates, sendMessageToQueue(newStateBase)...)
	for _, newState := range newStatesWithQueue {
		for _, worker := range newState.Workers {
			newStates = append(newStates, worker.MakeStep(newStateBase)...)
		}
	}

	previousState.AddChildren(newStates)
	return newStates
}

func analyzeSendWithProbability(previousState types.State) []types.State  {
	return []types.State { }
}

func sendMessageToQueue(state types.State) []types.State {
	return []types.State { }
}

func sendMessageToWorkers(state types.State) []types.State {
	return []types.State { }
}

func showStates(states []types.State) {

}

func filterStates(states []types.State) []types.State {
	return states
}