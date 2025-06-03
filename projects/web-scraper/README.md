### Explanation of the Code

Importing Packages:

 * net/http is used to make HTTP requests.
  
 * goquery is used to parse the HTML and query the document in a jQuery-like syntax.
  
### Making an HTTP request:

 * The http.Get() method is used to send a GET request to the specified URL.
  
 * The response body is then passed to goquery.NewDocumentFromReader() to create a document object.
  
### Scraping the headlines

 * The doc.Find() method allows you to query the HTML document.

 * In this example, we search for article h2 tags (which represent headlines) and then iterate over them to print each title.