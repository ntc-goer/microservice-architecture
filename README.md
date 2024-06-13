Order Service — Tạo một Order ở trạng thái APPROVAL_PENDING. \
Consumer Service — Xác minh rằng người tiêu dùng có thể đặt đơn hàng. \
Kitchen Service — Xác thực chi tiết đơn hàng và tạo một Ticket ở trạng thái CREATE_PENDING.\
Accounting Service — Phê duyệt thẻ tín dụng của người tiêu dùng.\
Kitchen Service — Thay đổi trạng thái của Ticket thành AWAITING_ACCEPTANCE.\
Order Service — Thay đổi trạng thái của Order thành APPROVED.

## Tech Stack
### Load Config
  + Viper
  + Remote key/value Storage Consul
### Dependency injection 
  + wire
### Service Communication
  + grpc 
  + Consider to git clone https://github.com/protocolbuffers/protobuf.git
  + nats
### Gateway
  + grpc-gateway
### Service Registration/
  + consul
### Database
  + Postgres
  + ent ORM -> Refer https://entgo.io/
### Server Load Balancing / Discovery
  + fabio
### Resty
### saga
### circuit-breaker

## Prerequisite
### consul
   + Installing consul: https://developer.hashicorp.com/consul/docs/install#precompiled-binaries
   + Start consul agent : consul agent -dev
   + Visit localhost:8500 for UI
### fabio
   + go install github.com/fabiolb/fabio@latest
### postgres
   + Run docker-compose: docker-compose up db db_admin
