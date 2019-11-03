"# go-kit-testing" 

Simple microservice where you can POST to a HTTP server to create a Person. A Person consists of a UUID and a name. 

You can retrieve a Person once it has been created by a GET method that includes the id of the Person you are searching for.

The repository used is a testing Postgres DB that is containerized within Docker.

Uses mux for the router to create the HTTP server and Go-Kit for endpoints, transport, etc. 