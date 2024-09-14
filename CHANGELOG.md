## 0.4.1 (2024-09-14)

### Feat

- refactor transaction amount handling

### Refactor

- refactor transaction creation and field handling
- refactor codebase to remove unused imports and fields

## 0.4.0 (2024-09-13)

### Feat

- improve handling of new block functionality
- refactor block service functionality
- enhance logging across multiple files
- enhance block scanning functionality (#60)
- implement new platform service command
- enhance block scanning functionality
- implement new platform service command
- refactor transaction handling and metadata management
- refactor transaction handling and RPC methods
- enhance transaction management functionalities (#59)
- enhance transaction management functionalities
- refactor timestamp serialization operations
- implement `pgx` storage option across modules (#58)
- implement `pgx` storage option across modules
- refactor database interactions for transactions
- update default constants and limits for better performance
- refactor transaction handling with new interfaces (#57)
- refactor transaction handling with new interfaces
- refactor block service functionality (#56)
- refactor block service functionality
- enhance BlockService functionality (#55)
- enhance BlockService functionality
- refactor transaction handling in `TransactionService` (#54)
- refactor transaction handling in `TransactionService`
- consolidate transaction service functions into new file (#53)
- consolidate transaction service functions into new file
- refactor server creation and code comments
- improve error handling and logging functionality
- refactor RPC methods for handling block transactions (#51)
- refactor RPC methods for handling block transactions
- refactor blockscanner to use EventBus dependency
- refactor block handling and logging operations (#50)
- refactor block handling and logging operations
- revamp transaction handling across multiple packages (#49)
- revamp transaction handling across multiple packages
- integrate `eventx` package across directories
- implement event bus with subscription and publishing capabilities
- implement the DomainEvent interface in NewBlockEvent (#48)
- implement the DomainEvent interface in NewBlockEvent
- consolidate changes in file structure for new feature
- refactor error handling for improved scalability
- refactor transaction handling logic
- consolidate scanner service commands and options
- enhance logging functionality for newBlock object
- refactor codebase to improve logging and collection handling (#47)
- refactor codebase to improve logging and collection handling
- implement logging and tracing for block operations (#46)
- implement logging and tracing for block operations
- refactor protobuf definitions and RPC methods
- refactor block scanning logic and logging initialization
- refactor server handling in scan package (#45)
- improve error handling and logging in server implementation
- refactor server handling in scan package
- refactor block scanning service initialization (#44)
- refactor block scanning service initialization
- consolidate new files and modify `contextx.go` in `scan` and `pkg/contextx` packages
- update API initialization and usage for improved performance
- update service names in configuration files
- update database configurations and services
- configure Buf build settings for v2 modules (#43)
- configure buf.build code generation plugins
- configure Buf build settings for v2 modules
- refactor interfaces and imports
- refactor persistence package with MongoDB repository
- refactor account balance handling and data types (#38)
- refactor account balance handling and data types
- update gRPC metadata handling in transaction processing

### Fix

- refactor error handling in transaction repository
- update OpenTelemetry dependencies to v1.28.0
- refactor configuration files and improve code consistency

### Refactor

- refactor context handling and error management in transactions
- refactor block service functions and improve readability
- refactor code for `FoundNewBlock` implementation (#61)
- refactor code for `FoundNewBlock` implementation
- refactor data processing logic
- refactor transaction handling and logging across files
- refactor comparison logic and conditions
- refactor logging and block handling in various files
- refactor `ScanBlockRequest` struct and related functions
- refactor transaction request naming conventions
- refactor codebase to remove unnecessary files and functionality
- refactor code logic for improved efficiency
- refactor transaction handling in process functions
- refactor BlockServiceServer and block.proto
- refactor transaction services across files
- refactor transaction service initialization and naming
- refactor function signatures for consistency
- refactor block streaming functionality across files (#52)
- refactor block streaming functionality across files
- update service clients and injector initialization across files
- refactor package structure and file naming conventions
- refactor file paths and update imports for better organization
- refactor project structure for `cmd` package
- refactor package names and imports across files
- refactor package names and import paths across project
- refactor initialization and error handling in block repository
- refactor global tracer initialization and usage
- refactor service initialization logic in wire files
- refactor package names for `scanner` consistency
- clarify function purpose and update comments
- refactor function and variable names across project
- refactor error handling across multiple modules
- refactor initialization and context handling
- refactor file naming conventions across project
- refactor server dependencies and imports
- refactor HTTP server setup and middleware handling
- refactor context handling and imports in various files
- refactor block service API integration
- refactor function names and references across files
- refactor code to remove `FetchAndStoreBlock` references
- refactor protobuf generation in build process
- refactor codebase to use `biz` instead of `blockB` types
- refactor method signatures and imports in MongoDB storage (#42)
- refactor method signatures and imports in MongoDB storage
- refactor directory structure for cleaner organization (#41)
- refactor package structure and imports for block repository
- update import paths and package names to use `mongodbx` (#40)
- refactor import paths and naming conventions
- refactor package import paths in gRPC files
- update import paths and package names to use `mongodbx`
- refactor service server initializations to use packages
- refactor Restful interface and related changes
- refactor directory structure for cleaner organization
- refactor import paths and function names across files
- refactor package structure and wire files
- update address handling in account service (#39)
- refactor error handling and address parsing in `GetAccount` method
- update address handling in account service
- refactor account-related protobuf and gRPC methods
- refactor transaction service interface and messages

### Perf

- improve performance with asynchronous block scanning

## 0.3.0 (2024-08-13)

### Feat

- create GORM implementation for block repository
- refactor PostgreSQL client initialization in app infra (#37)
- refactor service configurations and update ports
- refactor PostgreSQL client initialization in app infra
- implement detailed logging for transactions in service code (#36)
- implement detailed logging for transactions in service code
- refactor transaction handling in service layer
- improve transaction handling and imports
- integrate transaction service server in gRPC adapter
- introduce transaction service to biz package in transaction domain
- refactor transaction request message structure
- implement gRPC server for block and network services
- implement gRPC adapter with Wire dependency injection (#35)
- implement gRPC adapter with Wire dependency injection
- update `GRPC` struct with `URL` field usage
- implement OpenTelemetry instrumentation for gRPC server
- implement new function for retrieving context data (#34)
- implement new function for retrieving context data
- implement network server functionality in grpc and wire implementations (#33)
- implement network server functionality in grpc and wire implementations
- create network service struct and methods in domain biz directory
- register health, reflection, and grpc servers
- implement gRPC server interceptors
- create transaction service and proto messages
- define network related messages and services
- define message structures for block related requests and responses

### Fix

- improve error handling for gRPC connection settings
- refactor error handling and messages across multiple files

### Refactor

- update package imports and remove unnecessary file
- refactor block service to use GORM instead of MongoDB
- enhance NewTransactionFromTon function and fields
- refactor logging and error handling in `impl` struct
- refactor network service initialization across files
- refactor dependency injection for block service integration
- refactor server middleware and implement new functions (#32)
- refactor server middleware and implement new functions
- refactor import paths and mock initialization in test files
- update import paths and function signatures in block service
- restructure account domain files
- refactor file structure and build process for protobuf files

## 0.2.0 (2024-07-31)

### Feat

- implement List method for block entities (#30)
- refactor MongoDB operations in block package
- implement List method for block entities
- refactor error handling and data fetching logic (#29)
- refactor error handling and data fetching logic
- improve error handling and logging in block creation (#28)
- improve error handling and logging in block creation
- refactor MongoDB integration in application (#27)
- implement MongoDB integration for data retrieval and storage
- refactor MongoDB integration in application
- improve MongoDB testing and update container version
- refactor MongoDB implementation and methods
- refactor block repository files in domain directory
- integrate MongoDB storage functionality (#25)
- integrate MongoDB storage functionality
- refactor block fetching logic in FetchAndStoreBlock function
- implement FetchAndStoreBlock functionality across files (#24)
- implement FetchAndStoreBlock functionality across files
- consolidate grpc client implementation

### Fix

- implement retry logic and request parameters logging in API client
- refactor MongoDB query methods

### Refactor

- refactor block service function signatures
- remove retry limit parameter from API client initialization
- update logging messages for block scanning and service logs
- refactor data manipulation methods across files
- pass `block` variable as pointer in `ctx.Info` function
- improve block processing efficiency (#23)
- improve block processing efficiency
- refactor `restful` struct to `scan` struct (#22)
- refactor `restful` struct to `scan` struct
- refactor service initialization to use dependency injection
- refactor service initialization and context handling (#21)
- refactor service initialization and context handling
- update function calls for improved code readability
- update service name and port for block-grpc service

## 0.1.1 (2024-07-29)

## 0.1.0 (2024-07-29)

### Feat

- refactor block data structure and methods
- refactor block data retrieval and processing
- refactor gRPC implementation in block package (#20)
- refactor gRPC implementation in block package
- refactor server initialization process (#19)
- refactor server initialization process
- improve server logging and error handling
- create new `Server` type and `NewServer` function
- create new gRPC server for improved communication
- refactor application to support gRPC configuration (#18)
- refactor application to support gRPC configuration
- introduce new gRPC adapter implementation files
- consolidate start command functionality into new file
- update `ScanBlock` function with new API calls (#16)
- update `ScanBlock` function with new API calls
- integrate Tonx package and client into project
- implement `ScanBlock` function in `block_service.go`
- update protobuf definitions for block and transaction entities
- implement GetBlock method with timestamp initialization
- create new block service with GetBlock and GetBlocks methods
- update Swagger documentation for block scanning API
- refactor scan package dependencies (#15)
- update Block Scan API documentation and Swagger files
- enhance API documentation and health check endpoint
- refactor scan package dependencies
- refactor block scanning process and data types (#12)
- refactor block scanning process and data types
- refactor Client struct initialization and usage across codebase
- refactor network selection and scanning logic
- add blockchain scanning functionality (#11)
- add blockchain scanning functionality
- improve application logging and output messages
- add Testnet field to Network struct (#10)
- refactor configuration handling in stats command
- refactor configuration handling and file reading
- add Testnet field to Network struct
- create new structs for application and network configurations
- create new stats command under cmd/get/ (#8)
- refactor codebase to improve code readability
- create new stats command under cmd/get/
- implement new 'get' command functionality
- create interfaces and implementations for new services and commands
- consolidate error handling and API responses into new packages
- consolidate network-related functions into new file
- add account model (#5)
- add account model
- add network model (#6)
- add network model
- generate more test cases and protobufs efficiently
- gen-pb
- consolidate project structure and dependencies
- enhance Block message comments and add new messages (#3)
- enhance Block message comments and add new messages
- enhance Transaction Message Structure
- refactor protobuf message structure (#1)
- update protobuf message definitions
- refactor protobuf message structure
- consolidate and refactor proto file structures

### Refactor

- improve block sending logic and error handling
- update data type of block ID field throughout the application
- refactor `GetBlock` method for improved performance
- improve error handling and logging in server
- update grpc middleware package references
- refactor variable names and remove unnecessary code
- update return types to adapterx.Service in all functions
- update block height fields to uint32 data type (#17)
- update block height fields to uint32 data type
- consolidate command flags and remove duplicate function definition
- update API client naming conventions across files
- refactor network configuration handling
- refactor configuration handling across files
- refactor network configuration in `app/infra/configx`
- refactor code structure for improved organization
- standardize file naming conventions across domains
- update data types to bytes in Block and Transaction messages (#2)
- update data types to bytes in Block and Transaction messages
