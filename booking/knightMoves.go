package booking

type Location struct {
	X int
	Y int
}

func KnightMove(x, y int) int {
	dirX := []int{-2, -2, -1, -1, 1, 1, 2, 2}
	dirY := []int{1, -1, 2, -2, 2, -2, 1, -1}
	count := 0
	visited := make(map[Location]bool)
	loc := Location{
		X: 0,
		Y: 0,
	}
	queue := make([]Location, 0)
	queue = append(queue, loc)
	visited[loc] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[0]
			queue = queue[1:]
			if item.X == x && item.Y == y {
				return count
			}
			for i := range dirX {
				newX := item.X + dirX[i]
				newY := item.Y + dirY[i]
				newLoc := Location{
					X: newX,
					Y: newY,
				}
				if !visited[newLoc] && newX > 0 && newY > 0 {
					visited[newLoc] = true
					queue = append(queue, newLoc)
				}
			}
		}
		count++
	}
	return -1
}
