# GoProm

A simple API made in go to see how to setup Prometheus and Grafana could work for monitoring.

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
11. Then scroll to the button and click "Save and Test", everything should be green !
12. Finally, navigate to the Dashboard page using the side panel of the app and create your stuff :D .

# Side notes

The application can be run using the docker image built using `Dockerfile`.

The `docker-compose.yml` is only used in development, it run our application and a grafana application alongisde to visualize prometheus metrics.
