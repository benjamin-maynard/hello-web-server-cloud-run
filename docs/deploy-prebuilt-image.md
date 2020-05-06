# Deploy Pre-Built Image

These instructions give you a quick start with Cloud Run. You will deploy a pre-built image of the hello-web-server-cloud-run service to your own Google Cloud Project.

### Requirements

* These instructions assume you are using [Cloud Shell](https://cloud.google.com/shell/docs) for deployment.
* You can use your local machine, but ensure you have gcloud setup, and authenticated.

## Deployment Steps

### Setup Environment
```
GCLOUD_PROJECT=$(gcloud config get-value project)
CLOUD_RUN_SERVICE_NAME=hello-web-server-cloud-run
CONTAINER_IMAGE=gcr.io/bjm-public-container-images/hello-web-server-cloud-run:latest
# See https://cloud.google.com/run/docs/locations for below regions
CLOUD_RUN_REGION=europe-west1
```

### Enable Cloud Run API's
```
gcloud services enable run.googleapis.com
```

### Deploy Cloud Run Application
```
gcloud run deploy $CLOUD_RUN_SERVICE_NAME \
    --image=$CONTAINER_IMAGE \
    --platform=managed \
    --cpu=1 \
    --memory=512Mi \
    --port=8080 \
    --allow-unauthenticated \
    --region=$CLOUD_RUN_REGION \
    --project=$GCLOUD_PROJECT
```

## Access Environment
If successful, you should see an output in your console such as:
```
✓ Deploying new service... Done.                                                           
  ✓ Creating Revision...
  ✓ Routing traffic...
  ✓ Setting IAM Policy...
Done.
Service [hello-web-server-cloud-run] revision [hello-web-server-cloud-run-00001-fop] has been deployed and is serving 100 percent of traffic at https://hello-web-server-cloud-run-3lyhxu4vqa-ew.a.run.app
```

Browse to the URL in the output. You should see the expected JSON response.