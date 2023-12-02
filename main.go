package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/go-template/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
