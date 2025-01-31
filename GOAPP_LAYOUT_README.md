Structuring Go projects can be approached in several ways, but here's a common and effective structure: 

Project Root 

• cmd/: Contains the main application entry points. Each subdirectory within cmd/ represents a different executable. 
	• app1/: Contains main.go for the first application. 
	• app2/: Contains main.go for the second application. 


• internal/: Contains packages that are only used within this project. 
	• config/: Configuration management logic. 
	• database/: Database interaction code. 
	• models/: Data models and structures. 
	• services/: Business logic and service implementations. 


• pkg/: Contains packages that can be reused in other projects. 
	• util/: Reusable utility functions. 
	• logging/: Logging utilities. 


• go.mod: Module definition file. 
• go.sum: Module dependency checksums. 
• README.md: Project description and documentation. 

Key Points 

• cmd/: This directory is for executable binaries, and each binary should have its own subdirectory. 
• internal/: Packages in this directory are private to the project and cannot be imported by other projects. 
• pkg/: Packages here are intended to be reusable in other projects. 
• Subdirectories: Subdirectories within internal/ and pkg/ can be organized by domain or functionality. 
• Testing: Test files are typically placed alongside the source files, following the *_test.go naming convention. 

Example 
project-root/
├── cmd/
│   ├── app1/
│   │   └── main.go
│   └── app2/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   └── services/
│       └── user.go
├── pkg/
│   └── util/
│       └── helpers.go
├── go.mod
├── go.sum
└── README.md

Benefits 

• Clear Separation: Separates application entry points, project-specific code, and reusable packages. 
• Maintainability: Makes it easier to find and modify specific parts of the project. 
• Reusability: Promotes the creation of reusable packages. 
• Testability: Facilitates unit testing and integration testing. 

--------------------------------------------------------
Another example that has more of a production feel
--------------------------------------------------------

mywebapp/
    ├── cmd/
    |   ├── app/
    |   |   └── main.go
    ├── internal/
    |   ├── api/
    |   |   ├── http/
    |   |   |   ├── admin/
    |   |   |   |   ├── admin.go
    |   |   |   |   ├── admin_test.go
    |   |   |   ├── middleware/
    |   |   |   |   ├── middleware.go
    |   |   |   |   ├── middleware_test.go
    |   |   |   ├── server/
    |   |   |   |   ├── server.go
    |   |   |   |   ├── server_test.go
    |   |   |   |   └── router.go
    |   ├── config/
    |   |   ├── config.go
    |   ├── util/
    |   |   └── util.go
    ├── pkg/
    |   ├── admin/
    |   |   ├── repository/
    |   |   |   ├── admin_repository.go
    |   |   |   ├── admin_repository_test.go
    |   |   ├── service/
    |   |   |   ├── admin_service.go
    |   |   |   └── admin_service_test.go
    |   |   ├── admin.go
    |   |   └── admin_test.go
    ├── scripts/
    |   ├── build.sh
    |   ├── run.sh
    |   ├── test.sh
    ├── deployment/
    |   ├── Dockerfile
    |   ├── docker-compose.yml
    |   └── kubernetes.yml
    ├── .env
    ├── Makefile
    ├── README.md
    ├── go.mod
    └── go.sum
