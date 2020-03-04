package main

import "errors"

type application struct {
	ID          int    `json:"id"`
	Name        string `json:"title"`
	Mobility    bool   `json:"mobility"`
	Description string `json:"content"`
}

// For this demo, we're storing the application list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var applicationList = []application{
	application{ID: 1, Name: "Amazon Web Services", Mobility: false, Description: "Amazon web service es un servicio en la nube"},
}

// Return a list of all the applications
func getAllApplications() []application {
	return applicationList
}

// Fetch an application based on the ID supplied
func getApplicationByID(id int) (*application, error) {
	for _, a := range applicationList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Application not found")
}

// Create a new application with the title and content provided
func createNewApplication(name, description string, mobility bool) (*application, error) {
	// Set the ID of a new application to one more than the number of applications
	a := application{ID: len(applicationList) + 1, Name: name, Description: description, Mobility: mobility}

	// Add the application to the list of applications
	applicationList = append(applicationList, a)

	return &a, nil
}
