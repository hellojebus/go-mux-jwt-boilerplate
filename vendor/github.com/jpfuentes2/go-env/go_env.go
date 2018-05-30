/*
	The env.go package sources an environment file and sets the variables => values
	for the executing Go process using `os.Setenv()`

	If you would like to automatically souce your `pwd`/.env then see env/autoload/autoload.go

	Here is an example .env file:

		MY_VAR=hi
		ANOTHER_VAR=wat
		#COMMENTED_VAR=does-not-get-set

	Note the #COMMENTED_VAR will not be set.

	Usage:

		import "env"
		env.ReadEnv("/my/path/.env")
*/
package env

import (
	"io/ioutil"
	"os"
	"strings"
)

// ReadEnv reads an .env formatted file and sets the vars to the current
// environment so exec calls respect the settings.
func ReadEnv(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(content), "\n") {
		tokens := strings.SplitN(line, "=", 2)
		if len(tokens) == 2 && tokens[0][0] != '#' {
			k, v := strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
			os.Setenv(k, v)
		}
	}
	return nil
}
