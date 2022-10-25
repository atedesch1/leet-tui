package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type view string

const (
	questionListView view = "QuestionList"
	questionView     view = "Question"
)

type MainModel struct {
	currentView view
	models      map[view]tea.Model
}

func NewMainModel() MainModel {
	m := make(map[view]tea.Model, 0)
	m[questionListView] = newQuestionListModel()
	m[questionView] = newQuestionModel()

	return MainModel{
		currentView: questionListView,
		models:      m,
	}
}

func (m MainModel) getCurrentModel() tea.Model {
	return m.models[m.currentView]
}

func (m MainModel) Init() tea.Cmd {
	return m.getCurrentModel().Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), "q":
			return m, tea.Quit

		case tea.KeyCtrlLeft.String():
			m.currentView = questionListView

		case tea.KeyCtrlRight.String():
			m.currentView = questionView
		}

	case SelectedQuestionMsg:
		m.currentView = questionView
	}

	_, cmd := m.getCurrentModel().Update(msg)

	return m, cmd
}

func (m MainModel) View() string {
	return m.getCurrentModel().View()
}
