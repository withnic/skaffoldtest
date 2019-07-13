

* Need
  * docker, k8s, container-structure-test
 
 
 
 ## Run local
 
 * install docker.
 
 * Switch k8s setting on docker.
 
 * install [container-structure-test](https://github.com/GoogleContainerTools/container-structure-test)
 
 
```
# watch files and build docker image
$skaffold dev
```

curl http://localhost:8080 or see browser
 
 
 ## TEST (container-structure-test)

 ```
$skaffold build --profile test
 ```