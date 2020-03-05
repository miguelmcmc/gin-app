package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func performLogin(c *gin.Context) {
	// Obtain the POSTed name and password values
	name := c.PostForm("name")
	fmt.Println("Name-login: ", name)

	// var sameSiteCookie http.SameSite

	// Check if the name/password combination is valid
	if isUserValid(name) {
		// If the name/password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login", "payload": name}, "login-successful.html")

	} else {
		// If the name/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	// var sameSiteCookie http.SameSite

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showUserCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New User"}, "create-user.html")
}

func getUser(c *gin.Context) {
	// Check if the user ID is valid
	if userID, err := strconv.Atoi(c.Param("user_id")); err == nil {
		// Check if the user exists
		if user, err := getUserByID(userID); err == nil {
			// Call the render function with the title, user and the name of the
			// template
			render(c, gin.H{
				"title":   user.Name,
				"payload": user}, "user.html")

		} else {
			// If the user is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid user ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createUser(c *gin.Context) {
	// Obtain the POSTed title and content values
	name := c.PostForm("name")
	mail := c.PostForm("mail")

	if a, err := createNewUser(name, mail); err == nil {
		// If the user is created successfully, show success message
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-user-successful.html")
	} else {
		// if there was an error while creating the user, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
