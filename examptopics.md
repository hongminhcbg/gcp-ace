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

    Your API backend is running on multiple cloud providers. You want to generate reports for the network latency of your API.
    Which two steps should you take? (Choose two.)
    ---
    - Use Zipkin collector to gather data.  (https://cloud.google.com/trace/docs/zipkin)
    - Use Stackdriver Trace to generate reports.

# 80 -> 83 case study
    Business Requirements -
        HipLocal's investors want to expand their footprint and support the increase in demand they are seeing. Their requirements are:
        * Expand availability of the application to new regions.
        * Increase the number of concurrent users that can be supported.
        * Ensure a consistent experience for users when they travel to different regions.
        * Obtain user activity metrics to better understand how to monetize their product.
        * Ensure compliance with regulations in the new regions (for example, GDPR).
        * Reduce infrastructure management time and cost.
        * Adopt the Google-recommended practices for cloud computing.

    Technical Requirements -
        * The application and backend must provide usage metrics and monitoring.
        * APIs require strong authentication and authorization.
        * Logging must be increased, and data should be stored in a cloud analytics platform.
        * Move to serverless architecture to facilitate elastic scaling.
        * Provide authorized access to internal apps in a secure manner.

# 80
    Which database should HipLocal use for storing user activity?
    ---
    BigQuery because "user activity metrics to better understand how to monetize their product" => for data analysis

# 81

    HipLocal is configuring their access controls.
    Which firewall configuration should they implement?
    ---
    Allow traffic on port 443 for a specific tag.
    HipLocal could create a tag called "trusted-services" and apply it to their web servers and load balancers. They could then configure their firewall to allow traffic on port 443 only from resources with the "trusted-services" tag. This would prevent unauthorized access to their network while allowing legitimate HTTPS traffic.

    In summary: Option C provides the most secure and flexible approach to configuring HipLocal's firewall, allowing them to control access to their network on port 443 while maintaining security.

# 82
    HipLocal's data science team wants to analyze user reviews.
    How should they prepare the data?
    ---
    Use the Cloud Data Loss Prevention API for de-identification of the review dataset
    Data Loss Prevention API: This API is specifically designed for identifying and protecting sensitive data. It can be used to de-identify data, which means replacing sensitive information with non-sensitive substitutes. This is ideal for user reviews, as they might contain personal information like names, addresses, or other details that need to be protected.
    De-identification: De-identification is the process of removing or replacing sensitive information in a dataset while preserving its usefulness for analysis. This is crucial for HipLocal's data science team, as they need to analyze user reviews without compromising user privacy.

# 83

# 84

    You have an application deployed in production. When a new version is deployed, you want to ensure that all production traffic is routed to the new version of your application. You also want to keep the previous version deployed so that you can revert to it if there is an issue with the new version.
    Which deployment strategy should you use?
    ---
    Blue/green deployment
    Blue/Green Deployment: This strategy involves running two identical environments: a "blue" environment (the current production version) and a "green" environment (the new version). Traffic is initially routed to the blue environment. Once the green environment is fully deployed and tested, traffic is switched over to the green environment. This allows for a quick rollback to the blue environment if there are issues with the new version.

# 85

    You are porting an existing Apache/MySQL/PHP application stack from a single machine to Google
    Kubernetes Engine. You need to determine how to containerize the application. Your approach should follow Google-recommended best practices for availability.
    What should you do?
    ---
    Package each component in a separate container. Implement readiness and liveness probes.

# 86

    You are developing an application that will be launched on Compute Engine instances into multiple distinct projects, each corresponding to the environments in your software development process (development, QA, staging, and production). The instances in each project have the same application code but a different configuration. During deployment, each instance should receive the application's configuration based on the environment it serves. You want to minimize the number of steps to configure this flow. What should you do?
    ---
    The best answer is D. During each instance launch, configure an instance custom-metadata key named ‘environment’ whose value is the environment the instance serves. Use your deployment tool to query the instance metadata, and configure the application based on the ‘environment’ value.

# 87

    You are developing an ecommerce application that stores customer, order, and inventory data as relational tables inside Cloud Spanner. During a recent load test, you discover that Spanner performance is not scaling linearly as expected. Which of the following is the cause?
    ---
    The use of Version 1 UUIDs as primary keys that increase monotonically.
    correct https://cloud.google.com/spanner/docs/schema-and-data-model#choosing_a_primary_key
    Version 1 UUIDs and Monotonicity: Version 1 UUIDs are generated based on timestamps and MAC addresses. When used as primary keys, they can lead to performance issues in Cloud Spanner due to their non-monotonic nature. As new UUIDs are generated, they are not guaranteed to be in a sequential order, which can cause fragmentation in the underlying storage and lead to slower queries.

# 88

    You are developing an application that reads credit card data from a Pub/Sub subscription. You have written code and completed unit testing. You need to test the
    Pub/Sub integration before deploying to Google Cloud. What should you do?
    ---
    Create a service to publish messages, and deploy the Pub/Sub emulator. Publish a standard set of testing messages from the publishing service to the emulator.

# 89

    You are designing an application that will subscribe to and receive messages from a single Pub/Sub topic and insert corresponding rows into a database. Your application runs on Linux and leverages preemptible virtual machines to reduce costs. You need to create a shutdown script that will initiate a graceful shutdown.
    What should you do?
    ---
    Write a shutdown script that uses inter-process signals to notify the application process to disconnect from the database.

# 90

    You work for a web development team at a small startup. Your team is developing a Node.js application using Google Cloud services, including Cloud Storage and Cloud Build. The team uses a Git repository for version control. Your manager calls you over the weekend and instructs you to make an emergency update to one of the company's websites, and you're the only developer available. You need to access Google Cloud to make the update, but you don't have your work laptop. You are not allowed to store source code locally on a non-corporate computer. How should you set up your developer environment?
    ---
    Use Cloud Shell and the built-in code editor for development. Send your source code updates as pull requests.

# 91

    Your team develops services that run on Google Kubernetes Engine. You need to standardize their log data using Google-recommended practices and make the data more useful in the fewest number of steps. What should you do? (Choose two.)
    ---
    Write log output to standard output (stdout) as single-line JSON to be ingested into Cloud Logging as structured logs.
    Create aggregated exports on application logs to BigQuery to facilitate log analytics.

# 92

    You are designing a deployment technique for your new applications on Google Cloud. As part of your deployment planning, you want to use live traffic to gather performance metrics for both new and existing applications. You need to test against the full production load prior to launch. What should you do?
    ---
    Use canary deployment

# 93

    You support an application that uses the Cloud Storage API. You review the logs and discover multiple HTTP 503 Service Unavailable error responses from the
    API. Your application logs the error and does not take any further action. You want to implement Google-recommended retry logic to improve success rates.
    Which approach should you take?
    ---
    Retry each failure at increasing time intervals up to a maximum number of tries

# 94

    You need to redesign the ingestion of audit events from your authentication service to allow it to handle a large increase in traffic. Currently, the audit service and the authentication system run in the same Compute Engine virtual machine. You plan to use the following Google Cloud tools in the new architecture:
    ✑ Multiple Compute Engine machines, each running an instance of the authentication service
    ✑ Multiple Compute Engine machines, each running an instance of the audit service
    ✑ Pub/Sub to send the events from the authentication services.
    How should you set up the topics and subscriptions to ensure that the system can handle a large volume of messages and can scale efficiently?
    ---
    A is correct. This is the most flexible way to scale, allowing the authentication and audit services to be sized independently according to load.
    B is incorrect. This will cause messages to be duplicated, one copy per subscription.
    C is incorrect. This will allow the system to scale, but push subscriptions are less suited to handle large volumes of messages.
    D is incorrect. This will allow the system to scale, however each audit service will listen to all subscriptions.
    E. is incorrect. This will allow the system to scale, however it will require each audit service to listen to all subscriptions. Also push subscriptions are less suited to

# 95

    You are developing a marquee stateless web application that will run on Google Cloud. The rate of the incoming user traffic is expected to be unpredictable, with no traffic on some days and large spikes on other days. You need the application to automatically scale up and down, and you need to minimize the cost associated with running the application. What should you do?
    ---
    Python + Firestore + Cloud Run:
    Truly serverless (scales to zero)
    Pay only for actual usage
    Handles spikes well
    Firestore scales automatically
    Minimal management overhead
    Lowest cost when no traffic

# 96

    You have written a Cloud Function that accesses other Google Cloud resources. You want to secure the environment using the principle of least privilege. What should you do?
    ---
     Create a new service account that has a custom IAM role to access the resources. The deployer is given permission to act as the new service account.

# 97

    You are a SaaS provider deploying dedicated blogging software to customers in your Google Kubernetes Engine (GKE) cluster. You want to configure a secure multi-tenant platform to ensure that each customer has access to only their own blog and can't affect the workloads of other customers. What should you do?
    ---
    Deploy a namespace per tenant and use Network Policies in each blog deployment

# 98

    You have decided to migrate your Compute Engine application to Google Kubernetes Engine. You need to build a container image and push it to Artifact Registry using Cloud Build. What should you do? (Choose two.)
    ---
    To build a container image and push it to Artifact Registry using Cloud Build, you should:

    Run gcloud builds submit in the directory that contains the application source code. This command will trigger Cloud Build to build the container image and push it to Artifact Registry.

    In the application source directory, create a file named cloudbuild.yaml that contains the instructions for building and pushing the container image. The file should contain the following steps:

    steps:
    -name: 'grc.io/cloud-builders/docker'
    args: ['build','-t','grc.io/$PROJECT_ID','app-name','.']
    -name: 'grc.io/cloud-builders/docker'
    args: ['push','grc.io/$PROJECT_ID/app-name']

    This file will be used by Cloud Build to build and push the container image.
# 99

    You are developing an internal application that will allow employees to organize community events within your company. You deployed your application on a single Compute Engine instance. Your company uses Google Workspace (formerly G Suite), and you need to ensure that the company employees can authenticate to the application from anywhere. What should you do?
    ---
    Add an HTTP(S) load balancer in front of the instance, and set up Identity-Aware Proxy (IAP). Configure the IAP settings to allow your company domain to access the website.

# 100

    Your development team is using Cloud Build to promote a Node.js application built on App Engine from your staging environment to production. The application relies on several directories of photos stored in a Cloud Storage bucket named webphotos-staging in the staging environment. After the promotion, these photos must be available in a Cloud Storage bucket named webphotos-prod in the production environment. You want to automate the process where possible. What should you do?
    ---
    C.Add a build step in the cloudbuild.yaml file before the promotion step with the arguments:
    -name: gcr.io/cloud-builders/gsutil
    args: ['cp','-r','gs://webphotos-staging','gs://webphotos-prod']
    waitFor: ['-']
    
    You should add a build step in the cloudbuild.yaml file before the promotion step with the arguments shown above. This build step will use the gsutil tool to copy the photos from the webphotos-staging bucket to the webphotos-prod bucket. The -r flag tells gsutil to copy all files in the bucket recursively, and the waitFor parameter tells Cloud Build to wait for this step to complete before continuing with the promotion step.

# 101

    You are developing a web application that will be accessible over both HTTP and HTTPS and will run on Compute Engine instances. On occasion, you will need to SSH from your remote laptop into one of the Compute Engine instances to conduct maintenance on the app. How should you configure the instances while following Google-recommended best practices?
    ---
    Set up a backend with Compute Engine web server instances with a private IP address behind an HTTP(S) load balancer. Set up a bastion host with a public IP address and open firewall ports. Connect to the web instances using the bastion host

# 102

    You have a mixture of packaged and internally developed applications hosted on a Compute Engine instance , that is running Linux. These applications write log records as text in local files. You want the logs to be written to Cloud Logging. What should you do?
    ---
    Fluentd is for cloud logging agent
    Collectd is used for Monitoring agents

# 103

    You want to create `fully baked` or `golden` Compute Engine images for your application. You need to bootstrap your application to connect to the appropriate database according to the environment the application is running on (test, staging, production). What should you do?
    ---
    When creating the Compute Engine instance, create a metadata item with a key of ג€DATABASEג€ and a value for the appropriate database connection string. In your application, query the metadata server for the ג€DATABASEג€ value, and use the value to connect to the appropriate database.
    (https://cloud.google.com/compute/docs/metadata/querying-metadata)

# 104

    You are developing a microservice-based application that will be deployed on a Google Kubernetes Engine cluster. The application needs to read and write to a
    Spanner database. You want to follow security best practices while minimizing code changes. How should you configure your application to retrieve Spanner credentials?
    ---
    Configure the appropriate service accounts, and use Workload Identity to run the pods.

# 105

    You are deploying your application on a Compute Engine instance that communicates with Cloud SQL. You will use Cloud SQL Proxy to allow your application to communicate to the database using the service account associated with the application's instance. You want to follow the Google-recommended best practice of providing minimum access for the role assigned to the service account. What should you do?
    ---
    Assign the Cloud SQL Client role.

# 106

    Your team develops stateless services that run on Google Kubernetes Engine (GKE). You need to deploy a new service that will only be accessed by other services running in the GKE cluster. The service will need to scale as quickly as possible to respond to changing load. What should you do?
    ---
    Use a Horizontal Pod Autoscaler to scale the containers, and expose them via a ClusterIP Service.

# 107

    You recently migrated a monolithic application to Google Cloud by breaking it down into microservices. One of the microservices is deployed using Cloud
    Functions. As you modernize the application, you make a change to the API of the service that is backward-incompatible. You need to support both existing callers who use the original API and new callers who use the new API. What should you do?
    ---
    Leave the original Cloud Function as-is and deploy a second Cloud Function with the new API. Use Cloud Endpoints to provide an API gateway that exposes a versioned API

# 108

    You are developing an application that will allow users to read and post comments on news articles. You want to configure your application to store and display user-submitted comments using Firestore. How should you design the schema to support an unknown number of comments and articles?
    ---
    Store each comment in a subcollection of the article.