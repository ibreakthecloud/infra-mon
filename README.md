# Infra-Mon

Infra-Mon or Infrastructure Monitoring is a metric ingestion server. It records the cpu and memory percentage utilization against IP address of the client in-memory. It also generates the report with list of all the unique IP addresses/machines with it's highest CPU and memory utilization.


## Features
It exposes two endpoints
- `/metrics`
    - **methods**: `POST`
- `/report`
    - **methods:** `GET`
    
## Usage
By default, this server serves on port `8080`, but can be configured on the basis of environment variable `PORT`.

Running `make` will automatically generate a Docker image. Dockerfile is multi-staged build and exposes the server at port `8080`


### Ingestion

**Method**: POST

**Path**: /metrics

**Headers**: `content-type: application/json`

**Request Payload**: 
```
{
    "percentage_cpu_used": 80,
    "percentage_memory_used": 45
}
```

### Report

**Method**: GET

**Path**: /report

**Headers**: `content-type: application/json`

**ResponsePayload**: 
```
[
    {
        "ip": "10.20.30.40",
        "max_cpu": 65,
        "max_memory": 85
    },
    {
        "ip": "49.37.83.70",
        "max_cpu": 71,
        "max_memory": 90
    }
]
```

## Development

Running the below command will start the local server at port `8080`

```go
go run main.go
```