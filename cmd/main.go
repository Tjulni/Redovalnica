package main

import(
	"fmt"
	"context"
	"os"
	"log"

	"github.com/urfave/cli/v3"
	"github.com/Tjulni/Redovalnica/redovalnica"
)

func main(){
	cmd := &cli.Command{
		Name: "redovalnica",
		Usage: "Upravljanje redovalnice študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
			},
			&cli.IntFlag{
				Name: "minOcena",
				Usage: "Najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name: "maxOcena",
				Usage: "Največja možna ocena",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}