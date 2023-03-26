# Golang_tutorial

### site: https://go.dev/doc/tutorial/

## memo

When importing a package contained in another module.<br>
``go mod init example.com/<dirname>``
<br><br>

Add new module requirements and sums.<br>
``go mod tidy``
<br><br>

Run code<br>
``go run .``
<br><br>

The command specifies that example.com/greetings should be replaced with ../greetings for the purpose of locating the dependency.<br>
``$ go mod edit -replace example.com/greetings=../greetings``
<br><br>

Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program.
`log.Fatal(err)`
<br><br>

☆: Create a messages map to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: make(map[key-type]value-type). You have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.

