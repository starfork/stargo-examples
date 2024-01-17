package main

import (
	"flag"
	"service/examples/internal/server"
	pb "service/examples/pkg/pb/v1"

	"github.com/joho/godotenv"
	"github.com/starfork/stargo"
)

func main() {
	_ = godotenv.Load("../../config/.env.development")
	cf := flag.String("c", "../../config/debug.yaml", "config file path")
	flag.Parse()
	sc := server.LoadConfig(*cf)
	app := stargo.New(
		stargo.Org("park"),
		stargo.Name("examples"),
		stargo.Config(sc.Server),
	)

	pb.RegisterExamplesServer(app.Server(), server.New(app))

	app.Run()

}
