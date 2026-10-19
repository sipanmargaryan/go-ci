package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T){
	got := countWords("Go is good Go")
	want := map[string]int{
		"go": 2,
		"is": 1,
		"good": 1,
	}
	if !reflect.DeepEqual(got, want){
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCountWordsEmptyString(t *testing.T) {
    got := countWords("")

    if len(got) != 0 {
        t.Errorf("got %v, want an empty map", got)
    }
}