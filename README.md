# go-locator

Go-locator is a Go package designed to get offline information about any IP address and save it to a file!

## Usage

```bash
    go-locator.exe [--ip IP] [-f FILE] [--output -o OUTPUT]
```
- --ip: For single IP
- -f: For multiple IPs in a file
- -o: For saving in output file

## Demo

```bash
    go run . --ip 8.8.8.8          
    Processing...
    Country name: United States
    ISO country code: US
    Time zone: America/Chicago
    Coordinates: 37.751, -97.822
    Accuracy Radius: 1000
    Metro Code: 0
    Is Anonymous Proxy: false
    Is Satellite Provider: false
    Is In European Union: false
    Finished in: 392.4106ms
```

## License

This project is licensed under the [MIT License](https://github.com/go-nerds/go-locator/blob/main/LICENSE). See the [LICENSE](https://github.com/go-nerds/go-locator/blob/main/LICENSE) file for details.
