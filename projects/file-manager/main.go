package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sync"
)

// Navigate directories and list files
func listFilesInDir(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Files in %s:\n", dir)
    for _, file := range files {
        fmt.Println(file.Name())
    }
}

// Copy a file
func copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return err
    }

    return nil
}

// Move a file (renaming it)
func moveFile(src, dst string) error {
    return os.Rename(src, dst)
}

// Delete a file
func deleteFile(filepath string) error {
    return os.Remove(filepath)
}

// Search for files by name (recursively)
func searchFiles(dir, searchName string) {
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.Name() == searchName {
            fmt.Printf("Found: %s\n", path)
        }
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}

// Concurrently copy multiple files
func concurrentCopyFiles(files []string, destDir string) {
    var wg sync.WaitGroup
    wg.Add(len(files))

    for _, file := range files {
        go func(file string) {
            defer wg.Done()
            dst := filepath.Join(destDir, filepath.Base(file))
            err := copyFile(file, dst)
            if err != nil {
                log.Printf("Failed to copy %s: %v", file, err)
            } else {
                fmt.Printf("Successfully copied %s to %s\n", file, dst)
            }
        }(file)
    }

    wg.Wait()
}

func main() {
    fmt.Println("File Manager CLI Tool")

    // Example usage:
    // List files in the current directory
    listFilesInDir(".")

    // Copy a file
    err := copyFile("source.txt", "copy_of_source.txt")
    if err != nil {
        log.Printf("Failed to copy file: %v", err)
    }

    // Move a file
    err = moveFile("copy_of_source.txt", "moved_source.txt")
    if err != nil {
        log.Printf("Failed to move file: %v", err)
    }

    // Delete a file
    err = deleteFile("moved_source.txt")
    if err != nil {
        log.Printf("Failed to delete file: %v", err)
    }

    // Search for a file by name
    searchFiles(".", "somefile.txt")

    // Concurrently copy multiple files
    filesToCopy := []string{"file1.txt", "file2.txt"}
    concurrentCopyFiles(filesToCopy, "./backup")
}
