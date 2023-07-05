package main

import (
	"embed"

	"github.com/Arsfiqball/code-generator/cmd"
)

//go:embed cmd/templates/*
var templates embed.FS

func init() {
	cmd.Templates = templates
}

func main() {
	cmd.Execute()
}
