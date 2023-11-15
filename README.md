# Bulk Repository Transfer

Bulk Repository Transfer is a Go application designed for mass transfers of GitHub repositories from one user to another. This project is a personal endeavor, and I do not assume responsibility for its usage.

## Prerequisites

Before running the application, ensure that you have the following prerequisites installed:

- Go (version 1.16 or higher)
- GitHub API Token (for authentication)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/seu-usuario/bulk-repository-transfer.git
   ```

2. Navigate to the project directory:

   ```bash
   cd bulk-repository-transfer
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Create a file named `repositories.txt` in the project directory and list the repositories you want to transfer, each on a new line.

   ```plaintext
   repository1
   repository2
   ...
   ```

5. Build the application:

   ```bash
   go build -o bulk-repo-transfer
   ```

## Usage

Run the application using the following command:

```bash
go run main.go <OWNER_USERNAME> <TARGET_USERNAME>
```

- `<OWNER_USERNAME>`: The GitHub username of the owner whose repositories you want to transfer.
- `<TARGET_USERNAME>`: The GitHub username to which the repositories will be transferred.

### Example

```bash
go run main.go your-username target-username
```

## Important Note

This application involves transferring repositories, which can have significant consequences. Ensure you have the necessary permissions, and use it responsibly. By using this tool, you acknowledge that the project's author is not responsible for any issues or data loss that may arise from its usage.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
