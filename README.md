# String art
## Overview

String art displays the art of printable ascii charcters using various font styles. It offers both terminal and file output options.

## Features
- Choose from different font styles.
- Output directly to the terminal.
- Save results to a specified file.

# Getting Started
## Prerequisites

Ensure you have Go installed on your machine. If not, you can download it [here](https://go.dev/dl/).

## Clone the Repository

```
git clone https://github.com/nguonodave/str_art.git
```

## Navigate to the cloned repository and run the program (basic usage)

```
cd str_art
go run . "hello"
```

# Usage customization
## Choose Font Style

The following are the current implemented font styles:
- Standard
- Shadow
- Thinkertoy

**NOTE:** Not specifying the banner font to be used will use the `standard` font by default.

## Displaying on the terminal

The following is the allowed usage:

>`go run . [STRING] [OPTIONAL BANNER]`
>
>Examples:
>- `go run . hello shadow` will display:
>   ```
>   _|                _| _|          
>   _|_|_|     _|_|   _| _|   _|_|   
>   _|    _| _|_|_|_| _| _| _|    _| 
>   _|    _| _|       _| _| _|    _| 
>   _|    _|   _|_|_| _| _|   _|_|   
>   ```
>- `go run . hello thinkertoy` will display:
>   ```
>   o        o o     
>   |        | |     
>   O--o o-o | | o-o 
>   |  | |-' | | | | 
>   o  o o-o o o o-o 
>   ```
>- `go run . hello` or `go run . hello standard` will display:
>   ```
>    _              _   _
>   | |            | | | |         
>   | |__     ___  | | | |   ___   
>   |  _ \   / _ \ | | | |  / _ \  
>   | | | | |  __/ | | | | | (_) | 
>   |_| |_|  \___| |_| |_|  \___/  
>   ```

## Writing to a file

The following is the allowed usage:

>`go run . [OUTPUT FLAG] [STRING] [OPTIONAL BANNER]`
>
>Examples:
>- `go run . --output=hello.txt hello shadow` will display the shadow art in the hello.txt file.
>
>- `go run . --output=test.txt nice thinkertoy` will display the thinkertoy art in the test.txt file.
>
>- `go run . --output=wow.txt great` or `go run . --output=wow.txt great standard` will display the standard art in the wow.txt file.

# Credits

This project was inspired by the [01 Talent](https://01talent.com/) pedagogy at [Zone01 Kisumu](https://www.zone01kisumu.ke/) campus in Kenya.

# Contribution

Contributions are welcome! Fork the repository and enhance the project with new fonts, features, or color schemes.

# License

MIT License. See [LICENSE](LICENSE) file for details.
