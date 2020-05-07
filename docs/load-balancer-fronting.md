# Front Cloud Run with a HTTPs Load Balancer

These instructions give you an example of how to front Cloud Run with a Load Balancer. This can be used with Cloud Run Fully Managed to allow your application to sit behind third-party services like Cloud CDN, Akamai, Cloudflare, etc and to use your own SSL Certificiates.

### Requirements

* These instructions assume you are using [Cloud Shell](https://cloud.google.com/shell/docs) for deployment.
* You can use your local machine, but ensure you have gcloud setup, and authenticated.
* You should already have your Cloud Run Service deployed. We are using the service name `hello-web-server-cloud-run` for these example steps.

### Deploy the Cloud Run App
Do this via the GitHub Workers, once complete get the URL of the deployment
```
gcloud run services list --platform=managed --region=europe-west1
# Copy the HOSTNAME ONLY of the service in scope into the below variable
RUN_APP_HOSTNAME=<Hostname>
```

#### Reserve IP Addresses
```
gcloud compute addresses create hello-web-server-cloud-run-ipv4 --global --ip-version IPV4
```

####  Create DNS Records with DNS Provider
Create the DNS Records with your DNS provider using the above reserved static IP.
```
gcloud compute addresses describe hello-web-server-cloud-run-ipv4 --global
```

####  Create Internet Endpoint Group for europe-west2 Cloud Functions
```
# Create the Internet Network Endpoint Group
gcloud compute network-endpoint-groups create hello-web-server-cloud-run-cloudrun --network-endpoint-type="internet-fqdn-port" --global
gcloud compute network-endpoint-groups update hello-web-server-cloud-run-cloudrun --add-endpoint="fqdn=$RUN_APP_HOSTNAME,port=443" --global
```

####  Create Load Balancer Backend
```
gcloud compute backend-services create hello-web-server-cloud-run-cloudrun-backend --global --protocol=HTTP2 --timeout=20s
gcloud compute backend-services update hello-web-server-cloud-run-cloudrun-backend --custom-request-header "Host: $RUN_APP_HOSTNAME" --global
gcloud compute backend-services add-backend hello-web-server-cloud-run-cloudrun-backend --network-endpoint-group "hello-web-server-cloud-run-cloudrun" --global-network-endpoint-group --global
```

#### Create URL Map
```
gcloud compute url-maps create hello-web-server-cloud-run --default-service hello-web-server-cloud-run-cloudrun-backend --global
```

#### Create SSL Certificate
Create an SSL Certificate to use with the Load Balancer, the below command expects the following files `secrets/<fqdn>.crt` and `secrets/<fqdn>.key`
```
gcloud compute ssl-certificates create hello-web-server-cloud-run --certificate secrets/$LOAD_BALANCER_FQDN.crt --private-key secrets/$LOAD_BALANCER_FQDN.key
```

#### Create SSL Proxy
```
gcloud compute target-https-proxies create hello-web-server-cloud-run-frontend --url-map=hello-web-server-cloud-run --ssl-certificates=hello-web-server-cloud-run --global
```

#### Create Forwarding Rule
```
gcloud compute forwarding-rules create hello-web-server-cloud-run-frontend --ip-protocol=TCP --ports=443 --global --target-https-proxy=hello-web-server-cloud-run-frontend --address=hello-web-server-cloud-run-ipv4
```