package main

import "github.com/vegh1010/moviems/tools"

// CheckEnvironment verify if those environment has been define. If not, output missing variables
func CheckEnvironment() {
	//new instance of EnvironmentManager
	env := tools.NewEnvironmentManager()

	//add tags for validation
	env.Validate("URL")
	env.Validate("ACCESS_TOKEN")
	env.Validate("PORT")

	//check result
	env.Result()
}
