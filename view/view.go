package view

type StateEnum int
const (
    AgeQuestionState StateEnum = 1 << iota
    TestState
)

type View struct {
    Working bool
    CurrentItem int
    Age int
    State StateEnum
    Output []string
}

func NewView() View {
    return View {
        Working: true,
        CurrentItem: 0,
        Age: 10,
        State: AgeQuestionState,
        Output: []string {
            "Digite a sua idade:",
        },
    }
}

func (view *View) Write() string {
    output := ""
    for _, line := range view.Output {
        output += line + "\n"
    }
    view.Output = make([]string, 0)
    return output
}

func (view *View) Read(line string) {
    // TODO Implement the rest of me!
    switch view.State {
    case AgeQuestionState:
        view.State = TestState
        view.Output = append(view.Output, "Digite qualquer tecla para sair")
        break
    case TestState:
        view.Working = false
    }
}
