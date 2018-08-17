# Model

Esta é a documentação do módulo `Model`, que pode ser utilizado carregar
informações dos testes RPM.

## `func GetPercentile() []int`

Gets the percentile list.


## `func GetPercentileScoreByAge(age int) ([]int, error)`

Gets minimum score for a percentile based on the age.

## `func GetTestItems() []string`

This function get the each test item's id.


## `func GetHowManyOptionsForItem() []int`

Gets how many options there are for each item.


## `func GetCorrectAnswers() []int`

Gets the correct answer for each item.


## `func GetItemSeries() []string`

Gets which series the current item is in.


## `func GetValiditiesForSeries(series string) map[int]int`

Relates the total score to the required partial score for each series.
