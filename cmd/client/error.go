package main

func (m *model) ErrorState(err string, state string) {
	m.err.Err = err
	m.state = "repeat"
	m.err.ToState = state
	m.helpStr = ""
	m.choices = m.currentChoices[m.state]
	m.cursor = 0
}
