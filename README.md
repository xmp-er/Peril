# 🚀 Peril CLI Application

Welcome to Peril, a command-line tool for managing files, encoding files and interacting with Google Drive. This guide outlines the available commands and how to use them.

The encryption of file is performed using ChaCha20 encryption and can only be decrypted via 32-character password generated during encryption.

## 🛠️ Prerequisites

Dependencies are automatically resolved when you run the application.

## 📋 Available Commands

| Command                     | Description                                      | Example Usage                              | Notes                                     |
|-----------------------------|--------------------------------------------------|--------------------------------------------|-------------------------------------------|
| `open <name>`               | 📂 Create or open a file                         | `open example.txt`                         |                                           |
| `efile <name>`              | 🔒 Encrypt a file and delete the original        | `efile example.txt`                        | Password for decryption will be displayed |
| `dfile <name> <pass>`       | 🔓 Decrypt a file with a 32-character password   | `dfile example.txt.enc your-32-char-pass`  | Password must be exactly 32 characters    |
| `del <name_with_extension>` | 🗑️ Delete a file                                 | `del example.txt`                          |                                           |
| `up <Drive_Path>`           | ☁️ Upload a file to Google Drive                 | `up /path/to/file/example.txt`             |                                           |
| `down <Drive_Path>`         | 📥 Download a file from Google Drive             | `down /path/to/file/example.txt`           |                                           |
| `vi <name>`                 | ✏️ Open a file in the `vi` editor                | `vi example.txt`                           |                                           |
| `help` or `h`               | ❓ Display available options                      | `help`                                     |                                           |
| `exit`                      | 🚪 Exit the application                          | `exit`                                     |                                           |

## 🚦 Getting Started

Run Peril and use the commands listed above to manage your files. Type `help` if you need a quick reminder of available commands.

## ⚠️ Error Handling

If an error occurs, Peril will provide feedback. Make sure to follow the usage guidelines.

## 📝 Notes

- **🔑 Password Requirement**: Ensure your password for decryption is exactly 32 characters.
- **📁 File Paths**: Use the correct paths when uploading or downloading files.
- **📦 Dependencies**: Automatically resolved on application start.

---

Enjoy using Peril for efficient file management and Google Drive operations! 🎉
