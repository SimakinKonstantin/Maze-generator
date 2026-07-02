package solver

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"

type priorityQueueElem struct {
	Cell     cell.Cell
	Priority int
}

type priorityQueue struct {
	queue []priorityQueueElem
}

func (pq *priorityQueue) Pop() (priorityQueueElem, error) {
	if len(pq.queue) == 0 {
		return priorityQueueElem{
				Cell:     cell.Cell{X: -1, Y: -1, T: cell.Wall},
				Priority: -1}, PriorityQueueError{"ошибка" +
				"при попытке получить элемент из пустой очереди"}
	}

	minValue := pq.queue[0].Priority
	minIndex := 0

	for i, v := range pq.queue {
		if v.Priority < minValue {
			minValue = v.Priority
			minIndex = i
		}
	}

	toReturn := pq.queue[minIndex]
	pq.queue = append(pq.queue[:minIndex], pq.queue[minIndex+1:]...)

	return toReturn, nil
}

func (pq *priorityQueue) Push(elem priorityQueueElem) {
	pq.queue = append(pq.queue, elem)
}

func (pq *priorityQueue) Len() int {
	return len(pq.queue)
}
