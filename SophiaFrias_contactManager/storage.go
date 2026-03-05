package main

import (
	"encoding/json"
	"os"
)

// LoadContacts reads contacts from a JSON file.
// If the file doesn't exist, it returns an empty slice (not an error).
func LoadContacts(path string) ([]Contact, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{}, nil
		}
		return nil, err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Treat empty file as empty contacts.
	if len(b) == 0 {
		return []Contact{}, nil
	}

	var contacts []Contact
	if err := json.Unmarshal(b, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

// SaveContacts writes contacts to a JSON file atomically (best-effort).
func SaveContacts(path string, contacts []Contact) error {
	b, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}
