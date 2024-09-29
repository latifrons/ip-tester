# IP and Headers Echo Server

This is a simple Go server that echoes back the client's IP address, remote IP address, and all HTTP headers received in the request.

## Features

- Displays the remote IP address of the client
- Shows the client IP address (considering X-Forwarded-For header for proxied requests)
- Lists all HTTP headers received in the request
- Responds with JSON format

## How to Run

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Navigate to the project directory.
4. Run the server:

   ```
   go run main.go
   ```

5. The server will start running on `http://localhost:8080`.
