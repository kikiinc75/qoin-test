package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id    int
	dices []int
	score int
}

func (p *Player) rollDadu() {
	for i := 0; i < len(p.dices); i++ {
		p.dices[i] = rand.Intn(6) + 1
	}
}

func (p *Player) evaluasiDadu(pemain []*Player) {
	for i := 0; i < len(p.dices); i++ {
		switch p.dices[i] {
		case 6:
			p.score++
			p.dices = append(p.dices[:i], p.dices[i+1:]...)
			i--
		case 1:
			if p.id+1 < len(pemain) {
				pemain[p.id+1].dices = append(pemain[p.id+1].dices, 1)
			} else {
				pemain[0].dices = append(pemain[0].dices, 1)
			}
			p.dices = append(p.dices[:i], p.dices[i+1:]...)
			i--
		}
	}
}

func playGame(p []*Player) *Player {
	for len(p) > 1 {
		for _, val := range p {
			val.rollDadu()
			val.evaluasiDadu(p)
		}

		var remaining []*Player
		for _, val := range p {
			if len(val.dices) > 0 {
				remaining = append(remaining, val)
			}
		}
		p = remaining
	}

	return p[0]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var N, M int
	fmt.Print("masukkan jumlah player (N): ")
	fmt.Scan(&N)
	fmt.Print("masukkan jumlah dadu per player (M): ")
	fmt.Scan(&M)

	newPlayers := make([]*Player, N)
	for i := 0; i < N; i++ {
		newPlayers[i] = &Player{
			id:    i,
			dices: make([]int, M),
		}
	}

	pemenang := playGame(newPlayers)
	fmt.Printf("Player %d wins with %d points!\n", pemenang.id+1, pemenang.score)
}
