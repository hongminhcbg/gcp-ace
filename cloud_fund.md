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
