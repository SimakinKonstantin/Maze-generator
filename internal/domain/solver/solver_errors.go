package solver

// Ошибка при заросе некорректных ячеек для лабиринта.
type ValidityError struct {
	msg string
}

func (validityErr ValidityError) Error() string {
	return validityErr.msg
}

// Ошибка при работе с очередью с приоритетом.
type PriorityQueueError struct {
	msg string
}

func (pqErr PriorityQueueError) Error() string { return pqErr.msg }
