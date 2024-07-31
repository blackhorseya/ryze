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
