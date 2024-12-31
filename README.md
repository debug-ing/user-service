# User-Service

Sample User-Service golang

## Prerequisites

- Nix (for managing the environment)
- Go (for building the project)
- k6 (for test)

## Setup with nix

1. Install Nix

    ```bash
    curl -L https://nixos.org/nix/install | sh
    ```

2. Enter the Nix shell:

    ```bash
    nix-shell
    ```

3. Build the project:

    ```bash
    go build cmd/main.go
    ```

4. Run the project:

    ```bash
    ./main
    ```

The server will be running on `http://localhost:8080`, and it will respond with "Hello, World!" when accessed.