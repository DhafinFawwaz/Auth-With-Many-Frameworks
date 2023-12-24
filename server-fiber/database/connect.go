package database

import (
	"fiber-auth-template/config"
	"fmt"

	"github.com/supabase-community/supabase-go"
)

var DB *supabase.Client

func ConnectDatabase() {
	client, err := supabase.NewClient(config.Config("API_URL"), config.Config("API_KEY"), nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
		panic("failed to connect to database")
	}
	DB = client
}
