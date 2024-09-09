package main

import (
	"log"
	"os"
	"owlbytech/database"
	cfg "owlbytech/internal/config"
)

var help = `Usage:
      repository-cli <command> [arguments]

The commands are:
migrate         Run the migrations
      up		  Run the Up migrations files
	  	down      Run the Down migrations files
seed        Run the migrations
      up		  Run the Up seeders files
	  	down      Run the Down seeders files
convert-mjml    Convert all MJML to HTML`

func main() {

	if len(os.Args) < 2 {
		log.Fatal(help)
	}

	switch os.Args[1] {
	case "migrate":
		if len(os.Args) < 3 {
			log.Printf("You must specify the argument to migrate command")
			log.Fatal(help)
		}

		err := database.Migrate(os.Args[2])
		if err != nil {
			log.Fatalf("%v", err)
		}
	case "seed":
		if len(os.Args) < 3 {
			log.Printf("You must specify the argument to seed command")
			log.Fatal(help)
		}

		err := database.Seeder(os.Args[2])

		if err != nil {
			log.Fatalf("%v", err)
		}
	case "convert-mjml":
		err := cfg.ConvertAllMailingTemplates()
		if err != nil {
			log.Fatalf("Error converting all MJML to HTML: %v", err)
		}
	default:
		log.Fatal(help)
	}

}
