## A simple web server in Go

#### This code does the following:

 * Creates a simple Go web server using the net/http package.
 * Sets up a route (/hello) that responds with a JSON object containing the message "Hello, World!".
 * Uses json.NewEncoder().Encode() to write the response in JSON format.
 * 
#### Run the Server
To run the server, execute the following command in your terminal:

        go run main.go

You should see this log message:

        Server is running on http://localhost:8080

#### Test the API

Open your browser or use a tool like curl or Postman to test the API by visiting:

        http://localhost:8080/hello

You should receive the following JSON response:

        {
        "message": "Hello, World!"
        }

Alternatively, you can test with curl in your terminal:


        curl http://localhost:8080/hello

This should output:

        {"message":"Hello, World!"}

#### Project Summary
You've created a basic web server in Go.
You set up an endpoint that responds with a JSON message "Hello, World!" when visited.