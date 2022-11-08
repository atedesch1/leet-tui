package models

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type questionListModel struct {
	totalNumQuestions    int
	questions            []api.QuestionInfo
	pageToQuestionsIndex map[int]int
	page                 int
	limit                int
	categorySlug         string

	cursor  int
	loading bool
}

func newQuestionListModel() *questionListModel {
	return &questionListModel{
		totalNumQuestions:    0,
		questions:            []api.QuestionInfo{},
		pageToQuestionsIndex: make(map[int]int),
		page:                 0,
		limit:                10,
		categorySlug:         "",

		cursor:  0,
		loading: true,
	}
}

type problemsetQuestionListResponseMsg struct {
	questionList api.ProblemsetQuestionList
}

func (m *questionListModel) getProblemsetQuestionListCmd() tea.Msg {
	problemsetQuestionList, _ := api.GetProblemsetQuestionList(m.categorySlug, m.page*m.limit, m.limit)
	return problemsetQuestionListResponseMsg{
		questionList: problemsetQuestionList,
	}
}

func (m *questionListModel) goToPreviousPage() tea.Cmd {
	if m.page > 0 {
		m.page--
	}

	if _, ok := m.pageToQuestionsIndex[m.page]; !ok {
		m.loading = true
		return m.getProblemsetQuestionListCmd
	}
	return nil
}

func (m *questionListModel) goToNextPage() tea.Cmd {
	if (1+m.page)*m.limit <= m.totalNumQuestions {
		m.page++
	}

	if _, ok := m.pageToQuestionsIndex[m.page]; !ok {
		m.loading = true
		return m.getProblemsetQuestionListCmd
	}
	return nil
}

type SelectedQuestionMsg struct {
	question api.QuestionInfo
}

func (m *questionListModel) selectQuestionCmd() tea.Msg {
	return SelectedQuestionMsg{
		question: m.questions[m.pageToQuestionsIndex[m.page]*m.limit+m.cursor],
	}
}

func (m *questionListModel) Init() tea.Cmd {
	return m.getProblemsetQuestionListCmd
}

func (m *questionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case problemsetQuestionListResponseMsg:
		m.loading = false
		if m.totalNumQuestions == 0 {
			m.totalNumQuestions = msg.questionList.Total
		}
		m.pageToQuestionsIndex[m.page] = len(m.questions)
		m.questions = append(m.questions, msg.questionList.Questions...)

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 && m.page > 0 {
				cmd = m.goToPreviousPage()
				m.cursor = m.limit - 1
			}

		case "down", "j":
			if m.cursor < m.limit-1 {
				m.cursor++
			} else if m.cursor == m.limit-1 && (1+m.page)*m.limit <= m.totalNumQuestions {
				cmd = m.goToNextPage()
				m.cursor = 0
			}

		case "right", "l":
			cmd = m.goToNextPage()

		case "left", "h":
			cmd = m.goToPreviousPage()

		case "enter", " ":
			return m, m.selectQuestionCmd
		}

	}
	return m, cmd
}

func (m *questionListModel) View() string {
	s := "Choose a problem from below:\n\n"

	if m.loading {
		s += "  Loading..."
		return s
	}
	index := m.pageToQuestionsIndex[m.page]
	for i, choice := range m.questions[index : index+m.limit] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	s += fmt.Sprintf("\nPage %d/%d", m.page+1, m.totalNumQuestions/m.limit)

	return s
}
