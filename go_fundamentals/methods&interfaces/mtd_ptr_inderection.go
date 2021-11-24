package main

type Location struct {
	X, Y      float64
	isSpecial bool
}

func (loc Location) dist(loc2 Location) (x, y float64) {
	x, y = loc2.X-loc.X, loc2.Y-loc.Y
	return
}

func (loc *Location) makeSpecial() {
	loc.isSpecial = true
}

func main() {
	//indirection 1
	var loc1 Location = Location{3, 4, false}
	//here it interprets as (&loc1).makeSpecial()
	loc1.makeSpecial()

	//indirection 2
	var loc2 *Location = &Location{3, 4, false}
	//here it interprets as (*loc1).dist(Location)
	loc2.dist(loc1)
}
