## Setup Docker
Download and install Docker Desktop from [the offical website](https://www.docker.com/products/docker-desktop/) or install it from command line with brew.

## Run the project
Navigate to the root folder of the project and run `docker-compose up` to start the docker image.

## Test it 
Navigate to in the browser to http://localhost:8080/headers to test if the server is running.

## Live reload
Dockerfile is using CompileDaemon, the changes you make in the go files will be live-reloaded on the server inside the container.