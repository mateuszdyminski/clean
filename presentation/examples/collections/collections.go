package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type user struct {
	name    string
	surname string
	phone   string
}

var users = []user{
	user{
		name:    "tom",
		surname: "taylor",
		phone:   "12345678",
	},
	user{
		name:    "coco",
		surname: "chanel",
		phone:   "12345612",
	},
	user{
		name:    "tommy",
		surname: "hilfiger",
		phone:   "12345123",
	},
}

func main() {
	filter := os.Args[1]

	printUsers(filter)
}

func printUsers(filter string) {
	filtered := make([]user, 0, 0)
	for _, u := range users {
		if strings.Contains(u.name, filter) {
			filtered = append(filtered, u)
		}
	}
	sort.Sort(byName(filtered))

	for _, u := range filtered {
		fmt.Printf("user: %v \n", u)
	}
}

type byName []user

func (s byName) Len() int {
	return len(s)
}
func (s byName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byName) Less(i, j int) bool {
	return s[i].name < s[j].name
}
