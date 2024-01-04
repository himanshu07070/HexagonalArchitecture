
# Hexagonal Architecture Go Application

This repository houses a Go application structured on the principles of Hexagonal Architecture. The main goals of this project include file processing, size calculation, database updating using a decentralized consensus mechanism (RAFT), and comprehensive logging, tracing, debugging, and metrics of the API flow facilitated by the go-hclog logging mechanism.

## Project Structure

The application is meticulously organized to embody the principles of Hexagonal Architecture, effectively segregating business logic from external dependencies. Here's a glimpse of the project structure:

```
├── app
│   |
│   └── main.go
├── internal
├   |── api
│   │   ├── handler.go
│   │   
│   ├── core
│   │   ├── fileprocessor.go
│   │   |
│   │   └── fileprocessor_test.go
├   |── consensus
│   │   └── raft_adapter.go
│   ├── database
│   │   └── redis_adapter.go
│   └── logger
│   |   └── hclog_adapter.go
│   └── ports.go
├── go.mod
├── go.sum
└── README.md

```
## Usage

1: Clone the repository:
- git clone https://github.com/himanshu07070/HexagonalArchitecture.git

2: Navigate to the project directory:
- cd hexagonal-architecture-go

3: Build and run the application:
- go run app/main.go (1st terminal- start the backend server)
- redis-server (2nd terminal- start the redisdatabase server)

4: Open Postman( To request api)
- Mention this in the url - http://localhost:8080/upload
- Select the POST method
- Select Body
- Select form-data
  - key: file
  - type: file
  - value: upload any file from local system
  - Hit the send button to request the api
