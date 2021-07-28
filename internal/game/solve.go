package game

type SolveTreeNode struct {
	x     int
	y     int
	dirs  map[Direction]SolveTreeNode
	grown bool
}

type SolveTree struct {
	root          SolveTreeNode
	current       *SolveTreeNode
	currDirection Direction
}

func (g *Game) AvailableAdvances(x int, y int) []Direction {
	dirs := make([]Direction, 0)

	for _, d := range AllDirections {
		n := g.GetNeighbor(x, y, d)
		if n != nil && *n == ColorNone {
			dirs = append(dirs, d)
		}
	}

	return dirs
}

func (g *Game) AvailableAdvancesN(x int, y int) int {
	an := 0

	for _, d := range AllDirections {
		n := g.GetNeighbor(x, y, d)
		if n != nil && *n == ColorNone {
			an++
		}
	}

	return an
}
