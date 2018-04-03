package main

import "github.com/vegh1010/moviems/tools"

//Verify if those environment has been define. If not, output missing variables
func CheckEnvironment() {
	env := tools.NewEnvironmentManager()
	env.Validate("URL")
	env.Validate("ACCESS_TOKEN")
	env.Validate("PORT")

	env.Result()
}
