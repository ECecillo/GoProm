# GoProm

A simple API made in GO to see how Prometheus and Grafana can be setup and play around to see how they work.

# Running the application locally

Pre-requisite:

- Docker

```
docker compose up -d
```

The Dockerfile build an image of our application and the docker-compose run everything along with the Prometheus server and Grafana.

## Setting up Grafana

1. Go to localhost:3000 to access Grafana Interface
2. Default credential are `admin` for the login and the password.
3. Click skip if your don't want to change the default password.
4. On the Homepage, click in the searchbar located at the very top of the page.
5. Enter "Data Sources" and press ENTER.
6. Click on "Add new data source".
7. Then in the filter input above listed source, search "Prometheus".
8. Click on it.
9. Then once you are in the configuration page of the Prometheus source, go to the "Connection" section.
10. In the "Prometheus server URL", put `http://host.docker.internal:9090` since Grafana and Prometheus are not in the same network.
11. Then scroll to the button and click "Save and Test", everything should be green ✅ !
12. Finally, navigate to the Dashboard page using the side panel of the app and create your stuff :D .

## Loadtesting

Since we want to monitor and maybe in the future optimize some endpoint to support a load of request we need something that can help us reproduce this flow of request into our app.

For that we will use [K6](https://k6.io/) which will mimic how our users can send request to our API to see how the server can handle the traffic.

However, since I didn't wanted to include the image of app in my docker compose, it's up to you to see how you can install this tool [here](https://grafana.com/docs/k6/latest/set-up/install-k6/).

# Side notes

The application can be run using the docker image built using `Dockerfile`.

The `docker-compose.yml` is only used in development, it run our application and a grafana application alongisde to visualize prometheus metrics.
