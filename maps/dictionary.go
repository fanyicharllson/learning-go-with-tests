package maps

// TODO A map value is a pointer to a runtime.hmap structure. find out this

import (
	"fmt"
)

var DictionaryStringKey = "test"
var DictionarySearchWord = "this is just a test"

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryType map[string]string
type DictionaryErr string

var m map[string]string // nil map(nOT ALLOW)
var MakeDictionary = make(map[string]string) //another way to create map 

func Dictionary(dictionary map[string]string, word string) string {
	value := dictionary[word]
	fmt.Println("The dictionary value in Dictionary function is: ", value)
	return dictionary[word]
}

func (d DictionaryType) Search(key string) (string, error) {
	value, ok := d[key]
	fmt.Println("The dictionary value in search method: ", value)

	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d DictionaryType) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}

func (d DictionaryType) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d DictionaryType) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
