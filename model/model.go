package model

// Just a function to kickstart stuff
func GetTrue() bool {
    return true
}

// Gets the percentile table file path
func getPercentileFilePath() string {
    return "C:\\Users\\cris\\Documents\\scripts\\golang\\data\\dojoday\\percentile.csv"
}

// Gets the answers table file path
func getAnswerFilePath() string {
    return "C:\\Users\\cris\\Documents\\scripts\\golang\\data\\dojoday\\respostas.csv"
}

// Gets the validity table file path
func getValidityFilePath() string {
    return "C:\\Users\\cris\\Documents\\scripts\\golang\\data\\dojoday\\validade.csv"
}

// Gets the percentile according to the age and the score obtained by the
// subject on the test
func GetPercentileByAgeAndScore(age, score int) int {
    filepath := getPercentileFilePath()
    // TODO Implement this function
    return -1
}
