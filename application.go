package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": "homepage"}, "index.html")
}

func showApplicationsPage(c *gin.Context) {
	applications := getAllApplications()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": applications}, "applications.html")
}
func showUsersPage(c *gin.Context) {
	users := getAllUsers()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": users}, "users.html")
}

func showApplicationCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Application"}, "create-application.html")
}

func getApplication(c *gin.Context) {
	// Check if the application ID is valid
	if applicationID, err := strconv.Atoi(c.Param("application_id")); err == nil {
		// Check if the application exists
		if application, err := getApplicationByID(applicationID); err == nil {
			// Call the render function with the title, application and the name of the
			// template
			render(c, gin.H{
				"title":   application.Name,
				"payload": application}, "application.html")

		} else {
			// If the application is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid application ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createApplication(c *gin.Context) {
	// Obtain the POSTed title and content values
	name := c.PostForm("name")
	description := c.PostForm("description")
	mobility := c.PostForm("mobility")
	mobilityFlag := false
	if mobility == "on" {
		mobilityFlag = true
	}

	if a, err := createNewApplication(name, description, mobilityFlag); err == nil {
		// If the application is created successfully, show success message
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-application-successful.html")
	} else {
		// if there was an error while creating the application, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
