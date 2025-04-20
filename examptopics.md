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

# 109

# 110

# 111

    You manage an application that runs in a Compute Engine instance. You also have multiple backend services executing in stand-alone Docker containers running in Compute Engine instances. The Compute Engine instances supporting the backend services are scaled by managed instance groups in multiple regions. You want your calling application to be loosely coupled. You need to be able to invoke distinct service implementations that are chosen based on the value of an HTTP header found in the request. Which Google Cloud feature should you use to invoke the backend services?
    ---
    Internal HTTP(S) Load Balancing

# 112

    Your team is developing an ecommerce platform for your company. Users will log in to the website and add items to their shopping cart. Users will be automatically logged out after 30 minutes of inactivity. When users log back in, their shopping cart should be saved. How should you store users' session and shopping cart information while following Google-recommended best practices?

# 113

# 114

    You are developing a new application that has the following design requirements:
    Creation and changes to the application infrastructure are versioned and auditable.
    The application and deployment infrastructure uses Google-managed services as much as possible.
    The application runs on a serverless compute platform.
    How should you design the application's architecture?
    ---
    Store the application and infrastructure source code in a Git repository. 2. Use Cloud Build to deploy the application infrastructure with Terraform. 3. Deploy the application to a Cloud Function as a pipeline step.

# 115

    You are creating and running containers across different projects in Google Cloud. The application you are developing needs to access Google Cloud services from within Google Kubernetes Engine (GKE). What should you do?
    ---
    Use a Google service account to run the Pod with Workload Identity.

# 116

    You have containerized a legacy application that stores its configuration on an NFS share. You need to deploy this application to Google Kubernetes Engine
    (GKE) and do not want the application serving traffic until after the configuration has been retrieved. What should you do?
    ---
    Create a PersistentVolumeClaim on the GKE cluster. Access the configuration files from the volume, and start the service using an ENTRYPOINT script.

# 117

    Your team is developing a new application using a PostgreSQL database and Cloud Run. You are responsible for ensuring that all traffic is kept private on Google
    Cloud. You want to use managed services and follow Google-recommended best practices. What should you do?
    ---
    Enable Cloud SQL and Cloud Run in the same project. 2. Configure a private IP address for Cloud SQL. Enable private services access. 3. Create a Serverless VPC Access connector. 4. Configure Cloud Run to use the connector to connect to Cloud SQL.

# 118

    You are developing an application that will allow clients to download a file from your website for a specific period of time. How should you design the application to complete this task while following Google-recommended best practices?
    ---
    Generate and assign a Cloud Storage-signed URL for the file. Make the URL available for the client to download.

# 119

    Your development team has been asked to refactor an existing monolithic application into a set of composable microservices. Which design aspects should you implement for the new application? (Choose two.)
    ---
    Create an API contract agreement between the microservice implementation and microservice caller.
    Ensure that sufficient instances of the microservice are running to accommodate the performance requirements.

# 120

    You deployed a new application to Google Kubernetes Engine and are experiencing some performance degradation. Your logs are being written to Cloud
    Logging, and you are using a Prometheus sidecar model for capturing metrics. You need to correlate the metrics and data from the logs to troubleshoot the performance issue and send real-time alerts while minimizing costs. What should you do?
    ---
    Export the Prometheus metrics and use Cloud Monitoring to view them as external metrics. Configure Cloud Monitoring to create log-based metrics from the logs, and correlate them with the Prometheus data.

# 121

    You have been tasked with planning the migration of your company's application from on-premises to Google Cloud. Your company's monolithic application is an ecommerce website. The application will be migrated to microservices deployed on Google Cloud in stages. The majority of your company's revenue is generated through online sales, so it is important to minimize risk during the migration. You need to prioritize features and select the first functionality to migrate. What should you do?
    ---
    Migrate the Product catalog, which has integrations to the frontend and product database. (Least dependentcies)

# 122

    Your team develops services that run on Google Kubernetes Engine. Your team's code is stored in Cloud Source Repositories. You need to quickly identify bugs in the code before it is deployed to production. You want to invest in automation to improve developer feedback and make the process as efficient as possible.
    What should you do?
    ---
    Use Cloud Build to automate building container images from code based on Git tags.

# 123
    
    Your team is developing an application in Google Cloud that executes with user identities maintained by Cloud Identity. Each of your application's users will have an associated Pub/Sub topic to which messages are published, and a Pub/Sub subscription where the same user will retrieve published messages. You need to ensure that only authorized users can publish and subscribe to their own specific Pub/Sub topic and subscription. What should you do?
    ---
    Granting IAM at resource level is enough.
    If project level permission is given then user will be having publisher and subscriber roles for all the pub-sub topics created within the project. So this should be avoided according to the question asked

# 124

# 125

    You are developing an ecommerce web application that uses App Engine standard environment and Memorystore for Redis. When a user logs into the app, the application caches the user's information (e.g., session, name, address, preferences), which is stored for quick retrieval during checkout.
    While testing your application in a browser, you get a 502 Bad Gateway error. You have determined that the application is not connecting to Memorystore. What is the reason for this error?
    ---
    You configured your Serverless VPC Access connector in a different region than your App Engine instance.

# 126

    Your team develops services that run on Google Cloud. You need to build a data processing service and will use Cloud Functions. The data to be processed by the function is sensitive. You need to ensure that invocations can only happen from authorized services and follow Google-recommended best practices for securing functions. What should you do?
    ---
    Create a service account with the Cloud Functions Invoker role. Use that service account to invoke the function.

# 127

    You are deploying your applications on Compute Engine. One of your Compute Engine instances failed to launch. What should you do? (Choose two.)
    ---
    Determine whether your file system is corrupted.
    Check whether your instance boot disk is completely full.

# 128

    Your web application is deployed to the corporate intranet. You need to migrate the web application to Google Cloud. The web application must be available only to company employees and accessible to employees as they travel. You need to ensure the security and accessibility of the web application while minimizing application changes. What should you do?
    ---
    Configure Identity-Aware Proxy to allow employees to access the application through its public IP address.

# 129

    You have an application that uses an HTTP Cloud Function to process user activity from both desktop browser and mobile application clients. This function will serve as the endpoint for all metric submissions using HTTP POST.
    Due to legacy restrictions, the function must be mapped to a domain that is separate from the domain requested by users on web or mobile sessions. The domain for the Cloud Function is https://fn.example.com. Desktop and mobile clients use the domain https://www.example.com. You need to add a header to the function's
    HTTP response so that only those browser and mobile sessions can submit metrics to the Cloud Function. Which response header should you add?
    ---
    Access-Control-Allow-origin: https://www.example.com # NOT https://*.example.com, www is determine for '*'

# 130

    You have an HTTP Cloud Function that is called via POST. Each submission's request body has a flat, unnested JSON structure containing numeric and text data. After the Cloud Function completes, the collected data should be immediately available for ongoing and complex analytics by many users in parallel. How should you persist the submissions?
    ---
    ransform the POST request's JSON data, and stream it into BigQuery.

# 131

    Your security team is auditing all deployed applications running in Google Kubernetes Engine. After completing the audit, your team discovers that some of the applications send traffic within the cluster in clear text. You need to ensure that all application traffic is encrypted as quickly as possible while minimizing changes to your applications and maintaining support from Google. What should you do?
    ---
    Install Istio, enable proxy injection on your application namespace, and then enable mTLS.

# 132

    You migrated some of your applications to Google Cloud. You are using a legacy monitoring platform deployed on-premises for both on-premises and cloud- deployed applications. You discover that your notification system is responding slowly to time-critical problems in the cloud applications. What should you do?
    ---
    Use Cloud Logging and Cloud Monitoring to capture logs, monitor, and send alerts. Send them to your existing platform.
 
# 133

    You recently deployed your application in Google Kubernetes Engine, and now need to release a new version of your application. You need the ability to instantly roll back to the previous version in case there are issues with the new version. Which deployment model should you use?
    ---
    Perform a blue/green deployment, and test your new application after the deployment is. complete.
# 134

