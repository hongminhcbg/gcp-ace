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