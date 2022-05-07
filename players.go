package main

type players struct {
	players []*player
}

func (p players) you() *player {
	return p.players[0]
}
