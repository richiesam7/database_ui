# database_ui
A user interface to manage and view all your database table info: with an E-commerce view and Customer care view.

## About - Unique selling proposition (USP)
We always tend to stack our products into various components - front end stack, back end or server stack, database stack, and all the other stacks beyond and between. A big challenge in software engineering now is about _managing definitions for all the data that we want to store and visualize_. 
To store a data object,
- We end up defining a database schema (DB definition of the object),
- writing up API's to serve this schema and comply to it (server side definition of the object), 
- and a web based client which also has its own definition (client side definition of the object).

In my opinion, _this is absurd as it is error prone_. 

Imagine your DB schema storing an integer as the primary key, and you miss out on adding a client side validation in your web app to accept only integers, and the API allowing free text? What do you get - Yes, you are right, Garbage! 

Imagine, **your database schema definition doing all of this for you**. This product is build on the sole premise of _single-source-of-truth for data object definitions_, which is the point where it is stored, which in most cases in the industry is a _database_. 
The Go server and Angular client will all depend on this data object definition to create run time objects for the data that is being rendered on the client or passed along via the server. Everything is dynamic and contructed on the fly based on data object/table schema definition fetches. The schema definition is valid across the current user session, and hence is not be altered while user sessions are in progress (this can have unintended consequences).

## Second USP
Every single time, you need to bootstrap a product, you end up writing a server, a client, database, and all the associated boiler plate. Imagine all of that taken care for you by one product. And as a plus, you get a _e-commerce and customer care_ type of UI view for your data. 
This is what this product _database_ui_ gives you.

# Architecture and Components

A Postgre database, served by a Highly available Go server, sending/receiving data from an Angular-based web client is the essense of this product.

The below diagram explains the high level architecture of the product.
![image](https://user-images.githubusercontent.com/49153293/111121145-4be5c180-8592-11eb-902b-3503d5a5dcac.png)

## Server architecture and stack
- Microservices architecture
- Each service deployed on separate containers
- Postgres on Docker, with scalable config
- Managed by Kubernetes
- gRPC for communication across services
- Server accepts http incoming requests
- Redis as cache - CacheService as a microservice
- Redis server deployed as separate container

## Client architecture and stack
TBA

# Technology (v1)
- Database - Postgre
- Server - Golang
- Processing - Redis, Microservices, gRPC
- Client - Angular
- Infrastructure - Docker, Kubernetes
## Subsequent versions
- Databases - Postgre, _Oracle, My SQL_
- Server - Golang, _Kotlin_
- Processing - Redis, Microservices, gRPC, _Kafka, AWS services_
- Client - Angular, _React, Vue.js_
- Infrastructure - Docker, Kubernetes

## Releases
Expected release of v1 - August 2021
### What to expect
make install - Gets anll dependencies and packages
- make start - Start containers and microservices required for the app
- make demo - sets up demo project with dummy data
- Access url from browser. For example this url will show the details of a mobile phones from dummy data - http://localhost:1000/mobiles/111

Version 2 - Dec 2021
Reach out to me if you have any requirements or queries.
