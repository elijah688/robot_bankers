# Robot Bankers

The `robot` package simulates a bank with two robot bankers who perform transactions. The `Spendy` robot always withdraws the maximum amount possible from the bank account, whereas the `Stingy` robot only withdraws a small amount. The robots run concurrently and alternate in performing transactions. The bank account balance should oscillate between 0 and 10.

## Usage

To use the `robot` package, simply import it and create a new set of `Robots`:

```go
import "github.com/your-username/robot"

func main() {
    var bankBalance int
    condVar := sync.NewCond(&sync.RWMutex{})
    robots := robot.NewRobots(condVar, &bankBalance)

    robots.Run()
}
```

Once you create a new set of `Robots`, you can start them by calling the `Run` method. The robots will perform transactions indefinitely until the program is terminated.

Here's how you could add the Makefile to the README.md:

## Usage

To build the program, run:

```
make build
```

To run the program, run:

```
make run
```

To run the tests, run:

```
make test
```

To remove the binary, run:

```
make clean
```

## Contributions

Contributions are welcome! If you find a bug or have an idea for a feature, please open an issue or submit a pull request.
