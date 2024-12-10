package structure

type Room struct {
	Name     string
	X, Y     int
	IsStart  bool
	IsEnd    bool
	Occupied int
}

type Graph struct {
	Rooms map[string]*Room
	Links []Link
}

type Ant struct {
	ID       int
	Position string
}

type Link struct {
	Room1 string
	Room2 string
}
