package main

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/users", ensureLoggedIn(), showUsersPage)
	router.GET("/applications", ensureLoggedIn(), showApplicationsPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /user/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /user/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle GET requests at /user/view/some_user_id
		userRoutes.GET("/view/:user_id", ensureLoggedIn(), getUser)

		// Handle the GET requests at /user/create
		// Show the user creation page
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/create", ensureLoggedIn(), showUserCreationPage)

		// Handle POST requests at /user/create
		// Ensure that the user is logged in by using the middleware
		userRoutes.POST("/create", ensureLoggedIn(), createUser)
	}

	// Group application related routes together
	applicationRoutes := router.Group("/application")
	{
		// Handle GET requests at /application/view/some_application_id
		applicationRoutes.GET("/view/:application_id", ensureLoggedIn(), getApplication)

		// Handle the GET requests at /application/create
		// Show the application creation page
		// Ensure that the user is logged in by using the middleware
		applicationRoutes.GET("/create", ensureLoggedIn(), showApplicationCreationPage)

		// Handle POST requests at /application/create
		// Ensure that the user is logged in by using the middleware
		applicationRoutes.POST("/create", ensureLoggedIn(), createApplication)
	}
}
