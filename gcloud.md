# gcloud command line interface
    - most GCP svc can be managed by gcloud interface
some gcp svc have specific tool

    - cloud storage: gsutil
    - cloud big query: bq
    - cloud bigtable: cbt

command structure:

    gcloud GROUP SUBGROUP ACTION ...
    example command to create a VM

    gcloud compute instances create my-instance \
    --boot-disk-size=50GB \
    --zone=asia-southeast2-c \
    --machine-type=n1-standard-2 \
    --boot-disk-size=50GB \
    --image ubuntu-1804-bionic-v20190204 \
    --image-project ubuntu-os-cloud
    
    // add default zone and region
    $ gcloud compute project-info add-metadata --metadata=google-compute-default-region=asia-southeast1
    $ gcloud compute project-info add-metadata --metadata=google-compute-default-region=asia-southeast1,google-compute-default-zone=asia-southeast1-a
    
    // list image
    $ gcloud compute images list

Manage Instance Group (MIG):

    - managed as one unit

MIG scenarios:

    - want MIG managed app to survive zone failures => create multi zone MIG
    - want to create VMs of different config => create un-managed intance group 


