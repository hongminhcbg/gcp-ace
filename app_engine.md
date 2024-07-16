App engine

    - simplest way go deploy and scale App
    - idea like deploy a docker image to cloud

Features:

    - automatic loadbalancing && auto scaling 
    - app versioning
    - traffic spliting

Env:
    
    Sandbox: for testing
    Flexible: run with docker container, prod env

Flow to deploy:
    
    App engine -> GCS -> cloud build -> App engine env

