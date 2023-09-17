package main

type PathPoint struct {
	x, y   int
	parent *PathPoint
}

func (p *PathPoint) ConstructPath() *Queue {
	path := NewQueue()
	for p != nil {
		path.Push(p)
		p = p.parent
	}
	return path
}
