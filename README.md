# Address Book

Parse csv file and api to return contact details of the name passed in the variable endpoint

## Getting Started

Server will listen on port :8080 with firstname as the endpoint for showing the contact details
```
http://localhost:8080/{firstname}
```

### Prerequisites

Install docker to create an image from the Dockerfile in the respository

```
docker build -t address-book .
```

### Installing

The project will create an api which will use gorilla mux for routing, encoding/csv to read csv file
and dep tool for dependency management.

Create docker image from dockerfile

```
docker build -t address-book .
```

Run the container from the image to run server.

```
docker run --publish 6060:8080 --name address-book --rm address-book
```

Hit the endpoint to fetch the data on basis of firstname as variable endpoint as:

```
http://localhost:6060/bob
```
### Output

```
[{"firstname":"BOb","lastname":"Williams","address":{"street":"234 2nd Ave.","city":"Tacoma","state":"WA"},"code":26}]
```

## Built With

* [Gorilla/mux](https://github.com/gorilla/mux) - Gorilla mux routers
* [Dep](https://github.com/golang/dep) - dependency Management

## Authors

* **Himanshu Chaudhary** - *Initial work* - [Address-Book](https://github.com/Himanshuxone/address-book)

