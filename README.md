# Infra-Mon

Infra-Mon or Infrastructure Monitoring is a metric ingestion server. It records the cpu and memory percentage utilization against IP address of the client in-memory. It also generates the report with list of all the unique IP addresses/machines with it's highest CPU and memory utilization.


## Features
It exposes two endpoints
- `/metrics`
    - **methods**: `POST`
- `/report`
    - **methods:** `GET`
    
## Usage
By default this server serves on port `8080`. 

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
        "max_cpu": 95,
        "max_memory": 85
    },
    ...
]
```