## My Speaking
- This project give you the general view how microservice work.
- There are many service mesh like Istio , HCP Consul Dedicated , Nomad and Kubernetes give full microservice's options like service discovery , registration , retry , circuit-breaker .....
- But diving into each edge of microservice problem will make you fully understanding about how microservice work. What is the functionality of each component in a microservice project

## Project Structure
Order Service — Create an order with APPROVAL_PENDING state. \
Consumer Service — Verify an user can order. \
Kitchen Service — Verify order and create a Ticket as CREATE_PENDING state.\
Accounting Service — Verify user's credit card.\
Kitchen Service — Change ticket's state to AWAITING_ACCEPTANCE.\
Order Service — Change order state to APPROVED.
------------------
Service Registration
API Gateway Service 
Discovery + Load Balancer Service
Orchestration Service

## Tech Stack
### Load Config
  + Viper
  + Remote key/value Configuration:  Storage Consul
### Dependency injection (Build time DI)
  + wire
### Service Communication
  + grpc 
  + Consider to clone https://github.com/protocolbuffers/protobuf.git
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
### circuit-breaker
  + hystrix-go

## Architect Pattern
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
   + Add path to environment
   + Start consul agent : consul agent -dev
   + Visit localhost:8500 for UI
   + Add key-value configuration to consul Key / Value if you set environment variable APP_ENV different with "local"
### fabio
   + Install fabio: https://github.com/fabiolb/fabio/releases
   + Add path to environment
   + Place fabio.properties and fabio executive file in the same folder
   + Start fabio : fabio.exe -cfg fabio.properties
### postgres
   + Run docker-compose: docker-compose up db nats

## Development
### Local
| **Service**  | **Port** | **Service Name** | **Database Address** | **Database Name** | **Database UserName/Password** |
|--------------|----------|------------------|----------------------|-------------------|--------------------------------|
| Gateway      | 8080     | gateway          | -------------        | -------------     | -------------                  |
| Order        | 50000    | order            | localhost:5433       | orderdb           | orderuser/orderpwd             |
| Accounting   | 50001    | accounting       | -------------        | -------------     | -------------                  |
| Consumer     | 50002    | consumer         | -------------        | -------------     | -------------                  |
| Kitchen      | 50003    | kitchen          | localhost:5433       | kitchendb         | kitchenuser/kitchenpwd         |
| Mail         | 50004    | mail             | -------------        | -------------     | -------------                  |
| Orchestrator | 50005    | orchestrator     | PENDING              | PENDING           | PENDING                        |
| Broker       | 4222     | -------------    | -------------        | -------------     | -------------                  |
| ConsulUI     | 8500     | -------------    | -------------        | -------------     | -------------                  |
| FabioUI      | 9998     | -------------    | -------------        | -------------     | -------------                  |
| FabioLB      | 9999     | -------------    | -------------        | -------------     | -------------                  |

