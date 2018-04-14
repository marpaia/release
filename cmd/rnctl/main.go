package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := cli.NewApp()
	app.Name = "rnctl"
	app.Usage = "A workflow tool for managing the on-going aggregation of Kubernetes release notes"
	app.Version = "0.0.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "pull",
			Usage: "Pull recently merged PRs into the local database",
			Action: func(cliCtx *cli.Context) error {
				fmt.Println("Downloading recently merged Kubernetes pull requests...\n")

				bar := pb.StartNew(100)
				for i := 0; i < 100; i++ {
					bar.Increment()
					time.Sleep(15 * time.Millisecond)
				}

				bar.FinishPrint("Database updated!\n")
				return nil
			},
		},
		cli.Command{
			Name:  "classify",
			Usage: "Interactively classify an unclassified merged PR",
			Action: func(cliCtx *cli.Context) error {
				return errors.New("rnctl classify is unimplemented")
			},
		},
		cli.Command{
			Name:  "render",
			Usage: "Render classified release notes for a given subset of a release",
			Action: func(cliCtx *cli.Context) error {
				return errors.New("rnctl generate is unimplemented")
			},
		},
		cli.Command{
			Name:  "status",
			Usage: "Print the status of the current release notes database",
			Action: func(cliCtx *cli.Context) error {
				return errors.New("rnctl status is unimplemented")
			},
		},
	}

	app.RunAndExitOnError()
}
