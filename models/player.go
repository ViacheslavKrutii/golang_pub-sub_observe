package models

import "errors"

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
	newLobby = &lobby{players: []*player{p}}
	p.lobby = newLobby
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
	for _, v := range p.invites {

	}
}
