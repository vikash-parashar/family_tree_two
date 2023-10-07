package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Person struct represents an individual in the family tree.
type Person struct {
	Name          string
	Gender        string
	Relationships map[string][]*Person
}

// FamilyTree is a map that stores all family members.
var FamilyTree map[string]*Person

func main() {
	FamilyTree = make(map[string]*Person)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Family Tree Command Line Tool")
	fmt.Println("Enter commands or 'exit' to quit.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exiting the program.")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "add":
			if len(args) < 4 {
				fmt.Println("Usage: add <person|relationship> <name> <gender>")
				continue
			}
			entityType := args[1]
			name := args[2]
			gender := args[3]
			switch entityType {
			case "person":
				addPerson(name, gender)
			case "relationship":
				addRelationship(name)
			default:
				fmt.Println("Unknown entity type:", entityType)
			}

		case "connect":
			if len(args) < 6 || args[3] != "as" || args[4] != "of" {
				fmt.Println("Usage: connect <name 1> as <relationship> of <name 2>")
				continue
			}
			name1 := args[1]
			relationship := args[2]
			name2 := args[5]
			connect(name1, relationship, name2)

		case "count":
			if len(args) < 4 {
				fmt.Println("Usage: count <sons|daughters|wives> of <name>")
				continue
			}
			entity := args[1]
			name := args[3]
			switch entity {
			case "sons":
				countSons(name)
			case "daughters":
				countDaughters(name)
			case "wives":
				countWives(name)
			default:
				fmt.Println("Unknown entity:", entity)
			}

		case "father":
			if len(args) != 3 {
				fmt.Println("Usage: father of <name>")
				continue
			}
			name := args[2]
			fatherOf(name)

		default:
			fmt.Println("Unknown command:", command)
		}
	}
}

// Rest of the functions remain the same as in your previous code...

// Add a new person to the family tree.
func addPerson(name, gender string) {
	if gender == "" {
		gender = "unknown"
	}
	if strings.Contains(gender, "F") {
		gender = "female"
	} else if strings.Contains(gender, "M") {
		gender = "male"
	} else if strings.Contains(gender, "f") {
		gender = "female"
	} else if strings.Contains(gender, "m") {
		gender = "male"
	}
	person := &Person{
		Name:          name,
		Gender:        gender,
		Relationships: make(map[string][]*Person),
	}
	FamilyTree[name] = person
	fmt.Printf("Added %s to the family tree as %s\n", name, gender)
}

// Add a new relationship to a person in the family tree.
func addRelationship(relationship string) {
	validRelationships := []string{"father", "son", "daughter", "wife", "husband"}
	if !contains(validRelationships, relationship) {
		fmt.Println("Invalid relationship:", relationship)
		return
	}
	fmt.Printf("Added relationship: %s\n", relationship)
}

// Connect two people in the family tree based on a relationship.
func connect(name1 string, relationship string, name2 string) {
	person1, exists1 := FamilyTree[name1]
	person2, exists2 := FamilyTree[name2]

	if !exists1 || !exists2 {
		fmt.Println("Person not found in the family tree.")
		return
	}

	person1.Relationships[relationship] = append(person1.Relationships[relationship], person2)
	fmt.Printf("%s is now %s of %s\n", name1, relationship, name2)
}

// Count the number of sons of a person in the family tree.
func countSons(name string) {
	person, exists := FamilyTree[name]
	if !exists {
		fmt.Println("Person not found in the family tree.")
		return
	}

	sons, found := person.Relationships["son"]
	if !found {
		fmt.Println("No sons found.")
		return
	}

	fmt.Printf("%s has %d sons.\n", name, len(sons))
}

// Count the number of daughters of a person in the family tree.
func countDaughters(name string) {
	person, exists := FamilyTree[name]
	if !exists {
		fmt.Println("Person not found in the family tree.")
		return
	}

	daughters, found := person.Relationships["daughter"]
	if !found {
		fmt.Println("No daughters found.")
		return
	}

	fmt.Printf("%s has %d daughters.\n", name, len(daughters))
}

// Count the number of wives of a person in the family tree.
func countWives(name string) {
	person, exists := FamilyTree[name]
	if !exists {
		fmt.Println("Person not found in the family tree.")
		return
	}

	wives, found := person.Relationships["wife"]
	if !found {
		fmt.Println("No wives found.")
		return
	}

	fmt.Printf("%s has %d wives.\n", name, len(wives))
}

// Get the father of a person in the family tree.
func fatherOf(name string) {
	person, exists := FamilyTree[name]
	if !exists {
		fmt.Println("Person not found in the family tree.")
		return
	}

	fathers, found := person.Relationships["father"]
	if !found || len(fathers) == 0 {
		fmt.Println("No father found.")
		return
	}

	fmt.Printf("Father of %s is %s.\n", name, fathers[0].Name)
}

// Helper function to check if a string is in a slice of strings.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
