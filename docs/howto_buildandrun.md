# Getting started

## Build and run this project

1. clone the repository
2. Open a terminal and browse to /cmd
3. Type `go run .\main.go`
4. The service starts up and listen on `http://localhost:3000/api/v1/notifications`

The repository comes with one build-in receiver called GithubIssueReceiver. This receiver creates github issues whenever 
notifications of type `Warning` are sent to the API. Therefore, the following environment variables must be set:

> [!IMPORTANT]
> GithubIssueReceiver required environment variables
> - GITHUB_OWNER
> - GITHUB_REPO
> - GITHUB_TOKEN

## Run unit and integration tests
2. Open a terminal and browse to the root folder of the repository
3. Type `go test ./...` to run all tests in the repository


## Folder and package structure

`/cmd`
Contains the main entry point to start the http server for development purposes

`/tests`
Contains integration tests that span multiple components

`/docs`
Contains the documentation of this project

`/api`
Contains the notification service OpenApi v3 specification file

`/internal/application`
Contains the application itself

`/internal/messaging`
Simple In-Memory messaging component to dispatch incoming API calls to receivers

`/internal/routes`
Handler for /api/v1/notifications endpoint

`/internal/receivers`
Implementation of concrete receivers. E.g. GithubIssueReceiver









