```markdown
# Deployment Script

This script is designed to automate the deployment process for multiple directories. It reads a configuration file (`config.json`) which specifies a base directory, a list of directories, and a set of commands to execute in each directory. The script uses goroutines to run the commands concurrently, improving efficiency and reducing deployment time.

## Features

- Concurrent execution of commands for multiple directories.
- Customizable configuration through a `config.json` file.
- Automatic handling of production environments with additional command flags.

## Prerequisites

- Go programming language installed. Download from [golang.org](https://golang.org/dl/).

## Configuration

Create a `config.json` file with the following structure:

```json
{
  "basedir": "/path/to/base/directory",
  "directories": [
    "dir1",
    "dir2",
    "dir3"
  ],
  "commands": [
    {
      "name": "git pull",
      "command": ["git", "pull"]
    },
    {
      "name": "update database",
      "command": ["/path/to/php", "artisan", "migrate"],
      "forceProduction": true
    }
  ]
}
```

## Compilation Instructions

### Compile for Linux on Windows using PowerShell

1. **Open PowerShell**.

2. **Set environment variables**:

    ```powershell
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    ```

3. **Navigate to the project directory**:

    ```powershell
    cd C:\path\to\your\project
    ```

4. **Build the project**:

    ```powershell
    go build -o deploy
    ```

   This will create an executable named `deploy` in the current directory.

### Compile for the Current Operating System

1. **Open your terminal or command prompt**.

2. **Navigate to the project directory**:

    ```sh
    cd /path/to/your/project
    ```

3. **Build the project**:

    ```sh
    go build -o deploy
    ```

   This will create an executable named `deploy` in the current directory.

## Usage

1. **Transfer the compiled binary to your Linux machine** (if compiled on Windows):

    ```sh
    scp deploy user@linux-machine:/path/to/destination
    ```

2. **Connect to your Linux machine**, navigate to the destination directory, and make the binary executable:

    ```sh
    cd /path/to/destination
    chmod +x deploy
    ```

3. **Run the deployment script**:

    ```sh
    ./deploy
    ```

## Example `config.json`

Here is an example of a `config.json` file that you can use to customize your deployment:

```json
{
  "basedir": "/home/server/domains/domain.com",
  "directories": [
    "dir1",
    "dir2",
    "dir3",
    "dir4",
    "dir5",
    "dir6",
    "dir7",
    "dir8"
  ],
  "commands": [
    {
      "name": "git pull",
      "command": ["git", "pull"]
    },
    {
      "name": "update database",
      "command": ["/opt/cloudlinux/alt-php82/root/usr/bin/php", "artisan", "migrate"],
      "forceProduction": true
    }
  ]
}
```

This setup ensures that your deployment process is automated and efficient, leveraging Go's concurrency features to handle multiple directories and commands simultaneously.