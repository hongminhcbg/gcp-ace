I. Storage

    oganized and persistent storage

Category:

    - big query
    - cloud sql
    - cloud spaner 
    - datastore

Database type:

    Relational db OLTP -> cloud SQL, cloud spanner 
    relational OLAP db -> columnar storage, for admin and DA
    nosql db -> filestore (serverless transaction docs, mobile + web), bigtable (large db), not severless
    inmem database (cache) -> memstore, access in microsecond

Scenarios:

    start up quick -> datastore, filestore, filestore = datastore++ (multi devices)
    nosql, < 10Gi -> cloud datastore
    process 1Mi TPS -> cloud spanner
    process 1k TPS -> cloud SQL
    store ~1 peta bytes ->  big table nosql
    huge volume of IoT devices -> bigtable, nosql
    stream huge time series data (logs) -> bigtable, nosql

Cloud SQL

    
    


