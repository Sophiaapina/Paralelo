package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger, err := InitLogger("activity.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	contacts, err := LoadContacts(DefaultDataFile)
	if err != nil {
		logger.Errorf("failed to load contacts: %v", err)
		fmt.Fprintf(os.Stderr, "Error loading contacts: %v\n", err)
		// Continue with empty list to allow recovery
		contacts = []Contact{}
	}

	logger.Infof("program started, loaded %d contacts", len(contacts))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Contact Manager (Go) ===")
		fmt.Println("1) Add contact")
		fmt.Println("2) Edit contact")
		fmt.Println("3) Delete contact")
		fmt.Println("4) Show contacts")
		fmt.Println("5) Exit")
		fmt.Print("Choose an option: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			var c Contact
			c.Name = promptNonEmpty(reader, "Name: ")
			c.Phone = promptNonEmpty(reader, "Phone: ")
			c.Email = promptNonEmpty(reader, "Email: ")

			contacts = append(contacts, c)
			if err := SaveContacts(DefaultDataFile, contacts); err != nil {
				logger.Errorf("add contact failed (save): %v", err)
				fmt.Fprintf(os.Stderr, "Error saving contacts: %v\n", err)
			} else {
				logger.Infof("added contact: name=%q phone=%q email=%q", c.Name, c.Phone, c.Email)
				fmt.Println("Contact added.")
			}

		case "2":
			if len(contacts) == 0 {
				fmt.Println("No contacts to edit.")
				continue
			}
			printContacts(contacts)

			idx := promptIndex(reader, "Enter the contact number to edit: ", len(contacts))
			orig := contacts[idx]

			fmt.Println("Leave blank to keep current value.")
			newName := promptOptional(reader, fmt.Sprintf("Name [%s]: ", orig.Name))
			newPhone := promptOptional(reader, fmt.Sprintf("Phone [%s]: ", orig.Phone))
			newEmail := promptOptional(reader, fmt.Sprintf("Email [%s]: ", orig.Email))

			if newName != "" {
				contacts[idx].Name = newName
			}
			if newPhone != "" {
				contacts[idx].Phone = newPhone
			}
			if newEmail != "" {
				contacts[idx].Email = newEmail
			}

			if err := SaveContacts(DefaultDataFile, contacts); err != nil {
				logger.Errorf("edit contact failed (save): %v", err)
				fmt.Fprintf(os.Stderr, "Error saving contacts: %v\n", err)
				contacts[idx] = orig // best-effort rollback in memory
			} else {
				logger.Infof("edited contact #%d: from=%v to=%v", idx+1, orig, contacts[idx])
				fmt.Println("Contact updated.")
			}

		case "3":
			if len(contacts) == 0 {
				fmt.Println("No contacts to delete.")
				continue
			}
			printContacts(contacts)

			idx := promptIndex(reader, "Enter the contact number to delete: ", len(contacts))
			deleted := contacts[idx]

			contacts = append(contacts[:idx], contacts[idx+1:]...)
			if err := SaveContacts(DefaultDataFile, contacts); err != nil {
				logger.Errorf("delete contact failed (save): %v", err)
				fmt.Fprintf(os.Stderr, "Error saving contacts: %v\n", err)
				// best-effort rollback in memory
				contacts = append(contacts[:idx], append([]Contact{deleted}, contacts[idx:]...)...)
			} else {
				logger.Infof("deleted contact #%d: %v", idx+1, deleted)
				fmt.Println("Contact deleted.")
			}

		case "4":
			loaded, err := LoadContacts(DefaultDataFile)
			if err != nil {
				logger.Errorf("show contacts failed (load): %v", err)
				fmt.Fprintf(os.Stderr, "Error reading contacts: %v\n", err)
				continue
			}
			logger.Infof("show contacts: loaded %d contacts from file", len(loaded))
			if len(loaded) == 0 {
				fmt.Println("No contacts found.")
			} else {
				printContacts(loaded)
			}

		case "5":
			logger.Infof("program exit requested")
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1-5.")
		}
	}
}

func printContacts(contacts []Contact) {
	fmt.Println("\n--- Contacts ---")
	for i, c := range contacts {
		fmt.Printf("%d) %s | %s | %s\n", i+1, c.Name, c.Phone, c.Email)
	}
}

func promptNonEmpty(reader *bufio.Reader, label string) string {
	for {
		fmt.Print(label)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if s != "" {
			return s
		}
		fmt.Println("Value cannot be empty.")
	}
}

func promptOptional(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func promptIndex(reader *bufio.Reader, label string, max int) int {
	for {
		fmt.Print(label)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		n, err := strconv.Atoi(s)
		if err != nil || n < 1 || n > max {
			fmt.Printf("Please enter a number between 1 and %d.\n", max)
			continue
		}
		return n - 1
	}
}
