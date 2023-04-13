package main

import "fmt"

type Tower struct {
	n int
	pegs map[string][]int
}

func (tower *Tower) String() string {
	return fmt.Sprintf("%v", tower.pegs)
}

func NewTower(n int) *Tower {
	var disks []int
	for i := n; i > 0; i-- {
		disks = append(disks, i)
	}
	tower := &Tower{n,make(map[string][]int)}
	tower.pegs["A"] = disks  // descending [3, 2, 1]
	tower.pegs["B"] = make([]int, 0)
	tower.pegs["C"] = make([]int, 0)
	return tower
}

func (tower *Tower) swap(x, y string) {
	topDiskX := tower.topDisk(x)
	topDiskY := tower.topDisk(y)
	if topDiskX == 0 && topDiskY == 0 {
		return
	}
	if (topDiskX < topDiskY && topDiskX != 0) || topDiskY == 0 {
		tower.move(x, y)
	} else {
		tower.move(y, x)
	}
}

func (tower *Tower) topDisk(peg string) int {
	count := len(tower.pegs[peg])
	if count == 0 {
		return 0
	}
	return tower.pegs[peg][count-1]
}

func (tower *Tower) move(from, to string) {
	fmt.Println(tower)
	fmt.Printf("Move from %s to %s\n", from, to)
	countFrom := len(tower.pegs[from])
	topFrom := tower.pegs[from][countFrom - 1]
	tower.pegs[from] = tower.pegs[from][:countFrom - 1]
	tower.pegs[to] = append(tower.pegs[to], topFrom)
}

func (tower *Tower) isDone() bool {
	return len(tower.pegs["C"]) == tower.n
}

func (tower *Tower) solve() {
	var steps [][]string

	if tower.n % 2 == 0 {
		steps = [][]string{{"A", "B"}, {"A", "C"}, {"B", "C"}}
	} else {
		steps = [][]string{{"A", "C"}, {"A", "B"}, {"B", "C"}}
	}
	for {
		for _, step := range steps {
			tower.swap(step[0], step[1])
			if tower.isDone() {
				return
			}
		}
	}
}

func main() {
	tower := NewTower(4)
	tower.solve()
	fmt.Println(tower)
}
