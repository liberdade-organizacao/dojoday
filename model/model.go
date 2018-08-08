package model

import (
    "os"
)

// Just a function to kickstart stuff
func GetTrue() bool {
    return true
}

func GetSourceFolder() string {
    return os.Getenv("DOJODAY_FILES")
}

// Gets the percentile table file path
func getPercentileFilePath() string {
    return GetSourceFolder() + "percentile.csv"
}

// Gets the answers table file path
func getAnswerFilePath() string {
    return GetSourceFolder() + "respostas.csv"
}

// Gets the validity table file path
func getValidityFilePath() string {
    return GetSourceFolder() + "validade.csv"
}

// Gets the percentile list
func GetPercentile() []int {
    outlet := make([]int, 0)
    // TODO Implement this function
    return outlet
}

// TODO Get minimum score for a percentile based on the age
func GetPercentileScoreByAge(age int) []int {
    outlet := make([]int, 0)
    return outlet
}
