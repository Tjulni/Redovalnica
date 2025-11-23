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
				Value: 6,
			},
			&cli.IntFlag{
				Name: "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 0,
			},
			&cli.IntFlag{
				Name: "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			st := cmd.Int("stOcen")
			min := cmd.Int("minOcena")
			max := cmd.Int("maxOcena")
			
			redovalnica.StOcen = st
			redovalnica.MinOcena = min
			redovalnica.MaxOcena = max

			stud := map[string]redovalnica.Student{
				"123": {Ime: "Ana", Priimek: "Kovač"},
				"456": {Ime: "Marko", Priimek: "Novak"},
			}

			redovalnica.DodajOceno(stud, "123", 8)
			redovalnica.DodajOceno(stud, "123", 10)
			redovalnica.DodajOceno(stud, "456", 4)

			fmt.Println("Ocene:")
			redovalnica.IzpisVsehOcen(stud)

			fmt.Println("\nKončni uspeh:")

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}