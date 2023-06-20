## Description
Gometri is an implementation project for data query triggers, API data requests for opentelemetry. This project is a sample of the implementation of opentelemetry + golang. The schema here is that every result of a `ticket` request will be triggered to the opentelemetry metric and every query result data to the database is also the same

## Run Services
```
$ make run
```
or
```
$ make runDock
```
> If you dont have go you must install go, and if you dont have docker you must install docker

## License
[Mit License](https://github.com/rulanugrh/gometri/blob/main/LICENSE)