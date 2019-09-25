# stop

Stop adds functions for graceful app shutdown.

Basic usage is looks like:

```go
import "github.com/ninedraft/stop"

func main() {
    defer stop.Exit()

    // ... 
    defer cleanStuff()
    // do stuff
    // ...

    if err := doImportantThing(); err != nil {
        stop.Stop(100)
    }
}
```

In this scenario all defers will be called **before** your program exits with status 100.

## Why I should use it?
Compare with bare `os.Exit`.

```go
import "os"

func main() {    
    // ... 
    defer cleanStuff() // won't be called if bad stuff happens!
    // do stuff
    // ...

    if err := doImportantThing(); err != nil {
        os.Exit(100)
    }
}

```

Mostly it's important when you are writing to files and handling errors.

## I don't want to import another leftpad dependency!

Then just copy my code to your project :3 
It's really just a little more than a usefull snippet, which I have been using for many years.