# Testing Simple docker application with minikube

https://kubernetes.io/docs/tutorials/hello-minikube/


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


### Updating a new image 

```
Update the code and make sure it runs 
```

#### Build the new image

```
docker build -t dockerfonseka/dockergoapp:v2 .
```

#### Push to docker hub 

```
docker push dockerfonseka/dockergoapp
```

#### Update the kubernates image with the new image 

```
kubectl set image deployment/dockergoapp dockergoapp=dockerfonseka/dockergoapp:v2
```