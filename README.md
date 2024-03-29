## How to run it

### Setup Docker
Download and install Docker Desktop from [the offical website](https://www.docker.com/products/docker-desktop/) or install it from command line with brew.

### Run the containers
Navigate to the root folder of the project and run `docker-compose up` to start the docker image.

### Test the node server 
Navigate in the browser to http://localhost:3000 to test if the node server is running.

### Test the go server 
Navigate in the browser to http://localhost:8080 to test if the go server is running.

### Kafka
To check the messages on a topic, connect to the kafka container terminal `docker-compose exec kafka /bin/bash`. 
After that, use `kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic my-topic` to consume the messages from my-topic. When you navigate to the http://localhost:8080/kafka a message will be produced to this topic.

### Run react web app
Navigate to web-app and `npm run dev`

### Test react app 
Navigate in the browser to http://localhost:3000 

## More info
- Live reload of the go code
- Initializing scripts for database