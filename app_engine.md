App engine

    - simplest way to deploy and scale App
    - idea like deploy a docker image to cloud
    - only one app for each project
    - simple microservices

Features:

    - automatic loadbalancing && auto scaling 
    - app versioning
    - traffic spliting, default mode is split by IP

Env:

  - Sandbox: for testing
  - Flexible: run with docker container, prod env

Stardard vs flexible:

  - pricing: Hours vs memory * cpu * disk
  - scaling: Manual,basic,auto vs manual,auto
  - scale to zero: yes vs minimun one instance
  - startup time: seconds vs minutes
  - ssh to debug: no vs yes because based on VM

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

cron:

  cron.yaml

```
gcloud app deploy con.yaml

```

