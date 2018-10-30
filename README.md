# Infoblox Assignment

Parse csv file and api to return contact detailks of the name passed in the URL parameter

## Getting Started

Server will listen on post :8080 with firstname as the endpoint for showing the contact details
```
http://localhost:8080/{firstname}
```

### Prerequisites

You need to install docker to create an image from the Dockerfile in the respository

```
docker build -t infoblox .
```

### Installing

The project will create an api which will use gorilla mux for routing, tealeg/xlsx to read xlsx files
and dep tool for dependency management.

Create docker image from dockerfile

```
docker build -t infoblox .
```

Run the container from the image to run server.

```
docker run --publish 6060:8080 --name infoblox --rm infoblox
```

Hit the endpoint to fetch the data on basis of key argument as:

```
http://localhost:6060/one
```
### Output

```
[{"key":"one","value":1}]
```

## Built With

* [Gorilla/mux](https://github.com/gorilla/mux) - Gorilla mux routers
* [Dep](https://github.com/golang/dep) - dependency Management

## Authors

* **Himanshu Chaudhary** - *Initial work* - [Infoblox-Assignment](https://github.com/Himanshuxone/infoblox-assigment)

