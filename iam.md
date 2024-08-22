IAM:
    
    Member: Who?
    Roles: Set of permissions
    Permissions: what action (list, get, delete)? what resources (app.log)?
    Policy: assign roles to members or group

cmd:

    gcloud compute project-info describe
    gcloud auth list
    gcloud projects get-iam-policy glowing-furnace-304608
    gcloud projects add-iam-policy-binding glowing-furnace-304608 --member=user:in28minutes@gmail.com --role=roles/storage.objectAdmin
    gcloud projects remove-iam-policy-binding glowing-furnace-304608 --member=user:in28minutes@gmail.com --role=roles/storage.objectAdmin
    gcloud iam roles describe roles/storage.objectAdmin
    gcloud iam roles copy --source=roles/storage.objectAdmin --destination=my.custom.role --dest-project=glowing-furnace-304608

sa scenarios:

    VM want to publish msg, want to push obj to buckets => create a sa, config the vm use the sa
    VM in project A, want to access buckets in project B => in project B, add default sa from project A, assign per Storage Bucket Viewer

ACL:

    define who has access to your buckets, as well as what level of access they have.
    ACL vs IAM

    IAM apply to mass obj (if have object view => view all objs)
    ACL can be used to customized specific access to job
    
    Permission is allowed by IAM but NOT by ACL. Will the user be able to access the object?
    - Yes, u need either IAM or ACL

Roles:
    
    - basic role: Owner, Viewer, Editor
        viewer: read-only
        editor: viewer + edit
        owner: editor + IAM + billing
    - predefine-role:
        - fine granted roles and managed by google
        - example: storage object admin, storage object viewer, 
    - custom roles:
        - custom anything

On prem:
    
    - expose service account as json

Predefine Role:
    
    OrganizationAdmin: 
        - define resources hierarchy
        - define access management policies
        - manege users and groups
    BillAccountCreator:
        - create billing account
        - finance teams
    BillAccountAdmin:
        - manege bill account (active, inactive, export, link, unlink)
        - canot create billAccount
        - finance teams
    BillAccountUser:
        - Assign project to billAccount
        - project owner
    BillAccountViewer:
        - view details
        - auditor
    AppEngineAdmin: App(RU)
    AppEngineCreator: App(CD), svc+instance+versions(CRUD)
    AppEngineViewer: app+svc+instance+versions(R)
    AppEngineCodeViewer: versions.FileContent(R)
    StorageAdmin: storage.bucket.*, storage.object.*
    StorageObjectAdmin: storage.object.* (DOES NOT HAVE storage.bucket.*)
    StorageObjectCreator: storage.object.* (C) 
    BigQueryAdmin: bigquery.* 
    BigQueryDataOwner: bigquery.datesets.*, bigquery.tables.*, bigquery.models.* (DOES NOT HAVE bigquery.jobs.*)
    BigQueryJobUser: bq.jobs.* (C)
    BigQueryUser: BigQueryDataViewer + bq.jobs.*(R)
    BigQueryDataViewer: read data and metadate from table, can't create job
    CloudSqlAdmin: Full sql resource
    CloudSqlClient: connect to sql instance
    PrivateLogViewer: View log and private log, data access log is private log
        

Best practices:

    - basic role not recommended, prefer predefine role when possible
    - Use SA with minimum previleges
    - Google Acount: a person
    - Service Account: an application account
    - Group: a group of people or SAs
    - Organization policy: Forcus on WHAT, what can be done on specific resources (example restrict public ip on SQL)
    - user has GSuite account => grant required role to the project is the correct way to give them access
    - grant IAM with oganization level => give them acc all projects => Risk security

![Alt text](./imgs/edit_role.jpeg?raw=true "Title")
