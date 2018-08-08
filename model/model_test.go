package model

import (
    "testing"
)

func TestCanActuallyTest(t *testing.T) {
    isTrue := GetTrue()
    if isTrue != true {
        t.Error("true is not true? wtf??")
    }
}

func TestContainRequiredEnvironmentVariable(t *testing.T) {
    sourceFolder := GetSourceFolder()
    if len(sourceFolder) == 0 {
        t.Error("DOJODAY_FILES environment variable not set")
    }
}
