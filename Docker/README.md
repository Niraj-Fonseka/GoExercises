```
docker build -t dockerfonseka/dockergoapp .
```

```
docker run --rm -p 8080:8080 godockerapp
```

or 

```
docker run -d -p 8080:8080
```

### Kubernates ( minikube instructions )

#### Creating a deployment 

```
kubectl run dockergoapp --image=dockerfonseka/dockergoapp --port=8080
```


#### Exposing the service to the outside network 

```
kubectl expose deployment dockergoapp --type=LoadBalancer
```


#### Opening a webbrowser with the IP and the port 

```
minikube service dockergoapp
```