# 135

    You manage an ecommerce application that processes purchases from customers who can subsequently cancel or change those purchases. You discover that order volumes are highly variable and the backend order-processing system can only process one request at a time. You want to ensure seamless performance for customers regardless of usage volume. It is crucial that customers' order update requests are performed in the sequence in which they were generated. What should you do?
    ---
    Use a Pub/Sub subscriber in pull mode and use a data store to manage ordering.

# 136

    Your company needs a database solution that stores customer purchase history and meets the following requirements:
    ✑ Customers can query their purchase immediately after submission. 
    ✑ Purchases can be sorted on a variety of fields.
    ✑ Distinct record formats can be stored at the same time.
    Which storage option satisfies these requirements?
    ---
    Firestore in Native mode
    
    Firestore is for storing semi structured data. It is optimized for high reads and low writes. Since each document can store different collection types, ( MONGO DB ), fire store is suitable for the above requirements.

# 137

    You recently developed a new service on Cloud Run. The new service authenticates using a custom service and then writes transactional information to a Cloud
    Spanner database. You need to verify that your application can support up to 5,000 read and 1,000 write transactions per second while identifying any bottlenecks that occur. Your test infrastructure must be able to autoscale. What should you do?
    ---
    Create a Google Kubernetes Engine cluster running the Locust or JMeter images to dynamically generate load tests. Analyze the results using Cloud Trace.

# 138

    You are using Cloud Build for your CI/CD pipeline to complete several tasks, including copying certain files to Compute Engine virtual machines. Your pipeline requires a flat file that is generated in one builder in the pipeline to be accessible by subsequent builders in the same pipeline. How should you store the file so that all the builders in the pipeline can access it?
    ---
    Output the file contents to a file in /workspace. Read from the same /workspace file in the subsequent build step.

# 139

    Your company’s development teams want to use various open source operating systems in their Docker builds. When images are created in published containers in your company’s environment, you need to scan them for Common Vulnerabilities and Exposures (CVEs). The scanning process must not impact software development agility. You want to use managed services where possible. What should you do?
    ---
     Enable the Vulnerability scanning setting in the Container Registry

# 140

# 141
    You are building a CI/CD pipeline that consists of a version control system, Cloud Build, and Container Registry. Each time a new tag is pushed to the repository, a Cloud Build job is triggered, which runs unit tests on the new code builds a new Docker container image, and pushes it into Container Registry. The last step of your pipeline should deploy the new container to your production Google Kubernetes Engine (GKE) cluster. You need to select a tool and deployment strategy that meets the following requirements:
    • Zero downtime is incurred
    • Testing is fully automated
    • Allows for testing before being rolled out to users
    • Can quickly rollback if needed

    What should you do?
    ---
    Trigger another Cloud Build job that uses the Kubernetes CLI tools to deploy your new container to your GKE cluster, where you can perform a shadow test.

# 142

# 143

    You need to deploy a new European version of a website hosted on Google Kubernetes Engine. The current and new websites must be accessed via the same HTTP(S) load balancer's external IP address, but have different domain names. What should you do?
    ---
    Modify the existing Ingress resource with a host rule matching the new domain

# 144

    You are developing a single-player mobile game backend that has unpredictable traffic patterns as users interact with the game throughout the day and night. You want to optimize costs by ensuring that you have enough resources to handle requests, but minimize over-provisioning. You also want the system to handle traffic spikes efficiently. Which compute platform should you use?
    ---
    Cloud Run

# 145

    The development teams in your company want to manage resources from their local environments. You have been asked to enable developer access to each team’s Google Cloud projects. You want to maximize efficiency while following Google-recommended best practices. What should you do?
    ---
    Create groups, add the users to their groups, assign the relevant roles to the groups, and then provide the users with each relevant Project ID

    Project ID more friendly than project number

# 146

    Your company’s product team has a new requirement based on customer demand to autoscale your stateless and distributed service running in a Google Kubernetes Engine (GKE) duster. You want to find a solution that minimizes changes because this feature will go live in two weeks. What should you do?
    ---
    Deploy a Horizontal Pod Autoscaler, and scale based on the CPU toad

# 147

    Your application is composed of a set of loosely coupled services orchestrated by code executed on Compute Engine. You want your application to easily bring up new Compute Engine instances that find and use a specific version of a service. How should this be configured?
    ---
     Define your service endpoint information as metadata that is retrieved at runtime and used to connect to the desired service.

# 148

    You are developing a microservice-based application that will run on Google Kubernetes Engine (GKE). Some of the services need to access different Google Cloud APIs. How should you set up authentication of these services in the cluster following Google-recommended best practices? (Choose two.)
    ---
     Enable Workload Identity in the cluster via the gcloud command-line tool.
    Use gcloud to bind the Kubernetes service account and the Google service account using roles/iam.workloadIdentity.

# 149

    Your development team has been tasked with maintaining a .NET legacy application. The application incurs occasional changes and was recently updated. Your goal is to ensure that the application provides consistent results while moving through the CI/CD pipeline from environment to environment. You want to minimize the cost of deployment while making sure that external factors and dependencies between hosting environments are not problematic. Containers are not yet approved in your organization. What should you do?
    ---
    Use Cloud Build to deploy the application as a new Compute Engine image for each build. Use this image in each environment

# 150

    The new version of your containerized application has been tested and is ready to deploy to production on Google Kubernetes Engine. You were not able to fully load-test the new version in pre-production environments, and you need to make sure that it does not have performance problems once deployed. Your deployment must be automated. What should you do?
    ---
    Deploy the application via a continuous delivery pipeline using canary deployments. Use Cloud Monitoring to look for performance issues. and ramp up traffic as the metrics support it.

# 151

# 152

    You are a developer working on an internal application for payroll processing. You are building a component of the application that allows an employee to submit a timesheet, which then initiates several steps:

    • An email is sent to the employee and manager, notifying them that the timesheet was submitted.
    • A timesheet is sent to payroll processing for the vendor's API.
    • A timesheet is sent to the data warehouse for headcount planning.

    These steps are not dependent on each other and can be completed in any order. New steps are being considered and will be implemented by different development teams. Each development team will implement the error handling specific to their step. What should you do?
    ---
    Create a Pub/Sub topic for timesheet submissions. Create a subscription for each downstream development team to subscribe to the topic.

# 153

    You are designing an application that uses a microservices architecture. You are planning to deploy the application in the cloud and on-premises. You want to make sure the application can scale up on demand and also use managed services as much as possible. What should you do?
    ---
    B. Create a GKE cluster in each environment with Anthos, and use Cloud Run for Anthos to deploy your application to each cluster.

    Using Anthos to manage Kubernetes clusters in both cloud and on-premises environments allows for consistency in deployment and management across both environments. Deploying the application using Cloud Run for Anthos allows for easy scaling on demand and use of managed services such as Cloud SQL and Memorystore. Additionally, Cloud Run for Anthos can be deployed to both GKE clusters and on-premises Kubernetes clusters, allowing for a consistent deployment experience across environments.

# 154

    You want to migrate an on-premises container running in Knative to Google Cloud. You need to make sure that the migration doesn't affect your application's deployment strategy, and you want to use a fully managed service. Which Google Cloud service should you use to deploy your container?
    ---
    # knative: open source to build serverless app
    Cloud Run, not app engine because app-engine supported some specific lang

# 155

    Pub/Sub -> Data flow -> bigquery
    Data ingest: pub/sub
    pipeline: Dataflow
    transaction: filestore
    Analytics: BQ

# 156

    Your company just experienced a Google Kubernetes Engine (GKE) API outage due to a zone failure. You want to deploy a highly available GKE architecture that minimizes service interruption to users in the event of a future zone failure. What should you do?
    ---
    Regional Clusters
    - Span multiple zones within a region.
    - May not guarantee that all zones within the region are always available.
    - Can provide some level of redundancy, but might not be as resilient as multi-zone clusters in the event of a zone failure.

    Multi-Zone Clusters
    - Distribute your workload across multiple zones within a region.
    - Ensure that your application remains highly available and resilient to zone failures.
    - Automatically move your workloads to the remaining healthy zones if one zone fails.
    - Provide the highest level of availability among the GKE deployment options.

    In summary, while both regional and multi-zone clusters can provide redundancy, multi-zone clusters offer a higher level of availability by distributing your workload across multiple zones, ensuring minimal service disruption in the event of a zone failure.

