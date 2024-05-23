package maps

import "testing"

func TestWordCount(t *testing.T) {
	given := "bob bob bob"
	expected := map[string]int{"bob": 3}
	actual := WordCount(given)

	for k, v := range expected {
		if actual[k] != v {
			t.Fatalf("Expected %v:%v, got %v", k, v, actual[k])
		}
	}
}

func TestWordCountManyWords(t *testing.T) {
	given := "bob alice carol carol bob carol carol bob 		carol"
	expected := map[string]int{
		"bob":   3,
		"alice": 1,
		"carol": 5,
	}
	actual := WordCount(given)

	for k, v := range expected {
		if actual[k] != v {
			t.Fatalf("Expected %v:%v, got %v", k, v, actual[k])
		}
	}
}
