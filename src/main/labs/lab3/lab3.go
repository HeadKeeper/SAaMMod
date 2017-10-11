package lab3

import (
	"main/types"
	"fmt"
)

var (
	messageCreator types.MessageCreator
	workers []types.Worker
	queue types.Queue
)

func Setup() {
	fmt.Println("Enter probability and steps amount for message creator :")
	fmt.Printf("\tEnter probability (0.1 - 1): ")
	fmt.Scanf("%f", &messageCreator.ProbabilityToCreateMessage)
	fmt.Printf("\tEnter steps: ")
	fmt.Scanf("%d", &messageCreator.StepsForNewMessage)
	messageCreator.RemainStepsForNewMessage = messageCreator.StepsForNewMessage
	fmt.Println()

	fmt.Printf("Enter size of queue: ")
	fmt.Scanf("%d", &queue.Size)
	queue.AmountOfElementInside = 0
	fmt.Println()

	fmt.Printf("Enter amount of workers ( > 0): ")
	var workersSize int
	fmt.Scanf("%d", &workersSize)
	for index := 0; index < workersSize; index ++ {
		var worker types.Worker
		fmt.Printf("\tEnter probability for worker %d (0.1 - 1): ", index + 1)
		fmt.Scanf("%f", &worker.ProbabilityToPerform)
		worker.IsBusy = false
		workers = append(workers, worker)
	}
	fmt.Println()

	fmt.Println("Setup is complete")
}

func Perform() {
	stop := false
	for stop {

		//stop = true
	}
}

func NextStep() []types.State {
	var newStates []types.State

	switch messageCreator.WillSendMessageNextStep() {
	case 1:
		// Will send
		messageCreator.SendMessage()

		break
	case 0:
		// Not send
		for _, worker := range workers {
			//worker.
		}

		break
	case -1:
		// Send with probability

		break
	default:
		break
	}

	return newStates
}