
# EFK Stack on Kubernetes with ECK

Elastic Cloud on Kubernetes (ECK) simplifies the deployment and management of Elasticsearch and Kibana on Kubernetes.  It addresses Day 2 operations challenges like upgrades, scaling, and configuration management by automating these tasks through the Kubernetes operator pattern, ensuring consistent and reliable operation of your Elasticsearch cluster.

## Prerequisites & Assumptions

* Helm
* Kubectl

## Installation and Configuration

### 1. Install the ECK Operator & Fluent Bit agent(s)

```bash
kubectl create ns observer

helm repo add elastic https://helm.elastic.co
helm repo add fluent https://fluent.github.io/helm-charts

helm install elastic-operator elastic/eck-operator -n observer --kubeconfig /path/to/your/kubeconfig  # Replace with your kubeconfig path


```

### 2. Deploy Elasticsearch Cluster

```yaml
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: quickstart-es
  namespace: obs-system
spec:
  version: 8.17.1 # Use your desired version
  nodeSets:
  - name: default
    count: 1
    config:
      node.store.allow_mmap: false # Important for some environments
  http:
    service:
      spec:
        type: LoadBalancer # Or NodePort if LoadBalancer not available
```

```bash
kubectl apply -f elasticsearch.yaml # Save the above YAML as elasticsearch.yaml
```

### 3. Deploy Kibana Instance

```yaml
apiVersion: kibana.k8s.elastic.co/v1
kind: Kibana
metadata:
  name: quickstart-kb
  namespace: obs-system
spec:
  version: 8.17.1 # Use your desired version
  count: 1
  elasticsearchRef:
    name: quickstart-es
  http:
    service:
      spec:
        type: NodePort # Or LoadBalancer if available and preferred
```

```bash
kubectl apply -f kibana.yaml # Save the above YAML as kibana.yaml
```

## Data Ingestion

### 4. Configure Fluent Bit values file

Populate the ConfigMap section of the Fluent Bit values file for Fluent Bit output configuration:

```yaml

data:
  output.conf: |
    [OUTPUT]
      Name             es
      Match            *
      Host             <elasticsearch_service_ip> # Replace with Elasticsearch service IP
      Port             9200
      tls.verify       Off # Security risk - address in future iterations
      tls.debug        3
      Index            fluentbit-forwarder
      Logstash_Format  Off
      tls.ca_file      /fluent-bit/tls/tls.crt
      tls.crt_file     /fluent-bit/tls/tls.crt
      HTTP_User        elastic
      HTTP_Passwd      <elasticsearch_password> # Security risk - address in future iterations
      tls              On
      Suppress_Type_Name On
```

```bash
helm install fluent-bit \
  fluent/fluent-bit \
  -f fluentbit-values.yaml \
  -n observer # Save the above YAML as fluent-bit-config.yaml
```

**Important:** Replace `<elasticsearch_service_ip>` with the actual ClusterIP or LoadBalancer IP of your Elasticsearch service.  Replace `<elasticsearch_password>` with the Elasticsearch `elastic` user's password.  **Do not hardcode passwords in ConfigMaps in production.**  See "Security Considerations" below.

### 5. Mount Certificates via Secret in DaemonSet Manifest

Mount the certificates stored in the `quickstart-es-http-certs-public` secret directly into the Fluent Bit daemonset.  This is done within the daemonset's pod spec (consult your log monitoring system's documentation for how to modify the daemonset):

```yaml
volumeMounts:
- mountPath: /fluent-bit/tls
  name: tls-certs
  readOnly: true # Important for security
volumes:
- name: tls-certs
  secret:
    secretName: quickstart-es-http-certs-public
```

This configuration mounts the `tls.crt` (and any other files) within the `quickstart-es-http-certs-public` secret at `/fluent-bit/tls` inside the Fluent Bit container.  Ensure that the `tls.ca_file` and `tls.crt_file` paths in your Fluent Bit ConfigMap (`output.conf`) correctly point to the files within this mounted directory (e.g., `/fluent-bit/tls/tls.crt`).


Wait for the `fluent-bit` pods (or equivalent) to restart.

## Visualization and Analysis

### 6. Port-forward Kibana (if necessary)

If Kibana is deployed with a `ClusterIP` service, you'll need to port-forward to access it locally:

```bash
kubectl port-forward svc/quickstart-kb-http 5601:5601 -n obs-system --kubeconfig /path/to/your/kubeconfig
```

### 7. Access Kibana UI

Get the external IP or NodePort of the Kibana service:

```bash
kubectl get svc quickstart-kb-http -n obs-system --kubeconfig /path/to/your/kubeconfig
```

Access Kibana UI: `http://<kibana_ip>:<kibana_port>/login?`

### 8. Query Elasticsearch (for testing)

```bash
kubectl exec -it quickstart-es-default -- bash --kubeconfig /path/to/your/kubeconfig
curl -k https://<elasticsearch_service_ip>:9200/fluentbit-forwarder/_search?pretty=true -u "elastic:<elasticsearch_password>"
```

## Security Considerations (Caveats)

This example deployment prioritises functionality over comprehensive security. The following security considerations will be addressed in future iterations:

* **`tls.verify` is Off:** In a production environment, `tls.verify` should be enabled (`On`) to ensure the integrity of TLS communication. Future work will configure proper certificate validation.
* **Credentials in ConfigMap:** The Elasticsearch password (`HTTP_Passwd`) is stored in plain text within the ConfigMap. This is a significant security risk. Future iterations will use Kubernetes Secrets or a dedicated secrets management solution to store and manage sensitive credentials.

```
