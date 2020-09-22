package javascript

import (
	"fmt"
	"os"

	"github.com/Santiagozh1998/create-project/javascript/frontend"
)

func GetDependenciesFromArgs(array []string) []string {

	reverse := []string{}

	for i := (len(array) - 1); i >= 0; i-- {

		reverse = append(reverse, array[i])
	}

	reverse = reverse[:len(reverse)-4]

	return reverse
}

func FrontendOrBackendProject(dir string, array []string) {

	fmt.Println(array)
	whatIs := array[2]

	if whatIs == "react" || whatIs == "vue" {

		dependencies := []string{}

		if len(os.Args) > 3 {

			dependencies = array
			dependencies = GetDependenciesFromArgs(dependencies)
		}
		fmt.Println(dependencies)

		frontend.CreateFrontendProject(dir, whatIs, array[3], dependencies)
	}
}
