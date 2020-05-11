# hello-web-server-cloud-run

hello-web-server-cloud-run is a sample application written in Golang designed to be used as an example for deploying to Cloud Run.

It uses the `http` package to start a web server on a port defined by the $PORT variable. It loads the `K_SERVICE` and `K_REVISION` environment variables and returns them in JSON format when a user browses to the Cloud Run Service URL.