# 157

    Your team develops services that run on Google Cloud. You want to process messages sent to a Pub/Sub topic, and then store them. Each message must be processed exactly once to avoid duplication of data and any data conflicts. You need to use the cheapest and most simple solution. What should you do?
    ---
    Process the messages with a Dataflow streaming pipeline using Apache Beam's PubSubIO package, and write the output to storage.
    PubSubIO default fileter by msg_id

# 158
    
    You are running a containerized application on Google Kubernetes Engine. Your container images are stored in Container Registry. Your team uses CI/CD practices. You need to prevent the deployment of containers with known critical vulnerabilities. What should you do?
    ---
    Enable the Container Scanning API to perform vulnerability scanning
    • Programmatically review vulnerability reporting through the Container Scanning API, and provide an attestation that the container is free of known critical vulnerabilities
    • Use Binary Authorization to implement a policy that forces the attestation to be provided before the container is deployed

# 159

    You have an on-premises application that authenticates to the Cloud Storage API using a user-managed service account with a user-managed key. The application connects to Cloud Storage using Private Google Access over a Dedicated Interconnect link. You discover that requests from the application to access objects in the Cloud Storage bucket are failing with a 403 Permission Denied error code. What is the likely cause of this issue?
    ---
    The service account key has been rotated but not updated on the application server.

# 160

    You are using the Cloud Client Library to upload an image in your application to Cloud Storage. Users of the application report that occasionally the upload does not complete and the client library reports an HTTP 504 Gateway Timeout error. You want to make the application more resilient to errors. What changes to the application should you make?
    ---
    Write an exponential backoff process around the client library call.

# 161

    You are building a mobile application that will store hierarchical data structures in a database. The application will enable users working offline to sync changes when they are back online. A backend service will enrich the data in the database using a service account. The application is expected to be very popular and needs to scale seamlessly and securely. Which database and IAM role should you use?
    ---
    Use Firestore in Native mode and assign the roles/datastore.user role to the service account. 

# 162

    Your application is deployed on hundreds of Compute Engine instances in a managed instance group (MIG) in multiple zones. You need to deploy a new instance template to fix a critical vulnerability immediately but must avoid impact to your service. What setting should be made to the MIG after updating the instance template?
    ---
    Set the Minimum Wait time to 0 seconds.
    https://cloud.google.com/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#minimum_wait_time

# 163

    You made a typo in a low-level Linux configuration file that prevents your Compute Engine instance from booting to a normal run level. You just created the Compute Engine instance today and have done no other maintenance on it, other than tweaking files. How should you correct this error?
    ---
    Configure and log in to the Compute Engine instance through the serial port, and change the file

# 164

# 165

    Your company’s corporate policy states that there must be a copyright comment at the very beginning of all source files. You want to write a custom step in Cloud Build that is triggered by each source commit. You need the trigger to validate that the source contains a copyright and add one for subsequent steps if not there. What should you do?
    ---
    Build a new Docker container that examines the files in /workspace and then checks and adds a copyright for each source file. Changed files are explicitly committed back to the source repository.

# 166

# 167 case study

## company

    Company Overview -
    HipLocal is a community application designed to facilitate communication between people in close proximity. It is used for event planning and organizing sporting events, and for businesses to connect with their local communities. HipLocal launched recently in a few neighborhoods in Dallas and is rapidly growing into a global phenomenon. Its unique style of hyper-local community communication and business outreach is in demand around the world.


    Executive Statement -
    We are the number one local community app; it's time to take our local community services global. Our venture capital investors want to see rapid growth and the same great experience for new local and virtual communities that come online, whether their members are 10 or 10000 miles away from each other.


    Solution Concept -
    HipLocal wants to expand their existing service, with updated functionality, in new regions to better serve their global customers. They want to hire and train a new team to support these regions in their time zones. They will need to ensure that the application scales smoothly and provides clear uptime data, and that they analyze and respond to any issues that occur.


    Existing Technical Environment -
    HipLocal's environment is a mix of on-premises hardware and infrastructure running in Google Cloud Platform. The HipLocal team understands their application well, but has limited experience in global scale applications. Their existing technical environment is as follows:
    • Existing APIs run on Compute Engine virtual machine instances hosted in GCP.
    • State is stored in a single instance MySQL database in GCP.
    • Release cycles include development freezes to allow for QA testing.
    • The application has no logging.
    • Applications are manually deployed by infrastructure engineers during periods of slow traffic on weekday evenings.
    • There are basic indicators of uptime; alerts are frequently fired when the APIs are unresponsive.


    Business Requirements -
    HipLocal's investors want to expand their footprint and support the increase in demand they are seeing. Their requirements are:
    • Expand availability of the application to new regions.
    • Support 10x as many concurrent users.
    • Ensure a consistent experience for users when they travel to different regions.
    • Obtain user activity metrics to better understand how to monetize their product.
    • Ensure compliance with regulations in the new regions (for example, GDPR).
    • Reduce infrastructure management time and cost.
    • Adopt the Google-recommended practices for cloud computing.
    ○ Develop standardized workflows and processes around application lifecycle management.
    ○ Define service level indicators (SLIs) and service level objectives (SLOs).


    Technical Requirements -
    • Provide secure communications between the on-premises data center and cloud-hosted applications and infrastructure.
    • The application must provide usage metrics and monitoring.
    • APIs require authentication and authorization.
    • Implement faster and more accurate validation of new features.
    • Logging and performance metrics must provide actionable information to be able to provide debugging information and alerts.
    • Must scale to meet user demand.


    For this question, refer to the HipLocal case study.

# 167

    How should HipLocal redesign their architecture to ensure that the application scales to support a large increase in users?
    ---
    Use Memorystore to store session information and CloudSQL to store state information. Use a Google Cloud-managed load balancer to distribute the load between instances. Use managed instance groups for scaling.

# 168

# 169
    HipLocal's application uses Cloud Client Libraries to interact with Google Cloud. HipLocal needs to configure authentication and authorization in the Cloud Client Libraries to implement least privileged access for the application. What should they do?
    ---

    C: Create a service account for the application. Export and deploy the private key for the application. Use the service account to interact with Google Cloud.
    D: Create a service account for the application and for each Google Cloud API used by the application. Export and deploy the private keys used by the application. Use the service account with one Google Cloud API to interact with Google Cloud. => very complex

# 170
    
    You are in the final stage of migrating an on-premises data center to Google Cloud. You are quickly approaching your deadline, and discover that a web API is running on a server slated for decommissioning. You need to recommend a solution to modernize this API while migrating to Google Cloud. The modernized web API must meet the following requirements:

    • Autoscales during high traffic periods at the end of each month
    • Written in Python 3.x
    • Developers must be able to rapidly deploy new versions in response to frequent code changes

    You want to minimize cost, effort, and operational overhead of this migration. What should you do?
    ---
    Modernize and deploy the code on App Engine standard environment.

# 171

    You are in the final stage of migrating an on-premises data center to Google Cloud. You are quickly approaching your deadline, and discover that a web API is running on a server slated for decommissioning. You need to recommend a solution to modernize this API while migrating to Google Cloud. The modernized web API must meet the following requirements:

    • Autoscales during high traffic periods at the end of each month
    • Written in Python 3.x
    • Developers must be able to rapidly deploy new versions in response to frequent code changes

    You want to minimize cost, effort, and operational overhead of this migration. What should you do?
    ---
    kube secret

# 172

    You manage your company's ecommerce platform's payment system, which runs on Google Cloud. Your company must retain user logs for 1 year for internal auditing purposes and for 3 years to meet compliance requirements. You need to store new user logs on Google Cloud to minimize on-premises storage usage and ensure that they are easily searchable. You want to minimize effort while ensuring that the logs are stored correctly. What should you do?
    ---
    Store the logs in Cloud Logging as custom logs with a custom retention period.

