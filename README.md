# Load Balancer in Go

This project implements a simple HTTP load balancer in Go. The load balancer distributes incoming HTTP requests to a pool of backend servers using a round-robin algorithm.

## Features

- Round-robin load balancing
- Health check for backend servers (basic implementation)
- Simple and extensible design

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/haikali3/load-balancer-go.git
    cd load-balancer-go
    ```

2. Create three simple backend servers that will be used by the load balancer. You can use the following code for each server, changing the port number for each one:

    ```go
    package main

    import (
        "fmt"
        "net/http"
    )

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from server on port %s!", r.Host)
    }

    func main() {
        http.HandleFunc("/", handler)
        fmt.Println("Server is listening on port 3000")
        http.ListenAndServe(":3000", nil)
    }
    ```

    Save this file as `server3000.go` for the first server, then create `server3001.go` and `server3002.go` for the other two servers by changing the port number to 3001 and 3002, respectively.

3. Run each backend server in a separate terminal window:

    ```sh
    go run server3000.go
    go run server3001.go
    go run server3002.go
    ```

4. In the main project directory, run the load balancer:

    ```sh
    go run main.go
    ```

    The load balancer will listen on port 8080 and forward requests to the backend servers running on ports 3000, 3001, and 3002.

## Usage

Once the load balancer and backend servers are running, you can test the setup by making HTTP requests to the load balancer:

```sh
curl http://localhost:8080
