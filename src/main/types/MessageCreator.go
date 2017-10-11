package types

type MessageCreator struct {
	RemainStepsForNewMessage int

	StepsForNewMessage int
	ProbabilityToCreateMessage float32
}

func (thisCreator MessageCreator) Equals(creator MessageCreator) bool {
	remainStepsEqual := false
	if thisCreator.RemainStepsForNewMessage == creator.RemainStepsForNewMessage {
		remainStepsEqual = true
	}
	stepsEqual := false
	if thisCreator.StepsForNewMessage == creator.StepsForNewMessage {
		stepsEqual = true
	}
	probabilityEqual := false
	if thisCreator.ProbabilityToCreateMessage == creator.ProbabilityToCreateMessage {
		probabilityEqual = true
	}
	return remainStepsEqual && stepsEqual && probabilityEqual
}

func (thisCreator MessageCreator) WillSendMessageNextStep() int {
	if thisCreator.ProbabilityToCreateMessage >= 1 {
		if thisCreator.RemainStepsForNewMessage == 1 {
			return 1
		} else {
			return 0
		}
	} else {
		return -1
	}
}

func (thisCreator MessageCreator) SendMessage()  {
	thisCreator.RemainStepsForNewMessage = thisCreator.StepsForNewMessage
}
