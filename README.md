# Julia

## Developer Tools

### Templ:

- https://github.com/a-h/templ
- https://templ.guide/

### Air

- https://github.com/cosmtrek/air

## Instalation

    $ go mod download

## Run Local Server

    $ air

## Project Architecture

- We are doing templating but we are doing so with HTMX templ Components

### Folders

    .
    ├── /api
    |   ├── database.go  # Calls to backend database
    ├── /static
    |   └── /images
    ├── /view
    |   ├── /components     # Templ htmx components
    |   |   ├── component.templ
    |   |   └── ...
    |   ├── /Home
    |       └── Home.templ
    ├── .air.toml        # air config file
    ├── .gitignore
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── util.go
    └── main.go

## Some useful docs

- https://htmx.org/
- https://tailwindcss.com/docs/installation
- https://go.dev/doc/
- https://gobyexample.com/
- https://echo.labstack.com/

## Developer Notes:

- Imma move the api urls to env variables
- Imma refactor some stuff like those data formatting to json communications that I kinda just did to work
- Still need to build update function
- Still need to build delete function
- Styling might change
- This is (for now) a large ass boiler plate
- .air.toml file (cmd = "templ generate && go build -o ./tmp/main .") you might need to adjust the addressing of the templ command due to templ folder installation