# 173

    Your company has a new security initiative that requires all data stored in Google Cloud to be encrypted by customer-managed encryption keys. You plan to use Cloud Key Management Service (KMS) to configure access to the keys. You need to follow the "separation of duties" principle and Google-recommended best practices. What should you do? (Choose two.)
    ---
    Provision Cloud KMS in its own project. 
    Do not assign an owner to the Cloud KMS project.
    https://cloud.google.com/kms/docs/separation-of-duties#using_separate_project
    Instead, to allow for a separation of duties, you could run Cloud KMS in its own project, for example your-key-project. Then, depending on the strictness of your separation requirements, you could either:
    - (recommended) Create your-key-project without an owner at the project level, and designate an Organization Admin granted at the organization-level. Unlike an owner, an Organization Admin can't manage or use keys directly. They are restricted to setting IAM policies, which restrict who can manage and use keys. Using an organization-level node, you can further restrict permissions for projects in your organization.

# 174

# 175

    Your organization has recently begun an initiative to replatform their legacy applications onto Google Kubernetes Engine. You need to decompose a monolithic application into microservices. Multiple instances have read and write access to a configuration file, which is stored on a shared file system. You want to minimize the effort required to manage this transition, and you want to avoid rewriting the application code. What should you do?
    ---
    A. Create a new Cloud Storage bucket, and mount it via FUSE in the container.
    B. Create a new persistent disk, and mount the volume as a shared PersistentVolume.
    C. Create a new Filestore instance, and mount the volume as an NFS PersistentVolume.
    D. Create a new ConfigMap and volumeMount to store the contents of the configuration file.    
    ---
    Create a new Filestore instance, and mount the volume as an NFS PersistentVolume.
    A is incorrect, because Cloud Storage FUSE does not support concurrency and file locking.
    B is incorrect, because a persistent disk PersistentVolume is not read-write-many. It can only be read-write once or read-many.
    C is correct, because it’s the only managed, supported read-write-many storage option available for file-system access in Google Kubernetes Engine.
    D is incorrect, because the ConfigMap cannot be written to from the Pods.

    https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes
    https://cloud.google.com/filestore/docs/accessing-fileshares
    https://cloud.google.com/storage/docs/gcs-fuse

# 176

# 177

    You manage a microservices application on Google Kubernetes Engine (GKE) using Istio. You secure the communication channels between your microservices by implementing an Istio AuthorizationPolicy, a Kubernetes NetworkPolicy, and mTLS on your GKE cluster. You discover that HTTP requests between two Pods to specific URLs fail, while other requests to other URLs succeed. What is the cause of the connection issue?
    ---
     The Authorization Policy of your cluster is blocking HTTP requests for specific paths within your application.

# 178

    You recently migrated an on-premises monolithic application to a microservices application on Google Kubernetes Engine (GKE). The application has dependencies on backend services on-premises, including a CRM system and a MySQL database that contains personally identifiable information (PII). The backend services must remain on-premises to meet regulatory requirements.

    You established a Cloud VPN connection between your on-premises data center and Google Cloud. You notice that some requests from your microservices application on GKE to the backend services are failing due to latency issues caused by fluctuating bandwidth, which is causing the application to crash. How should you address the latency issues?
    ---
    Increase the number of Cloud VPN tunnels for the connection between Google Cloud and the on-premises services

# 179

    Cloud debugger is decrecated

# 180

    You are designing an application that consists of several microservices. Each microservice has its own RESTful API and will be deployed as a separate Kubernetes Service. You want to ensure that the consumers of these APIs aren't impacted when there is a change to your API, and also ensure that third-party systems aren't interrupted when new versions of the API are released. How should you configure the connection to the application following Google-recommended best practices?
    ---
    Use an Ingress that uses the API's URL to route requests to the appropriate backend.
    Selected Answer: A
    https://cloud.google.com/kubernetes-engine/docs/concepts/ingress#deprecated_annotation
    https://cloud.google.com/kubernetes-engine/docs/concepts/ingress#features_of_https_load_balancing

# 181
    
    Your team is building an application for a financial institution. The application's frontend runs on Compute Engine, and the data resides in Cloud SQL and one Cloud Storage bucket. The application will collect data containing PII, which will be stored in the Cloud SQL database and the Cloud Storage bucket. You need to secure the PII data. What should you do?
    ---
    C. 1. Configure a private IP address for Cloud SQL
    2. Use VPC-SC to create a service perimeter
    3. Add the Cloud SQL database and the Cloud Storage bucket to the same service perimeter

# 182

    You are designing a chat room application that will host multiple rooms and retain the message history for each room. You have selected Firestore as your database. How should you represent the data in Firestore?
    ---
    Create a collection for the rooms. For each room, create a document that contains a collection for documents, each of which contains a message.

# 183

    You are developing an application that will handle requests from end users. You need to secure a Cloud Function called by the application to allow authorized end users to authenticate to the function via the application while restricting access to unauthorized users. You will integrate Google Sign-In as part of the solution and want to follow Google-recommended best practices. What should you do?
    ---
    Deploy from a source code repository and grant users the roles/cloudfunctions.invoker role

# 184

# 185

    You are building a highly available and globally accessible application that will serve static content to users. You need to configure the storage and serving components. You want to minimize management overhead and latency while maximizing reliability for users. What should you do?
    ---
    D. 
    1. Create a Standard storage class, multi-regional Cloud Storage bucket. Put the static content in the bucket.
    2. Reserve an external IP address, and create an external HTTP(S) load balancer.
    3. Enable Cloud CDN, and send traffic to your backend bucket.

# 186 Case study

        Company Overview -
    HipLocal is a community application designed to facilitate communication between people in close proximity. It is used for event planning and organizing sporting events, and for businesses to connect with their local communities. HipLocal launched recently in a few neighborhoods in Dallas and is rapidly growing into a global phenomenon. Its unique style of hyper-local community communication and business outreach is in demand around the world.


    Executive Statement -
    We are the number one local community app; it's time to take our local community services global. Our venture capital investors want to see rapid growth and the same great experience for new local and virtual communities that come online, whether their members are 10 or 10000 miles away from each other.


    Solution Concept -
    HipLocal wants to expand their existing service, with updated functionality, in new regions to better serve their global customers. They want to hire and train a new team to support these regions in their time zones. They will need to ensure that the application scales smoothly and provides clear uptime data, and that they analyze and respond to any issues that occur.


    Existing Technical Environment -
    HipLocal's environment is a mix of on-premises hardware and infrastructure running in Google Cloud Platform. The HipLocal team understands their application well, but has limited experience in global scale applications. Their existing technical environment is as follows:
    • Existing APIs run on Compute Engine virtual machine instances hosted in GCP.
    • State is stored in a single instance MySQL database in GCP.
    • Release cycles include development freezes to allow for QA testing.
    • The application has no logging.
    • Applications are manually deployed by infrastructure engineers during periods of slow traffic on weekday evenings.
    • There are basic indicators of uptime; alerts are frequently fired when the APIs are unresponsive.


    Business Requirements -
    HipLocal's investors want to expand their footprint and support the increase in demand they are seeing. Their requirements are:
    • Expand availability of the application to new regions.
    • Support 10x as many concurrent users.
    • Ensure a consistent experience for users when they travel to different regions.
    • Obtain user activity metrics to better understand how to monetize their product.
    • Ensure compliance with regulations in the new regions (for example, GDPR).
    • Reduce infrastructure management time and cost.
    • Adopt the Google-recommended practices for cloud computing.
    ○ Develop standardized workflows and processes around application lifecycle management.
    ○ Define service level indicators (SLIs) and service level objectives (SLOs).


    Technical Requirements -
    • Provide secure communications between the on-premises data center and cloud-hosted applications and infrastructure.
    • The application must provide usage metrics and monitoring.
    • APIs require authentication and authorization.
    • Implement faster and more accurate validation of new features.
    • Logging and performance metrics must provide actionable information to be able to provide debugging information and alerts.
    • Must scale to meet user demand.


    For this question refer to the HipLocal case study.

