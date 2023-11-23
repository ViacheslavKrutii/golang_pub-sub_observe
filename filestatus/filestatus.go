package filestatus

import (
	"Proj/golang_pub-sub_observe/pubsub"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func Filestatus(directory string, broker *pubsub.Broker) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Рекурсивно додати каталог до відстежувача
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return watcher.Add(path)
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				broker.Publish("FS Event", "Write")
			} else if event.Op&fsnotify.Create == fsnotify.Create {
				broker.Publish("FS Event", "Create")
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				broker.Publish("FS Event", "Remove")
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Помилка відстеження:", err)
		}
	}
}
