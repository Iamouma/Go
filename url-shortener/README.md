### Explanation of the Code

 * generateShortCode(): Generates a random 6-character string that serves as the short code.

 * createShortURL(): Accepts a long URL in a POST request, generates a short code, stores the long URL and the short code in the database, and returns the short URL.

 * redirectToLongURL(): Redirects to the long URL when the user accesses the short URL by querying the database using the short code.

 * initDB(): Initializes the SQLite database and creates the urls table if it doesn't already exist.

### Run the URL Shortener

To run the project, use the following command:

        go run main.go

This will start the server on http://localhost:8080.

### Test the URL Shortener

Create a Short URL: You can create a short URL by sending a POST request to /shorten with a JSON body like:

        {
        "long_url": "https://www.example.com"
        }

Example using curl:

        curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://www.example.com"}' http://localhost:8080/shorten


The response will look like:

        {
        "long_url": "https://www.example.com",
        "short_url": "http://localhost:8080/abc123"
        }

Redirect to Long URL: You can now visit http://localhost:8080/abc123 in your browser, and it will redirect you to https://www.example.com.