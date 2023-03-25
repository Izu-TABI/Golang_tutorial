# Golang_tutorial

### site: https://go.dev/doc/tutorial/

## memo

When importing a package contained in another module.<br>
``go mod init example.com/<filename>``
<br><br>

Add new module requirements and sums.<br>
``go mod tidy``
<br><br>

Run code<br>
``go run .``

The command specifies that example.com/greetings should be replaced with ../greetings for the purpose of locating the dependency.<br>
``$ go mod edit -replace example.com/greetings=../greetings``


