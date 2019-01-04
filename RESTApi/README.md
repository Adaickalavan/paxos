# REST Api

## Instructions

1. **Setup Docker**
    + Install Docker following the instructions for [Windows](https://docs.docker.com/docker-for-windows/) or for [Mac](https://docs.docker.com/docker-for-mac/).

2. **Source code**
    + Git clone repository into `$GOPATH/src/paxos` folder in your computer

3. **Run Docker-Compose to Start Application** *(requires internet connectivity)*
    + In a terminal, navigate to the project folder, i.e. `C:/goWorkspace/src/paxos/RESTAPI`. 
    + Start the application by running docker-compose.
        ```bash
        docker-compose up
        ```

4. **Operation**
    + Several example commands and their expected outputs are given below
        ```go
        $ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' http://localhost:8080/messages

        {
        "digest": "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
        }
        ```
        ```go
        $ curl http://localhost:8080/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae

        {
        "message": "foo"
        }
        ```
        ```go
        $ curl http://localhost:8080/messages/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

        {
        "err_msg": "Message not found"
        }
        ```

## Project structure

The project structure is as follows:

```txt
RESTApi
├── vendor                  # dependencies
|   ├── database
|   |   ├── connection.go   # database connection object
|   |   └── product.go      # database access object
|   ├── document
|   |   └── message.go      # database access object
|   ├── handler
|   |   └── respond.go      # generic http response functions
|   └── ...                 # libraries for gorilla/mux and mongoDB
├── Docker-compose.yml      # to compose 2 containerized services
├── Dockerfile              # Dockerfile to build `restapi` api image
├── Gopkg.lock              # dependency version control file
├── Gopkg.toml              # dependency version control file
├── handlers.go             # handlers for RESTful operation
└── main.go                 # main file of Go code
```

## Notes on solution

1. **Hosting domain**
   + Upon running `docker-compose up` command, the application will run at your localhost.
   + The example commands given above are to be executed at our localhost.

2. **Alternative hosting domain**
   + This application is also hosted at heroku.
   + Example commands for Heroku are as follows:

3. **Bottleneck as more users are acquired**

4. **How you might scale your microservice**

5. **Application deployment process for long term maintainability**
