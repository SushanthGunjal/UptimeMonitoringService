# Uptime Monitoring System

* Checks whether the given url is up or not.
* The data used by the system is stored in the database.

## Tech stack used 


* Golang
* Mysql
    * Gorm as orm library
* Docker


## Installation

There are 3 ways to install the service

### 1. On local machine

Run the following commands
```
git clone https://github.com/SushanthGunjal/UptimeMonitoringService.git
cd UptimeMonitoringService
```

##
### Build

```
go build .
```

### Run
```
./UptimeMonitoringService
```

### 2. Using Docker

Run the following commands

```
git clone https://github.com/SushanthGunjal/UptimeMonitoringService.git
cd UptimeMonitoringService
```

### Build
```
docker build . -t UptimeMonitoringService  
```

### Run
```
docker run -p 8080:8080 UptimeMonitoringService  
```


### 3. Using Docker image
Pull the image from dockerhub by executing the following command
```
docker image pull sushanthgunjal/uptime_monitoring_service 
```

### Run
```
docker run -p 8080:8080 uptime_monitoring_service 
```



## API

### Check Url:

GET /urls/:id 

#### Response:


```
{
 "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0

}
```

### Add Url:

Add url to the database by 

POST/URL

#### Request:

```
{
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
}

```

#### Response:

```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0
}
```

### Update Url:

PATCH /url/:id

#### Request:

```
{
  “frequency”:                  60, 
  “status”:                     “active” 
}
```

#### Response:

```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  60, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0

}
```

### Delete URL:

DELETE /urls/:id

Following is the definition for the values
* url is checked at every fixed time(frequency) 
* crawl time is the time for which the system waits before giving up on the url
* failure threshold is the maximum failure count allowed for a particular Url. 