# 186

    HipLocal wants to reduce the latency of their services for users in global locations. They have created read replicas of their database in locations where their users reside and configured their service to read traffic using those replicas. How should they further reduce latency for all database interactions with the least amount of effort?
    ---
    Migrate the database to Cloud Spanner and use it to serve all global user traffic

# 187

    Define service level indicators (SLIs) and service level objectives (SLOs).
    Which Google Cloud product addresses HipLocal’s business requirements for service level indicators and objectives?
    ---
    Cloud Monitoring

# 188

    A recent security audit discovers that HipLocal’s database credentials for their Compute Engine-hosted MySQL databases are stored in plain text on persistent disks. HipLocal needs to reduce the risk of these credentials being stolen. What should they do?
    ---
    Grant the roles/secretmanager.secretAccessor role to the Compute Engine service account. Store and access the database credentials with the Secret Manager API

# 189

    HipLocal is expanding into new locations. They must capture additional data each time the application is launched in a new European country. This is causing delays in the development process due to constant schema changes and a lack of environments for conducting testing on the application changes. How should they resolve the issue while meeting the business requirements?
    ---
    Migrate data to Firestore in Native mode and set up instances in Europe and North America. Instruct the development teams to use the Cloud SDK to emulate a local Firestore in Native mode development environment. 

# 190

    You are writing from a Go application to a Cloud Spanner database. You want to optimize your application’s performance using Google-recommended best practices. What should you do?
    ---
    Write to Cloud Spanner using Cloud Client Libraries

# 191

    You have an application deployed in Google Kubernetes Engine (GKE). You need to update the application to make authorized requests to Google Cloud managed services. You want this to be a one-time setup, and you need to follow security best practices of auto-rotating your security keys and storing them in an encrypted store. You already created a service account with appropriate access to the Google Cloud service. What should you do next?
    ---
    Assign the Google Cloud service account to your GKE Pod using Workload Identity.

# 192

    You are planning to deploy hundreds of microservices in your Google Kubernetes Engine (GKE) cluster. How should you secure communication between the microservices on GKE using a managed service?
    ---
    Install Anthos Service Mesh, and enable mTLS in your Service Mesh

# 193

    You are developing an application that will store and access sensitive unstructured data objects in a Cloud Storage bucket. To comply with regulatory requirements, you need to ensure that all data objects are available for at least 7 years after their initial creation. Objects created more than 3 years ago are accessed very infrequently (less than once a year). You need to configure object storage while ensuring that storage cost is optimized. What should you do? (Choose two.)
    ---
    Set a retention policy on the bucket with a period of 7 years.
    Create an object lifecycle policy on the bucket that moves objects from Standard Storage to Archive Storage after 3 years.

# 194

    You are developing an application using different microservices that must remain internal to the cluster. You want the ability to configure each microservice with a specific number of replicas. You also want the ability to address a specific microservice from any other microservice in a uniform way, regardless of the number of replicas the microservice scales to. You plan to implement this solution on Google Kubernetes Engine. What should you do?
    ---
    Deploy each microservice as a Deployment. Expose the Deployment in the cluster using a Service, and use the Service DNS name to address it from other microservices within the cluster.

# 195

    You are building an application that uses a distributed microservices architecture. You want to measure the performance and system resource utilization in one of the microservices written in Java. What should you do?
    ---
    Instrument the service with Cloud Profiler to measure CPU utilization and method-level execution times in the service.

# 196

    Your team is responsible for maintaining an application that aggregates news articles from many different sources. Your monitoring dashboard contains publicly accessible real-time reports and runs on a Compute Engine instance as a web application. External stakeholders and analysts need to access these reports via a secure channel without authentication. How should you configure this secure channel?
    ---
    Add an HTTP(S) load balancer in front of the monitoring dashboard. Set up a Google-managed SSL certificate on the load balancer for traffic encryption.

    https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs

# 197

    You are planning to add unit tests to your application. You need to be able to assert that published Pub/Sub messages are processed by your subscriber in order. You want the unit tests to be cost-effective and reliable. What should you do?
    ---
    Use the Pub/Sub emulator

# 198

    You have an application deployed in Google Kubernetes Engine (GKE) that reads and processes Pub/Sub messages. Each Pod handles a fixed number of messages per minute. The rate at which messages are published to the Pub/Sub topic varies considerably throughout the day and week, including occasional large batches of messages published at a single moment.

    You want to scale your GKE Deployment to be able to process messages in a timely manner. What GKE feature should you use to automatically adapt your workload?
    ---
    Horizontal Pod Autoscaler based on an external metric
    https://cloud.google.com/kubernetes-engine/docs/concepts/custom-and-external-metrics#external-metrics

# 199

    You are using Cloud Run to host a web application. You need to securely obtain the application project ID and region where the application is running and display this information to users. You want to use the most performant approach. What should you do?
    ---
    Use HTTP requests to query the available metadata server at the http://metadata.google.internal/ endpoint with the Metadata-Flavor: Google header.

# 200

    You need to deploy resources from your laptop to Google Cloud using Terraform. Resources in your Google Cloud environment must be created using a service account. Your Cloud Identity has the roles/iam.serviceAccountTokenCreator Identity and Access Management (IAM) role and the necessary permissions to deploy the resources using Terraform. You want to set up your development environment to deploy the desired resources following Google-recommended best practices. What should you do?
    ---
    B. 1. Run the following command from a command line: gcloud config set auth/impersonate_service_account service-account-name@project.iam.gserviceacccount.com.
    2. Set the GOOGLE_OAUTH_ACCESS_TOKEN environment variable to the value that is returned by the gcloud auth print-access-token command. 

# 301

    You are developing a public web application on Cloud Run. You expose the Cloud Run service directly with its public IP address. You are now running a load test to ensure that your application is resilient against high traffic loads. You notice that your application performs as expected when you initiate light traffic. However, when you generate high loads, your web server runs slowly and returns error messages. How should you troubleshoot this issue?
    ---
    Check whether the Cloud Run service has scaled to a number of instances that equals the max-instances value. If necessary, increase the max-instances value.

# 302

    You are developing a new image processing application that needs to handle various tasks, such as resizing, cropping, and watermarking images. You also need to monitor the workflow and ensure that it scales efficiently when there are large volumes of images. You want to automate the image processing tasks and workflow monitoring with the least effort. What should you do?
    ---
    Implement Workflows to orchestrate the image processing tasks. Use Cloud Logging for workflow monitoring.

    GCP workflows https://cloud.google.com/workflows#all-features

# 303

    You are developing a web application that will be deployed to production on Cloud Run. The application consists of multiple microservices, some of which will be publicly accessible and others that will only be accessible after authentication by Google identities. You need to ensure that only authenticated users can access the restricted services, while allowing unrestricted access to the public services of the application. You want to use the most secure approach while minimizing management overhead and complexity. How should you configure access?
    ---
    Configure separate Cloud Run services for the public and restricted microservices. Enable Identity-Aware Proxy (IAP) only for the restricted services, and configure the Cloud Run ingress settings to ‘Internal and Cloud Load Balancing’.

# 304

    You are the lead developer for a company that provides a financial risk calculation API. The API is built on Cloud Run and has a gRPC interface. You frequently develop optimizations to the risk calculators. You want to enable these optimizations for select customers who registered to try out the optimizations prior to rolling out the optimization to all customers. Your CI/CD pipeline has built a new image and stored it in the Artifact Registry.

    Which rollout strategy should you use?
    ---
    Migrate the traffic to the new service by using a feature flag for registered customers.
    ---
    Migrate the traffic to the new service by setting Cloud Run’s traffic split based on the percentage of registered customers
    WRONG because "for select customers who registered"

# 305

    Your ecommerce application has a rapidly growing user base, and it is experiencing performance issues due to excessive requests to your backend API. Your team develops and manages this API. The Cloud SQL backend database is struggling to handle the high demand, leading to latency and timeouts. You need to implement a solution that optimizes API performance and improves user experience. What should you do?
    ---
    A. Use Apigee to expose your API. Use Memorystore for Redis to cache frequently accessed data. Implement exponential backoff in the application to retry failed requests.

    https://cloud.google.com/apigee/docs/api-platform/get-started/what-apigee
    ---
    Use Cloud Load Balancing to expose your API. Increase the memory for the database instances to handle more concurrent requests. Implement a custom rate-limiting mechanism in your application code to control API requests.
    WRONG because increase memory not long term solution

