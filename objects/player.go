package objects

import (
	"Proj/golang_pub-sub_observe/observer"
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

type player struct {
	name    string
	lobby   *lobby
	invites []invite
	role    role
}

type role interface {
	role()
}

type member struct {
}

func (m member) role() {}

type spectator struct {
}

func (s spectator) role() {}

func CreatePlayer(name string) *player {
	return &player{name: name, lobby: nil, role: nil}
}

func (p *player) CreateLobby() (newLobby *lobby, err error) {
	if p.lobby != nil {
		return nil, errors.New("you already in lobby")
	}
	newLobby = &lobby{players: []*player{p}, moves: make(map[int]*player, 0), turn: 0, subscribers: []observer.Observer{p}}
	p.lobby = newLobby
	p.role = member{}
	return newLobby, nil
}

func (p *player) LeaveLobby() {
	p.lobby = nil
}

func (p *player) InvitePlayer(p2 *player) {
	if p.lobby == nil {
		return
	}
	newInvite := invite{adress: p.lobby, whoInvite: p, whoInvited: p2}
	p2.invites = append(p2.invites, newInvite)
}

func (p *player) CheckInvites() {
	for idx, invite := range p.invites {
		r := rand.Intn(3)
		switch r {
		case 0:
			invite.JoinAsMember()
			slices.Delete(p.invites, idx, idx)
		case 1:
			invite.JoinAsSpectator()
			slices.Delete(p.invites, idx, idx)
		default:
			invite.RejectInvite()
			slices.Delete(p.invites, idx, idx)
		}
	}
}

func (p *player) Move() {
	_, ok := p.role.(member)
	if p.lobby != nil && ok {
		p.lobby.moves[p.lobby.turn] = p
		p.lobby.Notify(fmt.Sprintf("Player %s turn %d", p.name, p.lobby.turn))
		p.lobby.turn++
		return
	}

}

func (p *player) Observe(subject any) {
	fmt.Printf("%s received a notification %v\n", p.name, subject)
}
