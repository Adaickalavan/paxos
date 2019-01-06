# REST Api

## Instructions

1. **Setup Docker**
    + Install Docker following the instructions for [Windows](https://docs.docker.com/docker-for-windows/) or for [Mac](https://docs.docker.com/docker-for-mac/).

2. **Source code**
    + Git clone the repository into `$GOPATH/src/paxos` folder in your computer
        ```bash
        git clone https://github.com/Adaickalavan/paxos.git $GOPATH/src/paxos
        ```
        Here, git clone syntax follows the pattern: `git clone <repo> <local folder-name>`.

3. **Run Docker-Compose to Start Application** *(requires internet connectivity)*
    + Navigate to the project folder, i.e. `$GOPATH/src/paxos/RESTApi`.
    + Start the application by running docker-compose.
        ```bash
        docker-compose up
        ```
    + When docker-compose is run in machines with Hyper-V, go to "http://localhost:8080/messages" to interact with the app.
    + When docker-compose is run using Docker Toolbox, go to "http://192.168.99.100:8080/messages" to interact with the app. Here, "192.168.99.100" is the ip address of your docker-machine. Execute `$ docker-machine ip` to get ip address of your docker-machine.

4. **Operation**
    + Several example commands and their expected outputs are given below
        ```bash
        $ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' http://localhost:8080/messages

        {
        "digest": "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
        }
        ```
        ```bash
        $ curl http://localhost:8080/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae

        {
        "message": "foo"
        }
        ```
        ```bash
        $ curl http://localhost:8080/messages/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

        {
        "err_msg": "Message not found"
        }
        ```

## Project structure

The project structure is as follows:

```txt
RESTApi
├── web                         # web microservice
|   ├── vendor                  # dependencies
|   |   ├── database
|   |   |   ├── connection.go   # database connection object
|   |   |   └── product.go      # database access object
|   |   ├── document
|   |   |   └── message.go      # database document
|   |   ├── handler
|   |   |   └── respond.go      # generic http response functions
|   |   └── ...                 # libraries for gorilla/mux and mongoDB
|   ├── Dockerfile              # Dockerfile to build `web` microservice
|   ├── Gopkg.lock              # dependency version control file
|   ├── Gopkg.toml              # dependency version control file
|   ├── handlers.go             # handlers for RESTful operation
|   └── main.go                 # main file of Go code
└── Docker-compose.yml          # to compose 2 containerized services
```

## Notes on solution

1. **Hosting domain**
   + Upon running `docker-compose up` command, the application will run in our localhost.
   + The example commands given above are to be executed in our localhost.

2. **Bottleneck and scaling with users**
   + As more users are acquired, the REST API endpoint `/messages` needs to handle large volume of requests, thus creating a bottleneck. We may handle increased load using a load-balancer such as NGINX.

3. **Application deployment and long-term maintainability**
   + Currently, the application is deployed using Docker containers so that each microservice is well contained, well defined, and system agnostic. Both the `web` (REST api) and `mongo` (MongoDB) microservice operates from within Docker containers.
   + Using Docker,
     + individual software/tools need not be installed locally, and
     + entire local development environment can be checked into source control making it easier for other developers to collaborate.
   + Multiple containers distributed over multiple servers can be orchestrated using Kubernetes.  
