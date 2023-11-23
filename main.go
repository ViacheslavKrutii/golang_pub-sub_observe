package main

import (
	"Proj/golang_pub-sub_observe/filestatus"
	"Proj/golang_pub-sub_observe/objects"
	"Proj/golang_pub-sub_observe/pubsub"
	"Proj/golang_pub-sub_observe/users"
	"fmt"
)

func observerExample() {
	bob := objects.CreatePlayer("Bob")
	alice := objects.CreatePlayer("Alice")
	ignat := objects.CreatePlayer("Ignat")

	bob.CreateLobby()
	ignat.CreateLobby()
	bob.InvitePlayer(alice)
	ignat.InvitePlayer(alice)
	alice.CheckInvites()
	bob.Move()
	bob.Move()
	bob.Move()
	bob.Move()
	bob.Move()

}

func pubsubExample() {
	topic := "FS Event"
	const write = "Write"
	const create = "Create"
	const remove = "Remove"

	gregUser := users.NewUser("Greg", write)
	bobUser := users.NewUser("Bob", create, remove)

	broker := pubsub.NewBroker()

	greg := broker.AddSubscriber()
	bob := broker.AddSubscriber()

	broker.Subscribe(greg, topic)
	broker.Subscribe(bob, topic)

	go filestatus.Filestatus(".\\path", broker)

	go func() {
		for message := range greg.Messages() {
			if gregUser.IsInterested(message.GetMessageBody()) {
				fmt.Println(gregUser.Name, "found what he interested in!: ", message.GetMessageBody())
			}
		}
	}()

	go func() {
		for message := range bob.Messages() {
			if bobUser.IsInterested(message.GetMessageBody()) {
				fmt.Println(bobUser.Name, "found what he interested in!: ", message.GetMessageBody())
			}
		}
	}()

	fmt.Scanln()
	fmt.Println("Done!")
}

func main() {
	// observerExample()
	pubsubExample()
}
