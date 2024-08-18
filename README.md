# Peril CLI Application

Welcome to Peril, a command-line application for managing files and interacting with Google Drive. This guide will walk you through the available commands and how to use them.

## Prerequisites

Before you start, ensure you have installed all necessary dependencies. This is handled automatically when you run the application.

## Available Commands

Below are the commands you can use with Peril:

### 1. `open <name>`

- **Description**: Create or open a file with the given `<name>`.
- **Usage**: `open example.txt`

### 2. `efile <name>`

- **Description**: Encrypt a file with the given `<name>` and delete the original file.
- **Usage**: `efile example.txt`
- **Output**: The password for decrypting the file will be displayed.

### 3. `dfile <name> <pass>`

- **Description**: Decrypt a file with the given `<name>` using the provided `<pass>` (32-character password).
- **Usage**: `dfile example.txt.enc your-32-char-password`

### 4. `del <name_with_extension>`

- **Description**: Delete a file with the given `<name_with_extension>`.
- **Usage**: `del example.txt`

### 5. `up <Drive_Path>`

- **Description**: Upload a file to Google Drive at the specified `<Drive_Path>`.
- **Usage**: `up /path/to/file/example.txt`

### 6. `down <Drive_Path>`

- **Description**: Download a file from Google Drive from the specified `<Drive_Path>`.
- **Usage**: `down /path/to/file/example.txt`

### 7. `vi <name>`

- **Description**: Open a file with the given `<name>` in the `vi` editor.
- **Usage**: `vi example.txt`

### 8. `help` or `h`

- **Description**: Display the list of available commands and their usage.
- **Usage**: `help`

### 9. `exit`

- **Description**: Exit the application.
- **Usage**: `exit`

## Getting Started

To begin using Peril, run the application and enter any of the above commands based on what you'd like to do. If you're unsure, type `help` to see the list of available commands.

## Error Handling

Peril provides feedback for any errors encountered during execution. If a command is not recognized or if there is an issue, an error message will be displayed.

## Notes

- Ensure your password for decryption is exactly 32 characters.
- Use the correct file paths when uploading or downloading files to/from Google Drive.
- The application automatically resolves dependencies when launched.

---

Enjoy using Peril for your file management and Google Drive interactions!

