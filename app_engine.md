App engine

    - simplest way go deploy and scale App
    - idea like deploy a docker image to cloud
    - only one app for each project

Features:

    - automatic loadbalancing && auto scaling 
    - app versioning
    - traffic spliting

Env:
    
    Sandbox: for testing
    Flexible: run with docker container, prod env

Flow to deploy:
    
    App engine -> GCS -> cloud build -> App engine env

Repo structure
```yaml
root:
    - dispatch.yaml
    - cron.yaml
    svc1:
        - src
        - service1.yaml
        - app.yaml
    svc2:
        - src
        - service2.yaml
        - app.yaml
```
