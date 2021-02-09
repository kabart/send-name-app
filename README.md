Build docker image:

```
docker build --tag kabartkiew/send-name-app .
```

Push docker image to Docker Hub:

```
docker push kabartkiew/send-name-app:latest
```

Run docker image:

```
docker run -p 8080:8080 --name send-name-app kabartkiew/send-name-app
```

Run app in minikube:
```
minikube start
kubectl apply -f deploy-send-name.yaml
kubectl apply -f svc-send-name.yaml
minikube service svc-send-name
```
