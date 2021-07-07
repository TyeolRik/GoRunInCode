# GoRunInCode
This Code is developed in Ubuntu 20.04

## How to use

1. Change Directory to ```mainDirectory```

```sh
$ cd mainDirectory
```

2. Initialize Your terminal.

```sh
$ go mod init && go mod tidy && go mod vendor
```

3. Run this

```sh
$ go run .
```

## Result

You would get ```output.txt``` file in ```GoRunInCode/mainDirectory```. You would understand why this output returns by looking ```GoRunInCode/anotherDirectory/main.go```