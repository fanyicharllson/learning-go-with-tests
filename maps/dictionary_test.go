package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := DictionaryType{"test": DictionarySearchWord}
	// t.Run("Known word", func(t *testing.T) {
	// 	got, _ := dictionary.Search("test")
	// 	want := DictionarySearchWord
	// 	assertStrings(t, got, want)
	// })
	t.Run("UnKnown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := DictionaryType{}
		want := "this is just a new word"
		word := "new"
		err := dictionary.Add(word, want)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "new"
		definition := "this is just a new word"
		dictionary := DictionaryType{word: definition}
		err := dictionary.Add(word, "new test")

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := DictionaryType{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := DictionaryType{}

		err := dictionary.Update(word, definition)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := DictionaryType{word: "test definition"}

		err := dictionary.Delete(word)

		assertErrors(t, err, nil)

		_, err = dictionary.Search(word)

		assertErrors(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := DictionaryType{}

		err := dictionary.Delete(word)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

// reading nil to the panic err that will occur
func TestNilMapErr(t *testing.T) {
	val, ok := m["new"]

	if !ok {
		t.Fatal("Oh no! An error occured")
	}
	fmt.Println("Nil map value: ", val)

}
