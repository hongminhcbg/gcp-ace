# 36

    U want to re-architect a monolitic application, so that it follows a microservices model
    U want to accompolish this efficiently while minimizing the impact of this change to the business
    ---
    Replace the application's features with appropriate microservices in phases.

# 37

    Your existing application keep user state in a single MySQL db
    this state info is very user-specific and depends heavily on how long a user has been using the app
    The MySQL is causing challenges to maintain and enhence the schema for various user

    Which storage options should u choose?
    ---
    Cloud Datastore/Firestore [Link](https://cloud.google.com/datastore/docs/concepts/overview#what_its_good_for)

# 38

    You are building a new API. You want to minimize the cost of storing and reduce the latency of serving images.
    Which architecture should you use?
    ---
    CDN

# 40

    You are planning to deploy your application in a Google Kubernetes Engine (GKE) cluster. Your application can scale horizontally, and each instance of your application needs to have a stable network identity and its own persistent disk.
    Which GKE object should you use?
    ---
    StatefulSet: because it's own persistent disk

# 41

    Need to modify the build to exe unit test. When there is a failure
    U want to the build history to clearly error
    ---
    Create a build config file with separate cloud builder steps to compile and exe unit and intergration tests

# 42

    default svc of cloud func is  service-PROJECTA@gcf-admin-robot.iam.gserviceaccount.com

# 43

    auth service fails under intermittent load

# 44

    Application failure, want to collect application's information to troubleshoot the issue
    ---
    install log agent

# 45

    ---
    Create manual subnets.

# 46 

    which cloud svc The company should use to enable access to internal app?
    ---
    Cloud IAP

# 47

    The compaty want to remove on-call engineers and eliminate manual scaling?
    ---
    - App engine
    - Clound functions

# 48

    State is stored in a single instance MySQL database in GCP.
    But
    Business Requirements: Increase the number of concurrent users that can be supported
    ---
    move MySQL to cloud spanner

# 49

    Which service should HipLocal use for their public APIs?
    ---
    Cloud Endpoints

# 50

    Want to improve resilience of their MySQL?
    ---
    Replace the current single instance MySQL instance with Cloud SQL, and configure high availability

# 51

    Your app is running on multiple GKE cluster, how to see all log?
    ---
    gcloud logging read --resouce={...}

# 52

    U are using cloud build. Your app is built on every commit on master branch. 
    U want to release specific commits made to the master branch in an automated method?
    ---
    create a build trigger on a git tag pattern

# 53

    Move table mysql to bigtable

    Create table accounts (
        account_id (...),
        event_timestamp (...),
        amount int64,
        transaction_type
    ) primary key (account_id, event_timestap)

    How should design a row key for cloud big table
    ---
    Set Account_id_Event_timestamp as a key

# 54

    U want to view memory usage of your application deployed on CE
    ---
    install stackdriver monitoring agent because default memory metrics is not collected

# 55

    U have an analytics applications that run hundreds of queries on Bigquery. 
    U want to find out how much time there queries take to exe
    ---
    Use Stackdriver Monitoring to plot query execution times

# 56

    U are designing spanner table users,
    one user has multiple phones
    U want to search phones