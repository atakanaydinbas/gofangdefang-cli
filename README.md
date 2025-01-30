<p align="center">
  <img src="static/logo.png" width="250" title="hover text">
</p>

# GoFangDefang-CLI - Secure IOC Manipulation CLI Tool

GoFangDefan-CLI is a command-line interface (CLI) tool designed for the secure manipulation of Indicators of Compromise (IOCs). It facilitates the conversion between the original "fang" format (with special characters) and a safer "defang" format, preventing accidental execution of potentially malicious IOCs such as URLs, IPs, domains, or subdomains.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Defang](#defang)
  - [Defang File](#defang-file)
  - [Fang](#fang)
  - [Fang File](#fang-file)
- [Help](#help)

## Installation

To install gofangdefang-cli, follow these steps:

```bash
go install -v github.com/atakanaydinbas/gofangdefang-cli@latest
```

## Usage

### Defang

Defang an IOC by converting it to a safer format:

```bash
gofangdefang-cli defang <IOC>
```

### Defang File
Defang IOCs from a file:

``` bash
gofangdefang-cli defangfile -i <inputfilepath> -o <outputfilepath>
```

### Fang
Fang an IOC by converting it back to its original format:

``` bash
gofangdefang-cli fang <IOC>
```
### Fang File
Fang IOCs from a file:

``` bash
gofangdefang-cli fangfile -i <inputfilepath> -o <outputfilepath>
```

### Help
``` bash
gofangdefang-cli [command] --help
```
## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or create a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details
