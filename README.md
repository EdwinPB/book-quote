# Book Quote
## Stack

The stack of the application is:

- Go for the backend
- Go Templ for the frontend
- Tailwind CSS for styling
- Posgresql for storage
- HTMX for frontend interactions


## Setup local development

### Templ
Install templ in your machine having installed Go

```
go install github.com/a-h/templ/cmd/templ@latest
```

### Make
You will need make to run the local environment, if you don't have it installed in your machine, run the following command

```
brew install make
```

### Database
In order to run the app in development you should create the database to storage your data permanently.

```
go run ./cmd/db create
```

and run migrations.

```
go run ./cmd/db migrate
```

### Frontend dependencies
You need to run a our setup step to install the tailwindCSS standalone tooling. 
```
go run ./cmd/setup
```

### Running
To run the app in the development environment you will have to have two terminals open. \
In one of them run first: \
This is to run the watch which will be listening for the changes in your templ files

```
make templ-watch
```

Then, in the other terminal run: \

```
make dev
```


And visit http://localhost:3000 to see the app running.