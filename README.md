# Kubernetes 101 Workshop

This workshop is an introduction to Kubernetes. It covers the basics of Kubernetes architecture and the main resources used to deploy applications on a Kubernetes cluster.

## Workshop Agenda
- Getting familiar with kubernetes controle plane components
- Getting familiar with kubernetes workloads resources: 
    - Pods
    - Services
    - ConfigMaps
    - Secrets
    - ReplicaSet
    - Deployments

## Prerequisites
- [Docker](https://docs.docker.com/get-started/get-docker/) 
- [Kind](https://kind.sigs.k8s.io/#installation-and-usage)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)


## Build the docker image

``` bash
docker build -t go-server:v1 . 
```

## Run container 
```bash
docker run -d -p 8080:8080 -e NAME=ENICarthage -v ./secret.yaml:/var/secrets/secret.yaml go-server:v1
``` 

## Create a kind cluster
``` bash
kind create cluster --config ./k8s/0-kind-config.yaml
```
## Exec to the kind control plane
``` bash
export id=$(docker ps | grep control-plane | cut -d " " -f1)
docker exec -it $id sh
# crictl ps -a
# ls /etc/kubernetes/manifests
# systemctl status kubelet
```

## Load a local docker image into the kind cluster
``` bash
kind load docker-image go-server:v1 --name enicarthage
```

## Kubernetes Commands

```bash
kubectl config get-contexts
cat ~/.kube/config

# Get the list of nodes
kubectl get nodes

# Get the list of pods
kubectl get pods --all-namespaces

# Set the current namespace to demo
kubectl config set-context --current --namespace demo

# Create a pod with the specified image and environment variable
kubectl run mypod --image=go-server:v1 --env="NAME=ENICarthage" --dry-run=client -o yaml

# Describe the created pod
kubectl describe pod mypod 

# Get the logs of the pod
kubectl logs mypod

# Expose the pod as a service on port 8080
kubectl expose pod mypod --port=8080 

# CURL the service 
kubectl run alpine --image=alpine -- sleep 1h
kubectl exec -it alpine -- sh
apk add curl
curl mypod:8080/hello
cat /etc/resolv.conf

kubectl delete pod mypod --force
kubectl apply -f ./k8s/4-pod-with-secret.yaml
kubectl get secret mysecret -o yaml
echo "UGFzc3dvcmQ6IE15UEBBc3N3MHJk" | base64 -d
output: Password: MyP@ssw0rd

# Get the pod details in YAML format
kubectl get pod mypod -o yaml 

# Create a secret from a file
kubectl create secret generic mysecret --from-file=./secret.yaml

# Create a ConfigMap with a specified literal value
kubectl create cm myconfig --from-literal=NAME=ENICarthage
kubectl get cm myconfig -o yaml

# Creating a Pod with non existing configMap
kubectl  apply -f 6-pod-with-cm-fail.yaml
kubectl get po 
Output:
mypod          0/1     CreateContainerConfigError   0          2s

kubectl describe po mypod

Output: 
Warning  Failed     3m46s   kubelet            Error: configmap "wrongname" not found

# Apply a deployment configuration from a YAML file
kubectl apply -f ./k8s/7-deployment.yaml

# Get the list of deployments
kubectl get deploy

# Get the list of events
kubectl get events
```

## Clean up
```bash
kind delete cluster --name enicarthage
```