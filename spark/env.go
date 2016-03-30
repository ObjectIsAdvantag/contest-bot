// Loads env variables
package main

import (
	"os"
)

type Environment struct {
	isCorrect bool // is the environment correctly initialized
	sparkToken string
	tropoToken string
}

var env Environment

func init() {
	env.isCorrect = true

	env.sparkToken = os.Getenv("BOT_SPARK_TOKEN")
	if env.sparkToken == "" {
		env.isCorrect = false
	}

	env.tropoToken = os.Getenv("BOT_TROPO_TOKEN")
	if env.tropoToken == "" {
		env.isCorrect = false
	}
}

/*
func IsCorrect() bool {
	return env.correct
}

func SparkToken() string {
	return env.sparkToken
}

func TropoToken() string {
	return env.tropoToken
}
*/
