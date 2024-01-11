# Simplemedia
SimpleMedia is a social media API that allows users to register, connect with friends, create posts, and engage in various social activities. Below, you'll find a comprehensive list of available endpoints for different operations.

## Project Setup and Usage

### Requirements

##### Go 1.21.5
##### Docker and Docker Compose installed on your system.
##### A Postman collection has been included for convenient API testing. Import the collection to explore and interact with the API endpoints.

### Make Commands:
- all: Builds and starts both the application and database services.
- down: Stops and removes the running containers.
- re: Cleans up all containers and volumes, then rebuilds and starts everything from scratch.
- clean: Removes all containers, volumes, and unused Docker images for a thorough cleanup. " Warning: This will remove all Docker images on your system."
- db: Starts only the PostgreSQL database service.
- app: Starts only the application service.

## Endpoints
```http
POST    /register
POST    /login
POST    /user/logout
PUT     /user/pwd
PUT     /user/picture
        
POST    /comment/create
PUT     /comment/update
DELETE  /comment/delete
GET     /comment/get
GET     /comment/list
        
POST    /friends/add
POST    /friends/accept
DELETE  /friends/decline
DELETE  /friends/delete
GET     /friends/request
GET     /friends/list
        
POST    /likes/create
DELETE  /likes/delete
GET     /likes/getall
        
POST    /post/create
PUT     /post/update
GET     /post/get
DELETE  /post/delete
GET     /post/getall

```

#### Additional Information:
The PostgreSQL database data is persisted in the /home/postgres-data directory on your host machine.
To rebuild specific services, use docker-compose up --build <service-name> (e.g., docker-compose up --build app).
