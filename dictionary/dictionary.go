package dictionary

import (
	"errors"
	"fmt"
	"time"
)

// Defines the structure of Entry
type Entry struct {
	Definition string    `json:"definition"`
	Date       time.Time `json:"date"`
}

// String returns a string representation of Entry
func (e Entry) String() string {
	return fmt.Sprintf("Definition: %s\nDate: %s", e.Definition, e.Date.Format("2006-01-02 15:04:05"))
}

// Defines the structure of the dictionary (struct)
// with entries as the key and a map as the value
// The map uses a string as the key and Entry as the value
type Dictionary struct {
	entries map[string]Entry
}

// New creates a new dictionary and returns its address
func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

// Add adds a word with its definition to the dictionary
func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{
		Definition: definition,
		Date:       time.Now(),
	}
	d.entries[word] = entry
}

// Get retrieves the definition of a word from the dictionary
func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, errors.New("Word not found in the dictionary")
	}
	return entry, nil
}

// Remove removes a word from the dictionary
func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

// List returns a list of words and the entire dictionary
func (d *Dictionary) List() ([]string, map[string]Entry) {
	wordList := make([]string, 0, len(d.entries))
	for word := range d.entries {
		wordList = append(wordList, word)
	}

	return wordList, d.entries
}
