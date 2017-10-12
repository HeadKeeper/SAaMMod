package types

const (
	MESSAGE_CREATOR__SEND             = "SEND"
	MESSAGE_CREATOR__NOT_SEND         = "NOT_SEND"
	MESSAGE_CREATOR__PROBABILITY_SEND = "PROBABILITY_SEND"
)

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

func (thisCreator MessageCreator) WillSendMessageNextStep() string {
	if thisCreator.ProbabilityToCreateMessage >= 1 {
		if thisCreator.RemainStepsForNewMessage == 1 {
			return MESSAGE_CREATOR__SEND
		} else {
			return MESSAGE_CREATOR__NOT_SEND
		}
	} else {
		return MESSAGE_CREATOR__PROBABILITY_SEND
	}
}

func (thisCreator MessageCreator) SendMessage() MessageCreator {
	thisCreator.RemainStepsForNewMessage = thisCreator.StepsForNewMessage
	return thisCreator
}

func (thisCreator MessageCreator) MakeStep() MessageCreator {
	thisCreator.RemainStepsForNewMessage -= 1
	return thisCreator
}

func (thisCreator MessageCreator) Sent() bool {
	return thisCreator.RemainStepsForNewMessage == thisCreator.StepsForNewMessage
}

func (thisCreator MessageCreator) CreateCopy() MessageCreator {
	var copied MessageCreator
	copied.StepsForNewMessage = thisCreator.StepsForNewMessage
	copied.RemainStepsForNewMessage = thisCreator.RemainStepsForNewMessage
	copied.ProbabilityToCreateMessage = thisCreator.ProbabilityToCreateMessage
	return copied
}
