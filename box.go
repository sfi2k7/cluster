package main

type Box struct {
	NorthWest     *Location
	SouthEast     *Location
	Center        *Star
	Stars         []*Star
	VirtualCenter *Location
}

func NewBox(nw *Location, se *Location) *Box {
	b := Box{NorthWest: nw, SouthEast: se}
	x := (nw.Longitude + se.Longitude) / 2
	y := (nw.Latitude + se.Latitude) / 2
	b.VirtualCenter = &Location{Longitude: x, Latitude: y}
	return &b
}

func (b *Box) addPoint(l *Location) {
	d := distance(l.Latitude, l.Longitude, b.VirtualCenter.Latitude, b.VirtualCenter.Longitude)
	s := &Star{Location: l, Distence: d}
	if b.Center == nil || b.Center.Distence > d {
		b.Center = s
		b.recalculateDistenceFromCenter()
	}
	b.Stars = append(b.Stars, s)
}

func (b *Box) recalculateDistenceFromCenter() {
	for _, s := range b.Stars {
		s.Distence = distance(b.Center.Location.Latitude, b.Center.Location.Longitude, s.Location.Latitude, s.Location.Longitude)
	}
}

/*
   When a box is created trry to image a virtual center.
   When a point is added then its distence from VC is teted.
   if new Star has Min distence form VC then it becomes the center.
       And all existing points are updated with the distence from center
   Process is repeated for each point.

*/
