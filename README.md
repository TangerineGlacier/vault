# Tangerine Vault

A secure CLI tool for storing and managing your passkeys and sensitive credentials using AES-256-GCM encryption.


## Problem

For a long time, I had been storing sensitive information—like passkeys, GitHub Actions secrets, passwords, and recovery keys—in a single Obsidian note titled “Keys.” Needless to say, this wasn’t exactly a secure solution.

## Solution

There are robust tools out there—like Ansible Vault or HashiCorp Vault—that could’ve solved this. But I didn’t want the overhead of setting up and learning a new system just to securely encrypt and decrypt a simple file. I needed something fast, lightweight, and seamless to integrate into my workflow.

So I built a CLI tool to do exactly that. It’s written in Go, because I love Go, and it gets the job done—securely and efficiently—without getting in the way.


## Installation

```bash
# Clone the repository
git clone https://github.com/sreevatsan/tangerine-vault.git
cd tangerine-vault

# Install the tool
make install
```

## Usage

### Encrypt Content

```bash
tangerine-vault encrypt --dir /path/to/directory --name mysecret
```

This will:
1. Prompt you to enter the content you want to encrypt
2. Ask for a password
3. Save the encrypted content to `/path/to/directory/mysecret.txt`

### Decrypt Content

```bash
tangerine-vault decrypt /path/to/encrypted/file.txt
```

This will:
1. Prompt you for the password
2. Display the decrypted content

## Security Notes

- The tool uses AES-256-GCM for encryption
- Each encrypted file includes a random salt
- Keep your password secure and never share it
- The encrypted files are stored in base64 format with .txt extension
