IAM:
    
    Member: Who?
    Roles: Set of permissions
    Permissions: what action (list, get, delete)? what resources (app.log)?
    Policy: assign roles to members

cmd:

    gcloud compute project-info describe
    gcloud auth list
    gcloud projects get-iam-policy glowing-furnace-304608
    gcloud projects add-iam-policy-binding glowing-furnace-304608 --member=user:in28minutes@gmail.com --role=roles/storage.objectAdmin
    gcloud projects remove-iam-policy-binding glowing-furnace-304608 --member=user:in28minutes@gmail.com --role=roles/storage.objectAdmin
    gcloud iam roles describe roles/storage.objectAdmin
    gcloud iam roles copy --source=roles/storage.objectAdmin --destination=my.custom.role --dest-project=glowing-furnace-304608

