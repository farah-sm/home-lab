# EFK Stack on Kubernetes

This README provides a guide to setting up the EFK (Elasticsearch, Fluent Bit, Kibana) stack on a Kubernetes cluster.

## Prerequisites

* A running Kubernetes cluster.
* `kubectl` configured to interact with your cluster.
* `helm` installed.

Traditionally:
![elk](https://github.com/user-attachments/assets/2cb0d189-87f0-4d10-ab6a-971a1b22e0fa)
_____________________

Today's deployment:
![efk](https://github.com/user-attachments/assets/2344983b-3e89-4b80-b275-966ad2dd851d)

Main Difference:
* Logstash is an aggregator whereas fluentbit is a forwarder, Logstash is far more customisable, it has far more richer features, allowing you to filter out the logs, advanced labelling, filtering before forwarding to elastic search. 
* Fluentbit also has similar offerings, but typically fluentbit forwards straight from the node to  elastic search.  FluentBit is far more lightwight. 


## Installation

### 1. Create a Namespace

```Bash
kubectl create namespace logging
```
2. Add Helm Repositories
```Bash

helm repo add elastic https://helm.elastic.co
helm repo update
helm repo add fluent https://fluent.github.io/helm-charts
helm repo update
```
3. Install Elasticsearch
```Bash

helm install elasticsearch \
  --set replicas=1 \
  elastic/elasticsearch -n logging
```
4. Install Kibana
   
```Bash

helm install kibana \
  --set service.type=LoadBalancer \
  elastic/kibana -n logging
```
5. Retrieve Elasticsearch Credentials
```Bash

# for username
kubectl get secrets --namespace=logging elasticsearch-master-credentials -ojsonpath='{.data.username}' | base64 -d

# for password
kubectl get secrets --namespace=logging elasticsearch-master-credentials -ojsonpath='{.data.password}' | base64 -d
```
Important: Store these credentials securely.

6. Configure Fluent Bit
   
```
Create a fluentbit-values.yaml file.
Locate the HTTP_Passwd field and replace the placeholder with the Elasticsearch password from the previous step.
```

```YAML

service:
  type: ClusterIP
  annotations: {}
env:
  # Place your Elasticsearch credentials here
  HTTP_User: "elastic"
  HTTP_Passwd: "your_elasticsearch_password" 
```
7. Install Fluent Bit

```

helm install fluent-bit \
  fluent/fluent-bit \
  -f fluentbit-values.yaml \
  -n logging
```
Accessing Kibana
1. Get the external IP of the Kibana LoadBalancer:
```
```Bash

kubectl get service kibana -n logging
```
2. Open your browser and navigate to http://<EXTERNAL-IP>:<your-port>.

3. Log in with the Elasticsearch credentials.

Clean Up

```Bash

helm uninstall fluent-bit -n logging
helm uninstall elasticsearch -n logging
helm uninstall kibana -n logging

kubectl delete namespace logging
```
Important Notes
* This is a basic setup. For production, consider:
  * Increased Elasticsearch replicas for high availability.
  * Resource limits for pods.
  * TLS encryption for secure communication.
  * Advanced log processing with Fluentd or Logstash.
* Refer to the official documentation for Elasticsearch, Kibana, and Fluent Bit for the latest information and advanced configuration options.
