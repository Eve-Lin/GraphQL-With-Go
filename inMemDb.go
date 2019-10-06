package main

var tutorials []Tutorial

func dbInit() []Tutorial {

	author := &Author{Name: "Jane Doe", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQL Getting Started",
		Author: *author,
		Comments: []Comment{
			Comment{
				Body: "First Comment"},
		},
	}

	tutorials = append(tutorials, tutorial)

	return tutorials
}
