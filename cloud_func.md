# Cloud func Faas

    - auto scale, scale very fast
    - default timeout 60s, MAX is 60 mins
    - pay for mem + cpu

cloud func concept:

    - event: upload obj to gcs
    - triggers: which func to trigger when an event happens?
    - function: take data event data perform action

Event trigger from:

    - gcs upload
    - pubsub
    - HTTP
    - stack driver logging


# Cloud run (container to production in seconds)

    - serverless
    - zero infrastructure management
    - using container 
    - gcloud run deploy ...
    - public internet -> cloud run -> trigger cloud func
    - deploy a image to prod
    - nolimit in languages

Cloud Run vs App engine

    Cloud Run gives developers more control over their environment as they can define memory and CPU allocations at a container level. App Engine manages resource configurations at the application level and can abstract many of the management tasks away, which is great for developers preferring a more managed environment
