
## Getting Started

### Running Go
Enable dependency tracking for your code.
When your code imports packages contained in other modules, you manage those dependencies through your code's own module.
That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

To enable dependency tracking for your code by creating a go.mod file, run the go `mod init` command, giving it the name of the module your code will be in. The name is the module's module path.

In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path, see Managing dependencies.

For the purposes of this tutorial, just use example/hello.

```
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
else you can run

```
$ go run example/hello.go
```

### Testing
Within a directory only one package can exist, except a test package.

```
$ go test ./...  // scans through all the test packages and run within the directory

```

### Building the application

```
$ go build -o hello-application.exe .\main.go
$ ./hello-application  // to run the application

```
### Exercise
1. Write a package that provides a struct type Counter, with a method Next. The first time you call Next on a new Counter it should return 0,
    and if you call it again repeatedly it should return 1, 2, 3 and so on. Users should be able to set the current value of the counter to any value
    they want.
2. Add a Run method to your counter that loops repeatedly calling Next and printing the result to the configured output. As always, running the tests
    should not cause the counter to print anything in the terminal. How can you test the Run method without the test simply pausing forever?