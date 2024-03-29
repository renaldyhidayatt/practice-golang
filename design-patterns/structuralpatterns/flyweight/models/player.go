package models

import "structuralpatterns/flyweight/interfaces"

type Player struct {
	Dress      interfaces.Dress
	PlayerType string
	Lat        int
	Long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		PlayerType: playerType,
		Dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.Lat = lat
	p.Long = long
}
