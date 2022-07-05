package player

type PLayer struct {
	Name string
	Mark string
	id   int
}

func CreateNewPlayer(name string, id int, mark string) *PLayer {
	var newp PLayer
	newp.Name = name
	newp.id = id
	newp.Mark = mark
	return &newp
}
