//go:generate go-bindata template/...
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type Project struct {
	Name      string
	Directory string
}

var project = &Project{}

func main() {
	// var name string
	var name string

	app := &cli.App{
		Name:  "laragen",
		Usage: "Generate a laravel project skeleton with several practical packages",
		Flags: []cli.Flag{
			// &cli.StringFlag{
			// 	Name:        "directory",
			// 	Aliases:     []string{"d"},
			// 	Usage:       "Input project directory",
			// 	Destination: &directory,
			// },
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Usage:       "Input project name",
				Destination: &name,
			},
		},
		Action: func(c *cli.Context) error {
			dir, _ := os.Getwd()
			project.Directory = dir

			if c.NArg() == 0 {
				name = dir[strings.LastIndex(dir, string(os.PathSeparator))+1:]
			}

			project.Name = name

			if err := Generate(project); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Generate(project *Project) error {
	fmt.Printf("Project name: %s, directory name: %s\n", project.Name, project.Directory)
	// Check if current directory is empty
	if isEmpty, _ := IsEmpty(project.Directory); isEmpty == false {
		return fmt.Errorf("Directory is not empty")
	}

	ExecSlowCmd("composer create-project --prefer-dist --no-install laravel/laravel ./ 5.6.*")

	Copy(
		"template/README.md",
		"README.md",
	)

	Copy(
		"template/.env",
		".env",
	)

	Copy(
		"template/composer.json",
		"composer.json",
	)
	Copy(
		"template/Helpers/functions.php",
		"app/Helpers/functions.php",
	)
	ExecSlowCmd("composer install")
	ExecCmd("php artisan key:generate")
	ExecCmd("php artisan make:auth")
	// ExecSlowCmd("composer require barryvdh/laravel-debugbar --dev")
	// ExecSlowCmd("composer require dingo/api league/fractal:0.17 tymon/jwt-auth:1.*@rc")
	ExecCmd("php artisan vendor:publish --provider=\"Dingo\\Api\\Provider\\LaravelServiceProvider\"")
	ExecCmd("php artisan vendor:publish --provider=\"Tymon\\JWTAuth\\Providers\\JWTAuthServiceProvider\"")
	ExecCmd("php artisan jwt:secret")

	fmt.Println("Add demo API")
	Copy(
		"template/Api/ApiController.php",
		"app/Http/Controllers/Api/ApiController.php",
	)
	Copy(
		"template/Api/LoginController.php",
		"app/Http/Controllers/Api/LoginController.php",
	)
	Copy(
		"template/Api/UserController.php",
		"app/Http/Controllers/Api/UserController.php",
	)
	Copy(
		"template/Requests/BaseRequest.php",
		"app/Http/Requests/BaseRequest.php",
	)
	Copy(
		"template/Requests/CreateUserRequest.php",
		"app/Http/Requests/CreateUserRequest.php",
	)
	Copy(
		"template/Requests/UpdateUserRequest.php",
		"app/Http/Requests/UpdateUserRequest.php",
	)
	Copy(
		"template/Requests/DeleteUserRequest.php",
		"app/Http/Requests/DeleteUserRequest.php",
	)
	Copy(
		"template/Transformers/UserTransformer.php",
		"app/Transformers/UserTransformer.php",
	)
	Copy(
		"template/routes/api.php",
		"routes/api.php",
	)
	Copy(
		"template/routes/web.php",
		"routes/web.php",
	)

	fmt.Println("Add php unit test")
	Copy(
		"template/config/database.php",
		"config/database.php",
	)
	Copy(
		"template/phpunit.xml",
		"phpunit.xml",
	)
	Copy(
		"template/tests/TestCase.php",
		"tests/TestCase.php",
	)
	Copy(
		"template/tests/UserApiTest.php",
		"tests/UserApiTest.php",
	)

	fmt.Println("Add ReactJS")
	Copy(
		"template/package.json",
		"package.json",
	)
	Copy(
		"template/webpack.common.js",
		"webpack.common.js",
	)
	Copy(
		"template/webpack.dev.js",
		"webpack.dev.js",
	)
	Copy(
		"template/webpack.prod.js",
		"webpack.prod.js",
	)
	Copy(
		"template/.babelrc",
		".babelrc",
	)
	ExecSlowCmd("npm install")
	Copy(
		"template/js/app.js",
		"resources/assets/js/app.js",
	)
	Copy(
		"template/js/Main.js",
		"resources/assets/js/components/Main.js",
	)
	Copy(
		"template/sass/main.scss",
		"resources/assets/sass/main.scss",
	)
	Copy(
		"template/views/react.blade.php",
		"resources/views/react.blade.php",
	)

	fmt.Println("Add Jest Test")
	Copy(
		"template/jest.config.js",
		"jest.config.js",
	)
	Copy(
		"template/tests/__mocks__/fileMock.js",
		"tests/__mocks__/fileMock.js",
	)
	Copy(
		"template/js/tests/main.test.js",
		"resources/assets/js/tests/main.test.js",
	)

	fmt.Println("Add lint")
	Copy(
		"template/.eslintignore",
		".eslintignore",
	)
	Copy(
		"template/.eslintrc.js",
		".eslintrc.js",
	)
	Copy(
		"template/phpcs.xml",
		"phpcs.xml",
	)
	Copy(
		"template/.stylelintrc",
		".stylelintrc",
	)

	return nil
}
