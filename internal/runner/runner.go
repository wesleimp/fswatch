package runner

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/wesleimp/fswatch/internal/command"
)

// Runner holds all the needed informarmation for start running
type Config struct {
	Command []string
	Path    string
}

func Run(c Config) error {
	if len(c.Command) < 1 {
		return fmt.Errorf("no command provided")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	donec := make(chan bool)
	sigc := make(chan os.Signal, 1)

	cmd := command.Run(c.Command)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					command.Kill(cmd)
					cmd.Process.Wait()
					cmd = command.Run(c.Command)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigc
		command.Kill(cmd)
		cmd.Process.Wait()
		donec <- true
	}()

	err = filepath.Walk(c.Path, Traverse(watcher.Add))
	if err != nil {
		log.Fatal(err)
	}

	<-donec
	return nil
}

func Traverse(add func(string) error) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		return add(path)
	}
}
