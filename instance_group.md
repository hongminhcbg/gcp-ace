I. 

    - group of VM intances, managed as a single entity 
    - managed instance group:
        - create by template, auto scaling, auto healing
    - unmanaged IG: diff config

II. Setup MIG

    - create Instance template (https://cloud.google.com/compute/docs/instance-templates/create-instance-templates)
    - Home -> Compute Engine -> Instance Group
    - I suggest to change max number of instances is <= 8 (https://stackoverflow.com/q/63423413)

III. Setup with cmd
    
    - $ gcloud compute instance-groups
    