# 306

    You need to deploy a new feature into production on Cloud Run. Your company’s SRE team mandates gradual deployments to avoid large downtimes caused by code change errors. You want to configure this deployment with minimal effort. What should you do?
    ---
    Deploy the feature with “Serve this revision immediately” unchecked, and configure the new revision to serve a small percentage of traffic. Check for errors, and increase traffic to the revision as appropriate.

    Canary deployment
    ---
    Configure the application code to send a small percentage of users to the newly deployed revision. 
    WRONG because change client code very complex

# 307

    You are developing an external-facing application on GKE that provides a streaming API to users. You want to offer two subscription tiers, “basic" and “premium", to users based on the number of API requests that each client application is allowed to make each day. You want to design the application architecture to provide subscription tiers to users while following Google-recommended practices. What should you do?
    ---
    A. 
    1. Configure the service on GKE as a backend to an Apigee proxy.
    2. Provide API keys to users to identify client applications.
    3. Configure a Quota policy in Apigee for API keys based on the subscription tier.

# 308

    Your organization has users and groups configured in an external identity provider (IdP). You want to leverage the same external IdP to allow Google Cloud console access to all employees. You also want to personalize the sign-in experience by displaying the user's name and photo when users access the Google Cloud console. What should you do?
    ---
    Configure workforce identity federation with the external IdP, and set up attribute mapping.

# 309

    You are developing a new API that creates requests on an asynchronous message service. Requests will be consumed by different services. You need to expose the API by using a gRPC interface while minimizing infrastructure management overhead. How should you deploy the API?
    ---
    Deploy your API as a Cloud Run service. Create a Pub/Sub topic, and configure your API to push messages to the topic.

# 310

    You are about to deploy an application hosted on a Compute Engine instance with Windows OS and Cloud SQL. You plan to use the Cloud SQL Auth Proxy for connectivity to the Cloud SQL instance. You plan to follow Google-recommended practices and the principle of least privilege. You have already created a custom service account. What should you do next?
    ---
    A. Create and assign a custom role with the cloudsql.instances.connect permission to the custom service account. Adjust the Cloud SQL Auth Proxy start command to specify your instance connection name.
    ---
    B: Grant the custom service account the roles/cloudsql.client role. Adjust the Cloud SQL Auth Proxy start command to use the --unix-socket CLI option.
    WRONG because 
    --unix-socket: not for window
    cloudsqlclient = connect + get

# 311

    You are developing a secure document sharing platform. The platform allows users to share documents with other users who may be external to their organization. Access to these documents should be revoked after a configurable time period. The documents are stored in Cloud Storage. How should you configure Cloud Storage to support this functionality?
    ---
    Generate a signed URL for each document the user wants to share.

# 312

    You work for an environmental agency in a large city. You are developing a new monitoring platform that will capture air quality readings from thousands of locations in the city. You want the air quality reading devices to send and receive their data payload to the newly created RESTful backend systems every minute by using a curl command. The backend systems are running in a single cloud region and are using Premium Tier networking. You need to connect the devices to the backend while minimizing the daily average latency, measured by using Time to First Byte (TTFB). How should you build this service?
    ---
    D. 
    1 Run the air quality devices' backends in a managed instance group.
    2. Create an external Application Load Balancer, and connect it to the managed instance group.
    3. Configure a connection between the air quality devices and the Application Load Balancer.

# 313

    Your infrastructure team is responsible for creating and managing Compute Engine VMs. Your team uses the Google Cloud console and gcloud CLI to provision resources for the development environment. You need to ensure that all Compute Engine VMs are labeled correctly for compliance reasons. In case of missing labels, you need to implement corrective actions so the labels are configured accordingly without changing the current deployment process. You want to use the most scalable approach. What should you do?
    ---
    Use a Cloud Audit Logs trigger to invoke a Cloud Function when a Compute Engine VM is created. Check for missing labels and assign them if necessary.

# 314

    You are developing a discussion portal that is built on Cloud Run. Incoming external requests are routed through a set of microservices before a response is sent. Some of these microservices connect to databases. You need to run a load test to identify any bottlenecks in the application when it is under load. You want to follow Google-recommended practices. What should you do?
    ---
    Configure Cloud Trace to capture the requests from the load testing clients. Review the timings in Cloud Trace.

# 315

    Your team currently uses Bigtable as their database backend. In your application's app profile, you notice that the connection to the Bigtable cluster is specified as single-cluster routing, and the cluster’s connection logic is configured to conduct manual failover when the cluster is unavailable. You want to optimize the application code to have more efficient and highly available Bigtable connectivity. What should you do?
    ---
    
# 316

    You work for an ecommerce company. Your company is migrating multiple applications to Google Cloud, and you are assisting with the migration of one of the applications. The application is currently deployed on a VM without any OS dependencies. You have created a Dockerfile and used it to upload a new image to Artifact Registry. You want to minimize the infrastructure and operational complexity. What should you do?
    ---
    Deploy the image to Cloud Run.

# 317

    You recently deployed an Apigee API proxy to your organization across two regions. Both regions are configured with a separate backend that is hosting the API. You need to configure Apigee to route traffic to the appropriate local region backend. What should you do?
    apigee region 1 -> BE region 1
    apigee region 2 -> BE region 2
    ---
    Configure a TargetServer for each region's backend host names. Configure the API proxy to choose the TargetServer based on the system.region.name flow variable.

# 318

    You are a developer that works for a local concert venue. Customers use your company’s website to purchase tickets for events. You need to provide customers with immediate confirmation when a selected seat has been reserved. How should you design the ticket ordering process?
    ---
    Submit the seat reservation in an HTTP POST request to an Application Load Balancer. Configure the Application Load Balancer to distribute the request to a Compute Engine managed instance group that processes the reservation.
    ---
    pub/sub and async solution is WRONG because *immediate*

# 319

    You work for a financial services company that has a container-first approach. Your team develops microservices applications. You have a Cloud Build pipeline that creates a container image, runs regression tests, and publishes the image to Artifact Registry. You need to ensure that only containers that have passed the regression tests are deployed to GKE clusters. You have already enabled Binary Authorization on the GKE clusters. What should you do next?
    ---
    Create an attestor and a policy. Create an attestation for the container images that have passed the regression tests as a step in the Cloud Build pipeline.

# 320

    You have an application running in production on Cloud Run. Your team needs to change one of the application’s services to return a new field. You want to test the new revision on 10% of your clients using the least amount of effort. You also need to keep your service backward compatible.

    What should you do?
    ---
    Update the current service with the new changes. Deploy the new revision with no traffic allocated. Split the traffic between the current service and the new revision.

# 321

    Your team plans to use AlloyDB as their database backend for an upcoming application release. Your application is currently hosted in a different project and network than the AlloyDB instances. You need to securely connect your application to the AlloyDB instance while keeping the projects isolated. You want to minimize additional operations and follow Google-recommended practices. How should you configure the network for database connectivity?
    ---
    Provision a Shared VPC project where both the application project and the AlloyDB project are service projects.

# 322

    You have an on-premises containerized service written in the current stable version of Python 3 that is available only to users in the United States. The service has high traffic during the day and no traffic at night. You need to migrate this application to Google Cloud and track error logs after the migration in Error Reporting. You want to minimize the cost and effort of these tasks. What should you do?
    ---
    Deploy the code on Cloud Run. Configure your code to write errors to standard error.

# 323

    Your team is developing a new application that is packaged as a container and stored in Artifact Registry. You are responsible for configuring the CI/CD pipelines that use Cloud Build. Containers may be pushed manually as a local development effort or in an emergency. Every time a new container is pushed to Artifact Registry, you need to trigger another Cloud Build pipeline that executes a vulnerability scan. You want to implement this requirement using the least amount of effort. What should you do?
    ---
    Configure Artifact Registry to publish a message to a Pub/Sub topic when a new image is pushed. Configure the vulnerability scan pipeline to be triggered by the Pub/Sub message.

