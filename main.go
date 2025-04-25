package main

import "github.com/adamcheaib/blog-aggregator-goalng/internal/config"

func main() {
	config, _ := config.Read()
	config.SetUser("Rastaman Live up")

}
