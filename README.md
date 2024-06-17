## Comment
- This project give you the specific view how microservice work.
- There are many service mesh library like istio , kubernetes give full microservice's options like service discovery , registration , retry , circuit-breaker .....
- But diving into each edge of microservice problem will make you fully understanding about how microservice work. What is the functionality of each component in a microservice project

## Project Structue
Order Service — Create an order with APPROVAL_PENDING state. \
Consumer Service — Verify an user can order. \
Kitchen Service — Verify order và create a Ticket as CREATE_PENDING state.\
Accounting Service — Verify user's credit card.\
Kitchen Service — Change ticket's state to AWAITING_ACCEPTANCE.\
Order Service — Change order state to APPROVED.

## Tech Stack
### Load Config
  + Viper
  + Remote key/value Configuation:  Storage Consul
### Dependency injection (Build time DI)
  + wire
### Service Communication
  + grpc 
  + Consider to git clone https://github.com/protocolbuffers/protobuf.git
  + nats
### Gateway
  + grpc-gateway
### Service Registration
  + consul
### Database
  + Postgres
  + ent ORM -> Refer https://entgo.io/
### Server Load Balancing / Discovery
  + fabio
### Manage Log , Trace , Metric
  + OpenTelemetry
### Resty
### circuit-breaker
  + hystrix-go

## Design Pattern
- Database per service pattern. Refer https://microservices.io/patterns/data/database-per-service.html
- Remote procedure invocation pattern. Refer http://microservices.io/patterns/communication-style/messaging.html
- Circuit breaker pattern. Refer http://microservices.io/patterns/reliability/circuit-breaker.html
- Asynchronous messaging pattern. Refer http://microservices.io/patterns/communication-style/messaging.html.
- Saga pattern. Refer http://microservices.io/patterns/data/saga.html.
- API Gateway. Refer https://microservices.io/patterns/apigateway.html
- Server-side service discovery. https://microservices.io/patterns/server-side-discovery.html

## Prerequisite
### consul
   + Install consul: https://developer.hashicorp.com/consul/docs/install#precompiled-binaries
   + Start consul agent : consul agent -dev
   + Visit localhost:8500 for UI
### fabio
   + Install fabio: https://github.com/fabiolb/fabio/releases
   + Place fabio.properties and fabio executive file in the same folder
   + Start fabio : fabio.exe -cfg fabio.properties
### postgres
   + Run docker-compose: docker-compose up db db_admin
