package main

import "fmt"

var maze = [6][5]int{
	{0, 1, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{1, 1, 1, 0, 0},
	{0, 1, 0, 0, 1},
	{0, 1, 0, 0, 0},
}

type coordinate struct {
	x int
	y int
}

var start, end = coordinate{
	x: 0,
	y: 0,
}, coordinate{
	x: 5,
	y: 4,
}

// 操作数组
var actions = [4]coordinate{
	{x: -1, y: 0},
	{x: 0, y: -1},
	{x: 1, y: 0},
	{x: 0, y: 1},
}

func add(c1 coordinate, c2 coordinate) coordinate {
	return coordinate{c1.x + c2.x, c1.y + c2.y}
}

// 0未访问 1访问
var visits = [6][5]int{}
var counts = [6][5]int{}

func work(start, end coordinate) {

	Queue := []coordinate{start}
	visits[start.x][start.y] = 1
	for len(Queue) > 0 {
		if Queue[0] == end {
			fmt.Println("已经到达终点")
			break
		}
		curStep := Queue[0]
		fmt.Println("cur step x: ", curStep.x, " y: ", curStep.y)
		Queue = Queue[1:]
		for _, action := range actions {
			nextStep := add(curStep, action)
			if nextStep.x >= 6 || nextStep.y >= 5 || nextStep.x < 0 || nextStep.y < 0 {
				continue
			}
			// 表示没访问过
			if visits[nextStep.x][nextStep.y] != 1 && maze[nextStep.x][nextStep.y] != 1 {
				Queue = append(Queue, nextStep)
				visits[curStep.x][curStep.y] = 1
				counts[nextStep.x][nextStep.y] = counts[curStep.x][curStep.y] + 1
			}
		}
	}
}
func main() {
	work(start, end)
	for _, row := range counts {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
