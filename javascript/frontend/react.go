package frontend

import "fmt"

func CreatePackageJSONReact(nameproject string) string {

	packagejson := `{
	"name": "%s",
	"version": "1.0.0",
	"scripts": {
		"start": "react-scripts start",
		"build": "react-scripts build",
		"test": "react-scripts test",
		"eject": "react-scripts eject"
	}
}`

	return fmt.Sprintf(packagejson,
		nameproject)
}

func CreateFilesReact() (string, string) {

	App := `import React from "react";

const App = () => {
	return(
		<div>
			<h1>React project</h1>
		</div>
	)
}

export default App;`

	Index := `import React from "react";
import ReactDOM from "react-dom";
import App from "./App";

ReactDOM.render(
	<React.strictMode>
		<App />
	</React.strictMode>,
	document.getElementById("app")
);`
	return App, Index
}
