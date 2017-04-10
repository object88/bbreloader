package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gobwas/glob"
	"github.com/rjeczalik/notify"
	"github.com/urfave/cli"
)

const (
	source string = "source"
)

func main() {
	app := cli.NewApp()
	app.Name = "bbreloader"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "BBRELOADER_SOURCE",
			Name:   source,
		},
	}

	app.Action = func(c *cli.Context) error {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)

		err := Watch(c.String(source))
		if err != nil {
			return err
		}

		<-sigchan

		// Let's see if this changes.

		return nil
	}

	app.Run(os.Args)
}

func watch(notifyChan chan notify.EventInfo, callback func(*collectedEvents)) {
	lull := time.Duration(2 * time.Second)

	events := newCollectedEvents()

	globSpec := "**/*.go"
	g := glob.MustCompile(globSpec)

	for {
		select {
		case e := <-notifyChan:
			path := e.Path()
			if g.Match(path) {
				// We have a match!
				switch e.Event() {
				case notify.Create:
					events.created[path] = true
				case notify.Remove:
					events.removed[path] = true
				case notify.Rename:
					events.renamed[path] = true
				case notify.Write:
					events.written[path] = true
				}
			}
		case <-time.After(lull):
			processed := events
			events = newCollectedEvents()
			go callback(processed)
		}
	}
}

func Watch(basePath string) error {
	notifyChan := make(chan notify.EventInfo, 4096)

	err := notify.Watch(basePath, notifyChan, notify.All)
	if err != nil {
		notify.Stop(notifyChan)
		return err
	}

	go watch(notifyChan, func(events *collectedEvents) {
		if len(events.written) > 0 {
			fmt.Printf("Changed files: %#v\n", events.written)
			absPath, absErr := filepath.Abs(basePath)
			if absErr != nil {
				log.Fatalf("Failed to calculate the absolute path: %s", absErr)
			}

			cmd := exec.Command("go", "build", "-o", "./bin/bb")
			cmd.Dir = absPath
			runErr := cmd.Run()
			if runErr != nil {
				log.Fatalf("Build command failed: %s\n", runErr.Error())
				return
			}

			log.Printf("Finished compilation")
		}
	})

	return nil
}
