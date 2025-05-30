# Tangerine Vault

A CLI tool for AES encryption and decryption of text content.

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
