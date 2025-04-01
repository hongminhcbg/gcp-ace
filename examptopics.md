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
    U want to search phones?
    ---
    create parent table 'users', create child table 'phones' interleave with field customer_id, create index phone_number

# 57

    You are deploying a single website on App Engine that needs to be accessible via the URL http://www.altostrat.com/.
    What should you do?
    ---
    Verify domain ownership with Webmaster Central. Create a DNS CNAME record to point to the App Engine canonical name ghs.googlehosted.com.

# 58

    You are running an application on App Engine that you inherited. You want to find out whether the application is using insecure binaries or is vulnerable to XSS attacks.
    Which service should you use?
    ---
    Cloud security scanner

# 59

    U are working on a social media, user want to upload images.
    There img will be 2MB
    ---
    Change the application to create signed URLs for Cloud Storage. Transfer these signed URLs to the client application to upload images to Cloud Storage.

# 60

    Your application is built as a custom machine image. You have multiple unique deployments of the machine image. Each deployment is a separate managed instance group with its own template. Each deployment requires a unique set of configuration values. You want to provide these unique values to each deployment but use the same custom machine image in all deployments. You want to use out-of-the-box features of Compute Engine.
    What should you do?
    ---
    Place the unique configuration values in the instance template instance metadata.

# 61

    Your app performs well when tested locally, but it runs sinificantly slower after u deploy it to CE instance. U need to diagnose the problem
    What should u do
    ---
    Use Cloud Profiler to determine which functions within the application take the longest amount of time.

    Because cloud profiler is a performance analysis tool that allows u to identify per issue

# 63

    Your company has a dataware house, that keep your app information in BQ
    The BigQuery data warehouse keeps 2 PBs of user data. Recently, your company expanded your user base to include EU users and needs to comply with these requirements:
        1. Your company must be able to delete all user account information upon user request
        2. All EU user data must be stored in a single region specifically for EU users.
    ---
    - Create a dataset in the EU region that will keep information about EU users only.
    - Use DML statements in BigQuery to update/delete user records based on their requests.
# 64

    Your App Engine standard configuration is as follows:
    service: production
    instance_class: B1
    You want to limit the application to 5 instances.
    Which code snippet should you include in your configuration?
    ---
    basic_scaling: 
        max_instances: 5
# 65

    Your analytics system executes queries against a BigQuery dataset. The SQL query is executed in batch and passes the contents of a SQL file to the BigQuery
    CLI. Then it redirects the BigQuery CLI output to another process. However, you are getting a permission error from the BigQuery CLI when the queries are executed.
    You want to resolve the issue. What should you do?
    ---
    Grant the service account BigQuery Data Viewer and BigQuery Job User roles

# 66

    Your app in running on CE and it showing subtained failures for a small number of request
    U can narrowed the cause down to a single CE instance, but u can't ssh
    ---
    enable and check serial port output

# 67

    You configured your CE instance group to scale automatically acroding to overral CPU usage
    Howrever, your app's response latency increases sharply before the cluster has finished adding up instances

    Which two configuration changes should u make? 
    ---
    - Decrease the target CPU usage for the instance group autoscaler.

    - Decrease the cool-down period for instances added to the group. (LINK)[https://cloud.google.com/compute/docs/autoscaler#cool_down_period]
    - Because decrease the cool-down period for instances added to the group: The cool-down period is the time the autoscaler waits after a new instance is healthy before it collects usage metrics from the instance. A shorter cool-down period allows the autoscaler to react more quickly to changes in load, potentially starting to scale up sooner when there is a sudden increase in traffic.

# 68

    U have an application controlled by a MIG
    U deploy a new version on application, the number of instances should not increase
    ---
    Maxsurge = 0, max unavalable = 1 

# 69

    Your application requires service accounts to be authenticated to GCP products via credentials stored on its host Compute Engine virtual machine instances. You want to distribute these credentials to the host instances as securely as possible.
    What should you do?
    ---
    Use the instance's service account Application Default Credentials to authenticate to the required resources

# 70

    Your application is deployed in a Google Kubernetes Engine (GKE) cluster. You want to expose this application publicly behind a Cloud Load Balancing HTTP(S) load balancer.
    What should you do?
    ---
    Configure a GKE Ingress resource with type: LoadBalancer.

# 71

    Your company is planning to migrate their on-premises Hadoop environment to the cloud. Increasing storage cost and maintenance of data stored in HDFS is a major concern for your company. You also want to make minimal changes to existing data analytics jobs and existing architecture.
    How should you proceed with the migration?
    ---
    Create a Cloud Dataproc cluster on Google Cloud Platform, and then migrate your Hadoop code objects to the new cluster. Move your data to Cloud Storage and leverage the Cloud Dataproc connector to run jobs on that data

# 72

    Your data stored in GCS bucket
    Fellow developers have reported that data downloaded from GCS is resulting in slow API performance
    U want to research the issue to provide detailsto GCP team
    ---
    gsutil perfdiag: Run performance diagnostic

# 73

    You are using Cloud Build build to promote a Docker image to Development, Test, and Production environments. You need to ensure that the same Docker image is deployed to each of these environments.
    How should you identify the Docker image in your build?
    ---
    Use the digest of the Docker image

# 74

    Your company has created an application that uploads a report to a Cloud Storage bucket. When the report is uploaded to the bucket, you want to publish a message to a Cloud Pub/Sub topic. You want to implement a solution that will take a small amount to effort to implement.
    What should you do?
    ---
    Create a Cloud Function that is triggered by the Cloud Storage bucket. In the Cloud Function, publish a message to the Cloud Pub/Sub topic.
    
# 76

    Your company stores their source code in a Cloud Source Repositories repository. Your company wants to build and test their code on each source code commit to the repository and requires a solution that is managed and has minimal operations overhead.
    Which method should they use?
    ---
    Use Cloud Build with a trigger configured for each source code commit.

# 77

    You are writing a Compute Engine hosted application in project A that needs to securely authenticate to a Cloud Pub/Sub topic in project B.
    What should you do?
    ---
    Configure the instances with a service account owned by project A. Add the service account as a publisher on the topic.

# 78

    You are developing a corporate tool on Compute Engine for the finance department, which needs to authenticate users and verify that they are in the finance department. All company employees use G Suite.  
    What should you do?
    ---
    Enable Cloud Identity-Aware Proxy on the HTTP(s) load balancer and restrict access to a Google Group containing users in the finance department. Verify the provided JSON Web Token within the application.

# 79

    
