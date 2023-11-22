package models

type invite struct {
	whoInvite  *player
	whoInvited *player
	adress     *lobby
}

func (i invite) JoinAsMember() {
	i.whoInvited.lobby = i.adress
	i.whoInvited.role = member{}
	i.adress.players = append(i.adress.players, i.whoInvited)
}

func (i invite) JoinAsSpectator() {
	i.whoInvited.lobby = i.adress
	i.whoInvited.role = spectator{}
	i.adress.players = append(i.adress.players, i.whoInvited)

}

func (i *invite) RejectInvite() {
	i = nil
}
