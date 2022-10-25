package main

import (
	"github.com/atedesch1/leet-tui/api"
	"github.com/atedesch1/leet-tui/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	api.Init()
	p := tea.NewProgram(models.NewMainModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
