package frontend

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFrontendProject(path string, library string, nameproject string, dependencies []string) {

	fmt.Println("Creating frontend project.")
	fmt.Println("Creating project folders...")
	CreateFolder(fmt.Sprintf("%s/%s", path, nameproject))
	CreateFolder(fmt.Sprintf("%s/%s/public", path, nameproject))
	CreateFolder(fmt.Sprintf("%s/%s/src", path, nameproject))

	fmt.Println("Creating project file...")
	html := CreateAppHTML(nameproject)
	CreateFile(html, fmt.Sprintf("%s/%s/public/index.html", path, nameproject))

	if library == "react" {

		packagejson := CreatePackageJSONReact(nameproject)
		CreateFile(packagejson, fmt.Sprintf("%s/%s/package.json", path, nameproject))
		App, Index := CreateFilesReact()
		CreateFile(App, fmt.Sprintf("%s/%s/src/App.js", path, nameproject))
		CreateFile(Index, fmt.Sprintf("%s/%s/src/Index.js", path, nameproject))

		if len(dependencies) > 0 {

			CreateFolder(fmt.Sprintf("%s/%s/src/store", path, nameproject))
			CreateFolder(fmt.Sprintf("%s/%s/src/router", path, nameproject))

			dependenciesstring := dependencies[0]

			for i := 1; i < len(dependencies); i++ {

				dependenciesstring = dependenciesstring + " " + dependencies[i]
			}

			fmt.Println("Project created.")
			fmt.Println("Install all dependencies with this command: ")
			fmt.Printf("npm install -S react react-dom react-scripts %s \n", dependenciesstring)
		} else {

			fmt.Println("Project created.")
			fmt.Println("Install all dependencies with this command: ")
			fmt.Println("npm install -S react react-dom react-scripts")
		}
	}
}

func CreateFolder(path string) {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {

		err := os.Mkdir(path, 0755)
		if err != nil {

			fmt.Println(err)
		}
	}
}

func CreateFile(content string, path string) {

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {

		fmt.Println(err)
	}
}

func CreateAppHTML(nameproject string) string {

	html := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>%s</title>
</head>
<body>
	<div id="app"></div>
</body>
</html>
	`

	return fmt.Sprintf(html,
		nameproject)
}
