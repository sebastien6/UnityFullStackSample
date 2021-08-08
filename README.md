# UnityFullStackSample

## Design

### Backend

The backend is a golang GPRC microservice application with a RestAPI gateway. The gateway provides a web server to server the RestAPI call and the openapiv2 User Interface.

The GRPC listing application provide the service and store the data in a mongodb database.

A layer of abstraction was created between the GRPC service and the database making it easy in case of the replacement of the database by another one with limited code changes.

The api version and user input validation was declared as part of the proto file definition.

![base URL](https://github.com/sebastien6/UnityFullStackSample/blob/main/img/flow.png)

### unity app

- In the folder unity\Assets\Scripts\Runtime a new folder 'RestClient' was created with a file RestClient.cs that contain the code use to make the RestAPI calls.
- In the file, unity\Assets\Scripts\Test\TestListings.cs, a SerializeField baseUrl was added to set the URL for the backend RespAPI service to call. (example: http://localhost:8080/api/v1/games)

![base URL](https://github.com/sebastien6/UnityFullStackSample/blob/main/img/base-url.png)

- UIEntity.cs and Listing.cs were modified to accomodate the image upload.

When the application is run, the list of games is loaded and related images uploaded.

![unity app](https://github.com/sebastien6/UnityFullStackSample/blob/main/img/unityapp.png)

## Test

1. clone repository

```sh
git clone github.com/sebastien6/UnityFullStackSample
```
2. create a docker bridge network

```sh
docker network create -d bridge unity_app
```

3. go to directory app, build the docker image and launch the backend application
```sh
docker-compose build
docker-compose up
```

![backend service](https://github.com/sebastien6/UnityFullStackSample/blob/main/img/.png)

when the service is ready, the web URL for health check will return 'ok'

4. open Unity engine and open the unity app in folder unity of this repository

## Backend Web Service

To test the backend service open http://localhost:8080/healthz

![healthz](https://github.com/sebastien6/UnityFullStackSample/blob/main/img/healthz.png)

To access the openapiv2 UI, open http://localhost:8080/openapiv2

## Challenges

- first time coding in C# (understand language syntaxt, and concepts)
- first time scripting with unity (understand specificty of unity with C# like coroutine)
- deciding on the architecture (was with so many possibilities)
