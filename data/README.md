# Building and Deploying Kafka Broker and Zookeeper Pods on Kubernetes

This README provides step-by-step instructions for building Kafka broker and Zookeeper pods, pushing them to Docker Hub, deploying them on Kubernetes, and interacting with Kafka topics within a Kubernetes namespace.

## Docker

### Building Kafka Broker and Zookeeper Images (assuming you've logged into Docker CLI)

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Navigate to the directory containing the Dockerfiles:

    ```bash
    cd home-lab/data/kafka/zookeeper/
    ```

3. Build the Zookeeper image:

    ```bash
    docker build -t <your_dockerhub_username>/zookeeper:latest -f Dockerfile.zookeeper .
    ```

4. Push the built Zookeeper image to Docker Hub:

    ```bash
    docker push <your_dockerhub_username>/zookeeper:latest
    ```

## Kubernetes

### Creating Kubernetes Namespace and Deploying Pods

1. Create a Kubernetes namespace:

    ```bash
    kubectl create namespace kafka
    ```

2. Deploy Zookeeper pod within the namespace:

    ```bash
    kubectl run zookeeper --image=<your_dockerhub_username>/zookeeper:latest -n kafka
    ```

3. Expose Zookeeper pod as a service within the namespace:

    ```bash
    kubectl expose pod zookeeper -n kafka --port=2181 --name=zookeeper-service
    ```

4. Obtain the IP address of the Zookeeper pod:

    ```bash
    kubectl get pod -n kafka -o wide | grep zookeeper
    ```

    Note down the IP address under the "IP" column. You will use this IP address to configure the parameters of the server.properties and Dockerfile of the Kafka broker image.


------- Back to docker image creation (quick) -------

Update Kafka Dockerfile and server.properties:

 Navigate to the Kafka Dockerfile:

```bash
vi home-lab/data/kafka/Dockerfile
```

 Update the following environment variables with Zookeeper service IP:

```bash
ENV ZOOKEEPER_1_PORT_2181_TCP_ADDR XX.XX.XXX.XXX
ENV ZOOKEEPER_1_SERVICE_HOST XX.XX.XXX.XXX
ENV ZOOKEEPER_1_PORT_2181_TCP tcp://XX.XX.XXX.XXX:2181
ENV ZOOKEEPER_1_PORT tcp://XX.XX.XXX.XXX:2181
```

 Navigate to the Kafka server.properties file:

```bash
vi home-lab/data/kafka/server.properties
```

 Update `zookeeper.connect` with Zookeeper service IP:

```
zookeeper.connect=XX.XX.XXX.XXX:2181
```

### Building Kafka Broker Image and Pushing to Docker Hub

6. Build the Kafka broker image:

    ```bash
    docker build -t <your_dockerhub_username>/kafka-broker:latest -f Dockerfile.kafka-broker .
    ```

7. Push the built Kafka broker image to Docker Hub:

    ```bash
    docker push <your_dockerhub_username>/kafka-broker:latest
    ```
------ Back to Kubernetes -----------

5. Deploy Kafka broker pod within the namespace:

    ```bash
    kubectl run kafka --image=<your_dockerhub_username>/kafka-broker:latest -n kafka
    ```

6. Expose Kafka broker to port 9092:

    ```bash
    kubectl expose pod kafka -n kafka --port=9092 --name=kafka-broker-service
    ```

## Kafka

### Interacting with Kafka Topics

1. Exec into Kafka broker pod:

    ```bash
    kubectl exec -it kafka -n kafka -- /bin/bash
    ```

2. Create a Kafka topic:

    ```bash
    create_topic <topic-name>
    ```

3. List existing Kafka topics:

    ```bash
    list_topics
    ```

4. Produce messages to the topic:

    ```bash
    produce_txt <topic-name> <"Message">
    ```

5. Consume messages from the topic:

    ```bash
    consume_txt <topic-name>
    ```

You should be able to list topics from the Kafka pod if you create it on the Zookeeper pod and vice versa.

## Troubleshooting and Support

If you encounter any issues, feel free to ask questions in the comments section or reach out to me through my social media links in the video description. Your feedback is valuable!

## Documentation

For detailed documentation and configuration options, refer to the [official Kafka documentation](https://kafka.apache.org/documentation/) and [Kubernetes documentation](https://kubernetes.io/docs/).
