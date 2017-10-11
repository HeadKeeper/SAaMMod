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

func (thisQueue Queue) Push() {
	thisQueue.AmountOfElementInside += 1
}

func (thisQueue Queue) Pop() {
	if thisQueue.AmountOfElementInside > 0 {
		thisQueue.AmountOfElementInside -= 1
	}
}

func (thisQueue Queue) Clear()  {
	thisQueue.AmountOfElementInside = 0
}