# 324

    You have an application running on a GKE cluster. Your application has a stateless web frontend, and has a high-availability requirement. Your cluster is set to automatically upgrade, and some of your nodes need to be drained. You need to ensure that the application has a serving capacity of 10% of the Pods prior to the drain. What should you do?
    ---
    Configure the Pod replica count to be 10% more than the current replica count.

# 325

    You are designing an application that shares PDF files containing time-sensitive information with users. The PDF files are saved in Cloud Storage. You need to provide secure access to the files.

    You have the following requirements:
    • Users should only have access to files that they are allowed to view.
    • Users should be able to request to read, write, or delete the PDF files for 24 hours.

    Not all users of the application have a Google account. How should you provide access to data objects?
    ---
    generate signed URLs

# 326

# 327

    Your infrastructure team uses Terraform Cloud and manages Google Cloud resources by using Terraform configuration files. You want to configure an infrastructure as code pipeline that authenticates to Google Cloud APIs. You want to use the most secure approach and minimize changes to the configuration. How should you configure the authentication?
    ---
    Configure Terraform Cloud to use workload identity federation to authenticate to the Google Cloud APIs.
    https://cloud.google.com/docs/terraform/install-configure-terraform

# 328

    Your team has created an application that is hosted on a GKE cluster. You need to connect the application to a REST service that is deployed in two GKE clusters in two different regions. How should you set up the connection and health checks? (Choose two.)
    ---
    Use Cloud Service Mesh with sidecar proxies to connect the application to the REST service. (like anthos)
    Configure the REST service's firewall to allow health checks originating from the GKE check probe’s IP ranges. (must allow to health check)

# 329 

    You are using the latest stable version of Python 3 to develop an API that stores data in a Cloud SQL database. You need to perform CRUD operations on the production database securely and reliably with minimal effort. What should you do?
    ---
    C. 1. Use the Cloud SQL connector library for Python to connect to the Cloud SQL database through a Cloud SQL Auth Proxy.
    2. Grant an IAM role to the service account that includes the cloudsql.instances.connect permission.

# 330

    Your company manages an application that captures stock data in an internal database. You need to create an API that provides real-time stock data to users. You want to return stock data to users as quickly as possible, and you want your solution to be highly scalable. What should you do?
    ---
    Create a Memorystore for Redis instance, and use this database to store the most accessed stock data. Query this instance first when user requests are received, and fall back to the internal database.

# 331

    You are designing a microservices architecture for a new application that will be deployed on Cloud Run. The application requires high-throughput communication between the internal microservices. You want to use the most effective, lowest latency communication protocol for this application. What should you do?
    ---
    Configure the Cloud Run service to use HTTP/2. Implement gRPC for communication between the microservices. Use streaming gRPCs when a large amount of data has to be sent.

# 332

    Your company recently modernized their monolith ecommerce site to a microservices application in GKE. Your team uses Google Cloud's operations suite for monitoring and logging. You want to improve the logging indexing and searchabilty in Cloud Logging across your microservices with the least amount of effort. What should you do?
    ---
    Update your microservices code to emit logs in JSON format.
    "logging indexing" and "searchabilty" => index and search by json field

# 333

    You recently developed an application that will be hosted on Cloud Run. You need to conduct a load test. You want to analyze the load test logs second by second to understand your Cloud Run service's response to rapid traffic spikes. You want to minimize effort. How should you analyze the logs?
    ---
    Analyze the log data in BigQuery by configuring a BigQuery log sink with the appropriate inclusion filter for your application.
    
# 334

    You are a developer of a new customer-facing help desk chat service that is built on Cloud Run. Your customers use the chat option on your website to get support. The application saves each transcript as a text file with a unique identifier in a Cloud Storage bucket. After the conversation is done and before the chat window is closed, the customer receives a link to the chat transcript. You want to provide access to the chat transcript link for 2 hours. You need to configure this access using an approach that prioritizes security and follows Google-recommended practices. What should you do?
    ---
    Create a signed URL for each text file that expires after 2 hours.

# 335

    You are responsible for managing the security of internal applications in your company. The applications are deployed on Cloud Run, and use Secret Manager to store passwords needed to access internal databases. Each application can cache secrets for up to 15 minutes. You need to determine how to rotate the secrets. You want to avoid application downtime. What should you do?
    ---
    Store the new password in the secret. Reference the latest version of any secret required, and cache the secret for 15 minutes.

# 336

    You are developing a new ecommerce application that uses Cloud Functions. You want to expose your application's APIs to public users while maintaining a high level of security. You need to ensure that only authorized users can access your APIs and that all API traffic is encrypted and protected from unauthorized access. You want to use the most scalable and secure approach. What should you do?
    ---
    Deploy your Cloud Functions behind an Apigee proxy and use Apigee’s authentication and authorization features to secure your APIs.

# 337 
    
    You are deploying a microservices application to GKE. One microservice needs to download files from a Cloud Storage bucket. You have an IAM service account with the Storage Object Viewer role on the project with the bucket. You need to configure your application to access the Cloud Storage bucket while following Google-recommended practices. What should you do?
    ---
    Create a Kubernetes service account. Use an IAM policy to bind the IAM service account to a Kubernetes service account. Annotate the Kubernetes ServiceAccount object with the name of the bound IAM service account. Assign the Kubernetes ServiceAccount to the Pods that need to access the bucket.

# 338

# 339

    You are developing a new ecommerce website for your company. You want customers to receive a customized email notification when they place an order. You need to configure this email service while minimizing deployment effort. What should you do?
    ---
    Create a Cloud Function that is triggered by a create type event in Firestore,
    Minimal deployment effort:

# 340

    You are developing an online chat application where users can upload profile pictures. Uploaded profile pictures must comply with content policies. You need to detect inappropriate images and label those images automatically when they are uploaded. In the future, this process will need to be expanded to include additional processing tasks such as watermarking and image compression.

    You want to simplify orchestration and minimize operational overhead of the image scanning and labeling steps while also ensuring that additional steps can be added and removed easily later on. What should you do?
    ---
    Save user-uploaded images to a Cloud Storage bucket. Create an Eventarc trigger that connects the bucket to the Workflows event receiver when a new image is uploaded. Create a workflow in Workflows with multiple Cloud Functions that call the Vision API to process each new uploaded image.

    process image => cloud workflows

# 341

    Your application named ecom-web-app is deployed in three GKE clusters: ecom-web-app-dev, ecom-web-app-qa, and ecom-web-app-prod. You need to ensure that only trusted container images are deployed to the ecom-web-app-prod GKE cluster in the production environment while following Google-recommended practices. What should you do?
    ---
    Set up Binary Authorization, and define cluster-specific rules in clusterAdmissionRules nodes in the policy YAML file.

# 342

    You are responsible for improving the security of your Cloud Run services to protect these services against supply chain threats. You need to ensure that there are adequate security controls such as SLSA Level 3 builds for container images and non-falsifiable provenance for container images by using Google Cloud tools. What should you do?
    ---
    Use Cloud Build to build container images. Configure a Binary Authorization policy on the Cloud Run job.

# 343

    There are three teams developing an ecommerce application in the same Google Cloud project. Team A will build a set of RESTful APIs that exposes some core functionalities for the application. Team B and Team C will make requests to those APIs in their downstream processes running on Cloud Run services. You need to propose a solution for exposing the APIs in a way that maximizes security and minimizes management overhead for the three teams. How should you design this solution?
    ---
    1. Team A uses service accounts to authorize Cloud API Gateway. Team B and Team C each create a service account that has access to the APIs.
    2. Team B and Team C access the APIs in the Cloud Run service running their processes by using the service accounts attached to their service.

# 344

    Your application team is developing an ecommerce application. Your team has developed a new functionality that has a dependency on a third-party service. This third-party service will be deployed in a few days. However, you have been unable to ensure the reliability of this service. You need to choose a deployment strategy for the ecommerce application that will avoid disruption and can be rolled back quickly if issues are discovered. What should you do?
    ---
    Use a feature flag to enable the new functionality to users on demand. Gradually enable the new functionality to more users.

