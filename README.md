
# Hexagonal Architecture Go Application

This repository houses a Go application structured on the principles of Hexagonal Architecture. The main goals of this project include file processing, size calculation, database updating using a decentralized consensus mechanism (RAFT), and comprehensive logging, tracing, debugging, and metrics of the API flow facilitated by the go-hclog logging mechanism.

## About Hexagonal Architecture

Hexagonal Architecture, also known as Ports and Adapters Architecture or The Onion Architecture, is a design pattern that emphasizes the separation of concerns and the independence of components in a software system. It provides a way to structure an application to make it more modular, maintainable, and adaptable to changes. The key idea is to organize the application into layers, or "hexagons," with each layer having a specific responsibility and interacting with the others through well-defined interfaces or ports.

### The main components of Hexagonal Architecture are:

- Core Business Logic:

```
This is the heart of the application, containing the business rules and domain-specific logic.
It should be independent of external concerns, such as databases, frameworks, or user interfaces.

```
- Ports:
```
Ports define interfaces through which the core logic communicates with the external world.
They act as entry and exit points for the application, providing a clear separation between the core and external dependencies.
Ports are typically defined as interfaces or abstract classes.
```

- Adapters:
```
Adapters are implementations of the ports. They bridge the gap between the core logic and the external dependencies.
Each external dependency (database, UI, third-party services) has its adapter, providing a specific implementation of the port interface.
Adapters are responsible for translating the core logic's requests into operations that the external dependencies can understand.
```

- External Dependencies:
```
These are components or services that the application interacts with, such as databases, web frameworks, external APIs, and more.
External dependencies are accessed through the adapters, ensuring that the core logic remains independent and testable.
```

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
