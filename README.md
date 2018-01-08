# gocrawl
A proof-of-concept micro service based page crawler and mapper to demonstrate the technology components used, and how they would fit into an architectural pattern.  This is not what a production service would look like.

It is using docker-compose to build a linkerd service mesh and golang services .

## Overview

This is a demonstration of using linkerd in a docker-compose environment to route gRPC and HTTP requests.

The frontend is delivered through a simple golang delivered page, which talks to the crawling service thought gRPC calls.

The main characteristics of the service mesh are as follows:

- A router to proxy all the the requests to the right running services
- An admin service to introspect the performance of the service mesh
- A frontend, server-side, service to provide a user experience interface
- An API for the external clients to talk to (separation of concerns)
- services that handle requests (inbound initiated)
- services that request work (outbound initiated)

## Quick start

Get the repository of code

`git clone https://github.com/SonOfBytes/gocrawl.git`

Add vendor dependencies

`dep ensure`

Build the cli binary

`make`

Build the docker images

`docker-compose build`

Run the service

`docker-compose up`

Access the service web interface at [http://0.0.0.0:8080](http://0.0.0.0:8080)

Authenticate with

```
Username: someone
Password: hardcoded
```

And submit an url (e.g. http://example.com)

Submitting an URL through the web interface will trigger the start of a crawl through that site and it's linked pages that are in the same domain.

A spinner will be presented when the page is waiting for the backend to process a URL.

## Services

### linkerd

The [linkerd](https://linkerd.io/docs/) service is a mesh router that through transport introspection and name services routes inbound requests to the destination services.

We use HTTP for the frontend UI and unsecured gRPC for the backend microservices.

In this example we use docker name resolution through the filesystem name resolver to find the right container IP and port.

The HTTP services are listening on port 8080 and routes all traffic to frontend. [http://0.0.0.0:8080](http://0.0.0.0:8080)

The gRPC services are introspected to assess the Service name being requested and then routed via the linkerd filesystem resolver (see disco dircetory) to the container and port providing that gRPC service.

The linkerd admin page is available on port 8090. [http://0.0.0.0:8090](http://0.0.0.0:8090)

### frontend

`frontend` is a simple HTTP service that authenticates the user, provides a session key as a cookie and then presents the authorized user to the URL crawler interface.  Although this is a server side client, the same principles apply to browser side.

Once an URL has been submitted on the UI it is posted to the 'API' service for processing, and the 'Store' service via the 'API' service is polled for results.

Clicking through the links on the page will poll the 'Store' service for the results, and if not available yet present a spinner.

### api

`api` is a [gRPC](https://grpc.io/docs/) 'API' service that supports concurrent access and provides the frontend and cli with it's interface.

This service is stateless and can be horizontally scaled for resiliency and would allow for zero downtime deployments.

The API allows for 
- the implementation backend architecture to be abstracted
- acts as a control point for what services can be accessed from the client side (eg can only read from store, not write)
- some client independent recovery strategies can be implemented in a service outage (eg dead letter drop of requests or status page for the user)

The 'API' service allows controlled access to authentication, submission to the queue, and viewing store content.

### queue

`queue` is gRPC 'Queue' service that is a simple in memory list of URLs to be processed and a cache of URLs that have already been submitted or have already been processed recently.

This encapsulates the business logic of what is "valid" to be processed and in this case de-duplicates and determines when a URL can be submitted for reprocessing (i.e. every 5 minutes)

When requested for a URL it presents the first unprocessed URL in its queue and immediately puts it in the processed list with an expiration to prevent rapid re-requesting of the same URLs.

As implemented this is a stateful service that is not scalable but has a high processing throughput.  A more resilient queue service would be more appropriate in a real world production example (eg AWS SQS or rabbitMQ) if loss of requests was a concern.

There is also a race condition as implemented where two services can "get" the same URL before it is moved to the processed list.  However, this is an acceptable risk as any storage of the URL results will be de-duplicated in the 'Store' service. Apart from this logical race condition this service should be concurrent safe.

### retriever

`retriever` is a gRPC client that polls the 'Queue' service for URLs to process and find links in the same domain.  It then submits those links  to the 'Queue' service for further processing.

The 'Queue' service will act as an arbitrator of what is new and what has already been submitted or processed.  This allows the retriever to be scaled out horizontally to process sites in parallel.

An instance of the `retriever` process has multiple threads, however these have been throttled to 4 threads to clearly demonstrate how `retriever` can be scaled with more instances.

To launch 4 instances of `retriever`, and therefore 16 retrieval threads run

`docker-compose up --scale retriever=4`

`retriever` has the following controls
 
- service stays in the same top domain
- concurrency rate limited by the number of instances running
- caches URLs for 15 minutes to prevent too frequent revisits
- will only recurse 20 deep to limit loops and time spent crawling, although that is a very large potential tree

### store

`store` is a gRPC 'Store' service that supports concurrent access and provides a list of URLs for a given URL index.

As implemented this is a stateful service that is not scalable, however the separation of the store as a service allows the model abstraction from what could be a scalable solution (eg store shards, replicated key/value stores, etc)

## CLI

The CLI `bin/gocrawl` uses the client API interface as the frontend would.  It is there as a testing tool and to demonstrate the flow for commands.

For example (hypothetically),

```
$ bin/gocrawl authenticate someone hardcoded
wddJEWqaZkvh

$ bin/gocrawl submit wddJEWqaZkvh http://www.example.com
Job: tSobTNEifNnI

$ bin/gocrawl get wddJEWqaZkvh http://www.example.com
http://www.example.com/
http://www.example.com/about
http://www.example.com/help

$ bin/gocrawl get wddJEWqaZkvh http://www.example.com/about
http://www.example.com/
http://www.example.com/contact
http://www.example.com/help
```

The CLI does not have a polling feature, so a URL that has not been retrieved yet will return a Not Found error.

## Development notes

proto files are generated as follows (see `make proto`)

`protoc --go_out=plugins=grpc,:. <module>/pb/*.proto`

where module is the service component.

mocks are created with mockgen (see `make mock`)

`mockgen -destination=<module>/mocks/mock_<interface_ref>.go -package=mocks github.com/sonofbytes/gocrawl/<module>/server <interface>`

Tests are run with

`make test`

Test coverage report can be seen with

`make cover`