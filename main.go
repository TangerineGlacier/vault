package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/sreevatsan/tangerine-vault/crypto"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// getPassword reads a password from stdin without displaying it
func getPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Add a newline after password input
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

var rootCmd = &cobra.Command{
	Use:   "tangerine-vault",
	Short: "A CLI tool for AES encryption and decryption",
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt content and save to a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("dir")
		name, _ := cmd.Flags().GetString("name")
		inputFile, _ := cmd.Flags().GetString("file")

		var content string
		var err error

		if inputFile != "" {
			// Read content from file
			fileContent, err := os.ReadFile(inputFile)
			if err != nil {
				return fmt.Errorf("failed to read input file: %v", err)
			}
			content = string(fileContent)
		} else {
			// Read content from stdin
			fmt.Print("Enter content to encrypt: ")
			reader := bufio.NewReader(os.Stdin)
			content, err = reader.ReadString('\n')
			if err != nil {
				return err
			}
			content = strings.TrimSpace(content)
		}

		password, err := getPassword("Enter password: ")
		if err != nil {
			return err
		}

		// Confirm password
		confirmPassword, err := getPassword("Confirm password: ")
		if err != nil {
			return err
		}

		if password != confirmPassword {
			return fmt.Errorf("passwords do not match")
		}

		err = crypto.EncryptFile(content, password, dir, name)
		if err != nil {
			return fmt.Errorf("encryption failed: %v", err)
		}

		fmt.Printf("Content encrypted and saved to %s/%s.txt\n", dir, name)
		return nil
	},
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt content from a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("please provide the path to the encrypted file")
		}

		maxAttempts := 3
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			password, err := getPassword("Enter password: ")
			if err != nil {
				return err
			}

			decrypted, err := crypto.DecryptFile(args[0], password)
			if err == nil {
				fmt.Println("Decrypted content:")
				fmt.Println(decrypted)
				return nil
			}

			if attempt < maxAttempts {
				fmt.Printf("Wrong password. %d attempts remaining.\n", maxAttempts-attempt)
			} else {
				return fmt.Errorf("maximum password attempts reached")
			}
		}

		return nil
	},
}

func init() {
	encryptCmd.Flags().String("dir", ".", "Directory to save the encrypted file")
	encryptCmd.Flags().String("name", "encrypted", "Name of the encrypted file (without .txt extension)")
	encryptCmd.Flags().String("file", "", "Path to the file to encrypt")
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 