# 345

    You work for an organization that manages an ecommerce site. Your application is deployed behind an external Application Load Balancer. You need to test a new product recommendation algorithm. You plan to use A/B testing to determine the new algorithm’s effect on sales in a randomized way. How should you test this feature?
    ---
    Split traffic between versions using weights.

# 346

    You maintain a CI/CD pipeline for an application running on GKE. You use Cloud Build to create container images and push the images to Artifact Registry. When you build the image, you use the latest tag in your pipeline.

    You recently had to roll back a deployment 24 hours after rollout. The rollback process was difficult because the latest tag had been overwritten. You need to prevent this issue in the future. You want to use the most efficient approach. What should you do?
    ---
    Build a separate Docker image for each new version of the application, and tag it with the version number.

# 347

    You are designing a microservices application on GKE that will expose a public API to users. Users will interact with the application by using OAuth 2.0, and illegitimate requests should receive a 403 response code. You need the API to be resilient against distributed denial of service (DDoS) attacks and critical security risks such as SQL injection (SQL) and cross-site scripting (XSS).

    You want to design the application's architecture while following Google-recommended practices. What should you do?
    ---
    Use an external Application Load Balancer with Cloud Armor, and configure the load balancer to forward requests to Apigee to check the validity of the API requests. Configure GKE as the application's backend.

    Cloud Armor: DDoS,  helping to mitigate threats like SQL injection and XSS.
    Apigee: ratelimit, 
    ---
    reCAPTCHA Enterprise: Great for bot mitigation
    WRONG

# 348

    You are compiling a compliance report on vulnerability metadata for a specific set of images identified by Artifact Analysis. Metadata from images scanned more than 30 days ago are missing from the compliance report. You need to access the vulnerability metadata for these older images. What should you do?
    ---
    Push or pull the images from Artifact Registry.
    when u push or pull img, it re-trigger a fresh vulnerability scan

# 349

    Your team runs a Python job that reads millions of customer record files stored in a Cloud Storage bucket. To comply with regulatory requirements, you need to ensure that customer data is immediately deleted once the job is completed. You want to minimize the time required to complete this task. What should you do?
    ---
    Add a final step in the job that deletes all the objects in the bucket in bulk by using batch requests to the Cloud Storage API.
    ---
    gsutil rm ...
    WRONG because "immediately"

# 350

    You are a developer at a regulated financial company and are the lead of a risk calculation application that is running on Cloud Run. Binary Authorization for Cloud Run has been enabled as an organization policy, and there is one attestor. All applications in the company are attested. Each application's image is deployed as part of a CI/CD pipeline during a 1-hour change window at 11 PM local time. There is a new security issue that requires you to deploy a critical fix before the next change window. You have created a new image with the fix, and your manager has approved the image in an email message. What should you do?
    ---
    Use the breakglass approach to deploy the image.
    a method of bypassing security controls that normally guard a system or service

# 351

    You have a Cloud Run service that needs to connect to a Cloud SQL instance in a different project. You provisioned the Cloud Run service account with the Cloud SQL Client IAM role on the project that is hosting Cloud SQL. However, when you test the connection, the connection fails. You want to fix the connection failure while following Google-recommended practices. What should you do?
    ---
    C. Enable the Cloud SQL Admin API in both projects.

# 352

    You developed a Python script that retrieves information from files that are uploaded to Cloud Storage and writes the information to Bigtable. You have completed testing on your local environment and created the python-script service account with the Bigtable User IAM role. You want to deploy the code with the appropriate authentication while following Google-recommended practices. What should you do?
    ---
    A. 1. Deploy your code to Cloud Functions. Create a Cloud Storage trigger.
    2. Configure IAM binding for authentication.
    ---
    B. 1. Deploy your code to Cloud Functions. Create a Cloud Storage trigger.
    2. Create a service account key for authentication
    WRONG because sa key not necessary, key can be lost

# 353

    You are a developer at a large organization. Your team uses Git for source code management (SCM). You want to ensure that your team follows Google-recommended best practices to manage code to drive higher rates of software delivery. Which SCM process should your team use?
    ---
    Each developer creates a branch for their own work, commits their changes to their branch, and merges their code into the main branch daily.

# 354

    You work for an ecommerce company. You are developing a new application with the following requirements:
    • The application must have access to the most up-to-date data at all times.
    • Due to company policy, data older than 30 days must be automatically deleted.

    You need to determine which service should host the database, and how to configure the data deletion. You want to use the most efficient solution. What should you do?
    ---
    Configure Bigtable to host the database. Create a garbage collection policy in Bigtable that deletes data older than 30 days.

# 355

    Your application in production has recently been experiencing reliability issues, and you are unsure how the application will behave in the event of an unexpected failure. You want to assess the application's resilience. What should you do?
    ---
    B. Perform chaos engineering by intentionally introducing failures into the system. Observe how the application behaves, and ensure that it is able to recover from a failure.
    Because Chaos engineering is the recommended approach for assessing resilience in production systems.
    ---
    D. Load testing helps evaluate performance under high demand, not how the system recovers from failure or handles instability.

# 356

    Your team is responsible for developing multiple microservices. These microservices are deployed in Cloud Run and connected to a Cloud SQL instance. You typically conduct tests in a local environment prior to deploying new features. However, the external IP was recently removed from your Cloud SQL instance, and you are unable to perform the tests. You need to connect to the database to conduct tests with the most updated data. You want to follow Google-recommended practices. What should you do?
    ---
    Create a VM in the same VPC as the Cloud SQL instance. Connect to the VM by using Identity-Aware Proxy for TCP forwarding. Install and configure the Cloud SQL Auth Proxy.

# 357

    You have an application running on Cloud Run that receives a large volume of traffic. You need to deploy a new version of the application. You want your deployment process to minimize the risk of downtime while following Google-recommended practices. What should you do?
    ---
    "minimize the risk of downtime"
    Use traffic splitting to have a small percentage of users test out new features on the new revision of the application on the production Cloud Run service. If performance meets expectations, gradually increase the percentage of users until it reaches 100%
    ---
    Use Cloud Build to create a pipeline, and configure a test stage before the deployment stage. When all tests pass, deploy the application to Cloud Run, and direct 100% of users to this new version of the application. Roll back if any issues are detected.
    ROLLBACK => downtime

# 358

    You are a developer at an ecommerce company. You are tasked with developing a globally consistent shopping cart for logged-in users across both mobile and desktop clients. You need to configure how the items that are added to users’ carts are stored. How should you configure this cart service?
    ---
    Store the carts in a separate Firestore document, and configure each user ID as the document's key.

# 359

    You are deploying a containerized application to GKE. You have set up a build pipeline by using Cloud Build that builds a Java application and pushes the application container image to Artifact Registry. Your build pipeline executes multiple sequential steps that reference Docker container images with the same layers.

    You notice that the Cloud Build pipeline runs are taking longer than expected to complete. How should you optimize the Docker image build process
    ---
    Specify the cached image by adding the --cache-from argument in your build config file with the image as a cache source.

# 360

    You are developing a dashboard that aggregates temperature readings from thousands of IoT devices monitoring a city's ambient temperature. You expect a large amount of viewing traffic resulting in a large amount of data egress once the dashboard is live. The dashboard temperature display data doesn't need to be real-time and can tolerate a few seconds of lag. You decide to deploy Memorystore for Redis as the storage backend. You want to ensure that the dashboard will be highly available. How should you configure the service in Memorystore for Redis?
    ---
    Configure Memorystore to use read replicas.

# 361

    You are responsible for developing a new ecommerce application that is running on Cloud Run. You need to connect your application to a Cloud SQL database that is in a separate project. This project is on an isolated network dedicated to multiple databases without a public IP. You need to connect your application to this database. What should you do?
    ---
    Create a subnet on your VPC. Create a Serverless VPC Access connector on your project using the new subnet. In Cloud Run, create a Cloud SQL connection. Use Cloud SQL Language Connectors to interact with the database
