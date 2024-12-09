package structure

type Room struct {
	Name string
	X, Y int
	IsStart bool
	IsEnd bool
	Occupied bool
}

type Graph struct {
	Rooms map[string]*Room
	Links map[string][]string
}