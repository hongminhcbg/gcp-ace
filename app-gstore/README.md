Step to deploy

    $ gcloud auth login
    - Create app engine with UI
    - Enable app engine admin api
    - grant role storage bucket viewer to "targe serice account" (u can see when run $gcloud app deploy)
    $ gcloud app deploy

Step to update

    - update app version in app.yaml (https://cloud.google.com/appengine/docs/standard/reference/app-yaml?tab=python)
    $ gcloud app deploy
    # split traffic
    $ gcloud app services set-traffic v3 --splits=20240715t215942=.5,20240715t221452=.5
