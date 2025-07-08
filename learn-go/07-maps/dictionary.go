package maps

import "errors"

const (
	ErrorNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrorWordDoesNotExists = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrorNotFound):
		d[word] = definition
	case err == nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrorNotFound):
		return ErrorWordDoesNotExists
	case err == nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrorNotFound):
		return ErrorWordDoesNotExists
	case err == nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
