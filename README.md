# GoProm

A simple API made in go to see how to setup Prometheus and Grafana could work for monitoring.

The Dockerfile build an image of our application and the docker-compose run everything along with the Prometheus server and Grafana.

You can pull the Dockerfile of the API [here](https://github.com/ECecillo/GoProm/pkgs/container/goprom)

# Running the application locally

Pre-requisite:

- Docker

```
docker compose up -d
```

# Quick tour

| Server Name  | Port  | Description  |
|---|---|---|
| API  | [:8080](localhost:8080)  | Golang Rest API  |
| Prometheus  | [:9090](localhost:9090)  | Metrics collector and querying language  |
| Grafana  | [:3000](localhost:3000)  | Monitoring and Analytics  |


## Setting up Grafana

1. Go to [localhost:3000](localhost:3000) to access Grafana Interface
2. Default credential are `admin` for the login and the password.
3. Click skip if your don't want to change the default password.
4. On the Homepage, click in the searchbar located at the very top of the page.
5. Enter "Data Sources" and press ENTER.
6. Click on "Add new data source".
7. Then in the filter input above listed source, search "Prometheus".
8. Click on it.
9. Then once you are in the configuration page of the Prometheus source, go to the "Connection" section.
10. In the "Prometheus server URL", put `http://prometheus:9090` since Grafana and Prometheus are in the same network `goprom-net`.
11. Then scroll to the button and click "Save and Test", everything should be green !
12. Finally, navigate to the Dashboard page using the side panel of the app and create your stuff :D .

## Loadtesting our API

For short, load testing fake user request to see how our api can handle the charge.

I will use [k6](https://grafana.com/docs/k6/latest/) for this but we could have gone for a more lightweight alternative such as [hey](https://github.com/rakyll/hey).

Both of them require that you install on your machin these tools, you can pull the Docker image for k6 but you need to make sure that it have access to the network where our api server is running.

Once k6 is install, run the following command to run http request on our API :

```shell
k6 run loadtesting.js
```

[Source](https://grafana.com/docs/k6/latest/get-started/running-k6/)



# Side notes

The application can be run using the docker image built using `Dockerfile`.

The `docker-compose.yml` is only used in development, it run our application and a grafana application alongisde to visualize prometheus metrics.
