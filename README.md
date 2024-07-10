
# Brackets

This program receives an undefined number of string in arguments. For each argument, if the expression is correctly bracketed, the program prints on the standard output OK. Otherwise it prints Error. Symbols considered as brackets are parentheses ( and ), square brackets [ and ] and curly braces { and }. Every other symbols are simply ignored.

An opening bracket must always be closed by the good closing bracket in the correct order. A string which does not contain any bracket is considered as a correctly bracketed string.

## Documentation

This part contains instructions on how to make use of the program.

## Features

- Checks if a string is correctly bracketed.


### Usage

To use this program, download and install the latest version of Go from [here](https://go.dev/doc/install).

After cloning the projects repository, navigate to the RepeatAlpha directory and execute the following command in your terminal.
```bash
go run . <string>

```
For example:
```bash
$ go run . '(johndoe)' | cat -e


```
Expected output :
```bash
OK$

```

## Contributions
Users of this program are allowed to contribute to this project in terms of adding features, or fixing bugs. Just make a pull request.

## Author

- [@JosephOkumu](https://github.com/JosephOkumu)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Credits

[Zone01Kisumu](https://zone01kisumu.ke)