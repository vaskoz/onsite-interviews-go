package skyline

import (
	"container/heap"
	"sort"
)

// Point represents a 0-dimensional geometric point.
type Point struct {
	X, Y int
}

// Building represents a 2-dimensional representation of a building in a skyline.
type Building struct {
	LeftX, RightX, Height int
}

type edge struct {
	X, Height int
	Up        bool
}

// HeightHeap is a max-heap of building heights.
type HeightHeap []int

func (h HeightHeap) Len() int           { return len(h) }
func (h HeightHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h HeightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds a new height to the heap.
func (h *HeightHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the max height from the heap.
func (h *HeightHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Skyline returns all the points describing a skyline silhouette.
// Runs in O(N log N) time with some extra space O(N)
func Skyline(buildings []Building) []Point {
	edges := make([]edge, 0, 2*len(buildings))
	for _, building := range buildings {
		edges = append(edges, edge{building.LeftX, building.Height, true})
		edges = append(edges, edge{building.RightX, building.Height, false})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].X < edges[j].X
	})
	var points []Point
	h := &HeightHeap{}
	heap.Init(h)
	for _, edge := range edges {
		if edge.Up {
			if h.Len() == 0 {
				points = append(points, Point{edge.X, 0})
				points = append(points, Point{edge.X, edge.Height})
			} else if edge.Height > (*h)[0] {
				points = append(points, Point{edge.X, (*h)[0]})
				points = append(points, Point{edge.X, edge.Height})
			}
			heap.Push(h, edge.Height)
		} else if !edge.Up {
			var removeIndex int
			for i, v := range *h {
				if v == edge.Height {
					removeIndex = i
					break
				}
			}
			heap.Remove(h, removeIndex)
			if h.Len() == 0 {
				points = append(points, Point{edge.X, edge.Height})
				points = append(points, Point{edge.X, 0})
			} else if edge.Height > (*h)[0] {
				points = append(points, Point{edge.X, edge.Height})
				points = append(points, Point{edge.X, (*h)[0]})
			}
		}
	}
	return points
}
