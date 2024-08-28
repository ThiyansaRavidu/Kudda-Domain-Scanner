# Advanced Network Information Tool

This Go-based network information tool allows users to query a wide range of APIs to obtain detailed information about domains, IP addresses, and DNS records. It is designed to be an all-in-one tool for performing advanced DNS lookups, response checks, and other domain-related queries, making it ideal for system administrators, developers, and network enthusiasts.

**Made by Thiyansa**

## Features

- **DNS Lookup**: Retrieve DNS records such as A, AAAA, MX, NS, CNAME, and SOA.
- **Reverse DNS Lookup**: Identify the domain name associated with an IP address.
- **IP Geolocation Lookup**: Obtain geographical details of an IP address.
- **Reverse IP Lookup**: Find domains hosted on a specific IP address.
- **HTTP Headers**: View HTTP headers of a domain.
- **Page Links**: Extract all links from a webpage.
- **AS Lookup**: Fetch Autonomous System (AS) information for an IP.
- **Advanced DNS Lookup**: Detailed DNS query with advanced options.
- **Response Checker Advanced**: Evaluate HTTP response details including TLS and speed test information.
- **NS Lookup**: Obtain authoritative name servers for a domain.

## Installation

### Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download Go from the [official website](https://golang.org/dl/).
- **Git**: Git is required for cloning the repository. Download Git from [here](https://git-scm.com/downloads).

### Windows Installation

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/ThiyansaRavidu/Kudda-Domain-Scanner.git
    cd Kudda-Domain-Scanner
    ```

2. **Build the Project**:

    ```bash
    go build -o network-info.exe
    ```

3. **Run the Tool**:

    ```bash
    network-info.exe
    ```

### Linux Installation

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/ThiyansaRavidu/Kudda-Domain-Scanner.git
    cd Kudda-Domain-Scanner
    ```

2. **Build the Project**:

    ```bash
    go build -o network-info
    ```

3. **Run the Tool**:

    ```bash
    ./network-info
    ```

### MacOS Installation

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/ThiyansaRavidu/Kudda-Domain-Scanner.git
    cd Kudda-Domain-Scanner
    ```

2. **Build the Project**:

    ```bash
    go build -o network-info
    ```

3. **Run the Tool**:

    ```bash
    ./network-info
    ```

## Usage

Upon running the tool, you will be presented with a menu of API queries you can perform. Select an option by entering the corresponding number, then provide a domain or IP address when prompted.

### Example

1. Select **DNS Lookup** by entering `1`.
2. Enter a domain such as `example.com`.
3. View the detailed DNS records returned by the tool.

## Configuration

The tool is configured with several API endpoints to fetch data. You can add or modify API endpoints directly in the code under the `apiEndpoints` variable.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to suggest improvements or report bugs.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- **API Providers**: Special thanks to the data providers for their APIs.
- **Community**: A big shoutout to all contributors and users who make this tool better every day!

---

**Made by Thiyansa**
