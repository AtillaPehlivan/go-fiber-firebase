# Backend API Example

FiberGo Framework Docs : https://github.com/gofiber

### Info
This application using firebase ecosystem 
- Firebase Auth
- Cloud Storage
- Firestore 

## Development

### Start the application

```bash
go run main.go
```

### please fill the .env file for production


## Production

```bash
docker build -t mockerize-backend .
docker run -d -p 8080:8080 mockerize-backend
```

Go to `http://localhost:8080`:



## Architecture
Clean Architecture is a concept introduced by 
Robert C. Martin or also known as Uncle Bob. Simply put, the purpose of this architecture is to perform complete separation of concerns. Systems made this way can be independent of frameworks, testable (easy to write unit tests), independent of UI, independent of database, and independent of any external agency. When you use this architecture, it is simple to change the UI, the database, or the business logic.

<img alt="System Architecture" src="https://raw.githubusercontent.com/gofiber/recipes/master/docker-mariadb-clean-arch/assets/CleanArchitecture.jpg"/>

| Architecture Layer  |    Equivalent Layer    |             Filename             |
| :-----------------: | :--------------------: | :------------------------------: |
| External Interfaces | Presenters and Drivers | `middleware` and `routes` |
|     Controllers     |     Business Logic     |           `service.go`           |
|      Use Cases      |      Repositories      |         `repository.go`          |
|      Entities       |        Entities        |           `model.go`            |

Basically, a request will have to go through `route.go` (and `middleware.go`) first. After that, the program will call a repository or a use-case that is requested with `service.go`. That controller (`service.go`) will call `repository.go` that conforms to the `model.go` in order to fulfill the request that the `service.go` asked for. The result of the request will be returned back to the user by `route.go`.
