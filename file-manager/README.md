#### Explanation:
* List Files in a Directory:

Uses ioutil.ReadDir to list all files in a specified directory.

* Copy a File:

Opens the source file, creates a destination file, and copies the content using io.Copy.

* Move a File:

Uses os.Rename to move a file. This renames or moves the file to a new location.

* Delete a File:

Uses os.Remove to delete a file.

* Search for Files:

Recursively searches for a file by name using filepath.Walk.

* Concurrent File Copy:

Demonstrates concurrency by copying multiple files concurrently using goroutines and a sync.WaitGroup.

* Running the File Manager:
Create some sample files to test the file manager, or use files already in your working directory.

Run your program with:

        go run main.go

Example Output:

        File Manager CLI Tool
        Files in .:
        file1.txt
        file2.txt
        Successfully copied file1.txt to ./backup/file1.txt
        Successfully copied file2.txt to ./backup/file2.txt
