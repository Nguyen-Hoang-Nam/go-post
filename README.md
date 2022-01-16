# Go Post

Test api like Curl and Httpie

## Usage

```text
Test api with method, header and body.

Usage:
    go-post [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
    delete      Test delete method
    get         Test get method
    help        Help about any command
    post        Test post method
    put         Test put method

Flags:
        --header strings   A list of headers
    -h, --help             help for go-post
        --param strings    A list of params
        --sort-header      Sort header by alphabet
    -t, --toggle           Help message for toggle

Use "go-post [command] --help" for more information about a command.
```

## TODO

- Improve performance of encode, decode Json [here](https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go/47798475#47798475)
  and [here](https://github.com/TylerBrock/colorjson)
- Support WSS
- Support Grpc

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
