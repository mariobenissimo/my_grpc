## Client-Server Communication Project with Proto
This project demonstrates an implementation of client-server communication using gRPC (Google Remote Procedure Call) and the Proto protocol.

## Description
The project utilizes gRPC to facilitate communication between a client and a server using an interface defined in the .proto file. The Proto file defines the services and messages used for communication. This allows for efficient and typed communication between the client and the server.

The project is divided into two main components:

Client: The client sends requests to the server and receives responses. It uses the message and service definitions from the Proto file to automatically generate the necessary code for communication.
Server: The server receives requests from the client, processes the requests, and sends back responses using the message and service definitions from the Proto file.

## Dependencies
The project requires the following components to be installed:

gRPC: A framework for remote communication between applications.
Proto Compiler: A compiler to generate code from the .proto file.

## Usage Instructions

1. Generate code from the Proto file:

``` protoc --go_out=. routeguide.proto ```

2. Start the server:

``` cd /server ```

``` go run . ```

3. Start the client:

``` cd /client ```

``` go run . ```
