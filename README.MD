# Simple TCP Server and Client in Go

This is a simple TCP server written in Go that listens on a specified address and port, accepts incoming connections, and provides a basic command handling mechanism.

## Server setup
1. Clone the repository:
   ```bash
   git clone https://github.com/comr8/time-server.git
   ```
2. Build the server binary:
   ```bash
   go build server.go
   ```
3. Run the server:
   ```bash
   ./server
   ```
## Client setup
1. Clone the repository if it was not cloned in the previous step.
   ```bash
   git clone https://github.com/comr8/time-server.git
   ```
2. Build the client binary:
   ```bash
   go build client.go
   ```
3. Run the server:
   ```bash
   ./client
   ```
## Usage
The server listens for incoming connections and handles two commands:
- `time`: Returns the current server time.
- `exit`: Closes the connection.

To interact with the server, you can use tools like `telnet`, `PuTTY`, or create a custom client to send commands. If you prefer, you can also use the client application that I have developed, located in the `Client` directory.