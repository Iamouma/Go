package main

import "fmt"

type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}

type Document struct {
    content string
}

func (d Document) Read() string {
    return d.content
}

func (d *Document) Write(content string) {
    d.content = content
}

func main() {
    doc := &Document{content: "Initial content"}
    
    // using the Reader interface to read the document content
    var r Reader = doc
    fmt.Println("Content before writing:", r.Read())
    
    // using the Writer interface to write new content to the document
    var w Writer = doc
    w.Write("New content")
    
    // using the ReadWriter interface to read the updated document content
    var rw ReadWriter = doc
    fmt.Println("Content after writing:", rw.Read())
}
