package player

type PLayer struct {
	Name string
	Mark string
	id   int
}

func CreateNewPlayer(name string, id int) *PLayer {
	var newp PLayer
	newp.Name = name
	newp.id = id
	if id == 1 {
		newp.Mark = "X"
	} else if id == 2 {
		newp.Mark = "O"
	}
	return &newp
}
