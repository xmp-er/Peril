package helper

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/chacha20"
)

func EncryptAndDeleteOriginal(fileName, current_folder_location string) (string, error) {
	//Read the file Contents,generate a password,delete the original file,return the password
	fileName = current_folder_location + "/" + fileName
	fileData, err := os.Open(fileName + ".md")
	if err != nil {
		return "", fmt.Errorf("🔴[ERROR] error opening the file %v as %v, aborting encryption", (fileName + ".md"), err)
	}
	scanner := bufio.NewScanner(fileData)
	var finalString string
	isFirstLine := true
	for scanner.Scan() {
		curr_line := scanner.Text()
		if isFirstLine {
			finalString = curr_line
			isFirstLine = false
		} else {
			finalString = finalString + "\n" + curr_line
		}
	}
	defer fileData.Close()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("🔴[ERROR] error reading the file %v as %v, aborting encryption", (fileName + ".md"), err)
	}

	password, err := GeneratePassword(32)
	if err != nil {
		return "", fmt.Errorf("🔴[ERROR] error generating password for the file %v as %v, aborting encryption", (fileName + ".md"), err)
	}
	encryptedFileData, err := Encrypt([]byte(finalString), []byte(password))
	if err != nil {
		return "", fmt.Errorf("🔴[ERROR] error encrypting file %v as %v, aborting encryption", (fileName + ".md"), err)
	}

	//Storing the file as .enc format

	fileOutputName := fmt.Sprintf("%v.enc", fileName)
	err = os.WriteFile(fileOutputName, encryptedFileData, 0644)
	if err != nil {
		return "", fmt.Errorf("🔴[ERROR] error writing data to file %v from file %v as %v, aborting encryption", (fileName + ".md"), fileOutputName, err)
	}

	//log that the file was encrypted and the original file was deleted
	msg := fmt.Sprintf("🟢[DONE] the file %v.md was encrypted to %v.enc and the original file was deleted", fileName, fileName)
	fmt.Println(msg)

	//Delete the original file
	err = os.Remove(fileName + ".md")
	if err != nil {
		return "", fmt.Errorf("🔴[ERROR] error deleting the original file %v as %v", (fileName + ".md"), err)
	}

	return password, nil
}

func DecryptAndRecoverOriginal(fileName, pass, current_folder_location string) error {
	fileName = current_folder_location + "/" + fileName
	if _, err := os.Stat(fileName + ".enc"); os.IsNotExist(err) {
		//File does not exist, log it.
		return fmt.Errorf("🔴[ERROR] file %v does not exist at current loaction, aborting decryption", (fileName + ".enc"))
	}
	fileData, err := os.ReadFile(fileName + ".enc")
	if err != nil {
		return fmt.Errorf("🔴[ERROR] error opening the file %v as %v, aborting decryption", (fileName + ".enc"), err)
	}
	res, err := Decrypt(fileData, []byte(pass))
	if err != nil {
		return fmt.Errorf("🔴[ERROR] error decrypting the file %v as %v, aborting decryption", (fileName + ".enc"), err)
	}

	err = os.WriteFile(fileName+".md", res, 0644)
	if err != nil {
		return fmt.Errorf("🔴[ERROR] error writing decrypted content to file %v as %v", (fileName + ".md"), err)
	}

	err = os.Remove(fileName + ".enc")
	if err != nil {
		return fmt.Errorf("🔴[ERROR] error deleting the original file %v as %v", (fileName + ".md"), err)
	}

	msg := fmt.Sprintf("🟢[DONE] the file %v.enc was decrypted to %v.md and the original file was deleted", fileName, fileName)
	fmt.Println(msg)

	//Opening the decrypted file
	cmd := exec.Command("vi", fileName+".md")

	// Set the command to run in the current terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("🔴[ERROR] Error executing vi: %v\n", err)
	}
	return nil
}

// Got this shit from Claude, idk how it operates but it encrypts array of bytes via ChaCha20 encryption with a specefic key
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// Generate a random nonce
	nonce := make([]byte, chacha20.NonceSizeX)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)

	// Prepend the nonce to the ciphertext
	return append(nonce, ciphertext...), nil
}

// Got this shit from Claude, idk how it operates but it decrypts array of bytes via ChaCha20 encryption with a specefic key
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// Extract the nonce from the ciphertext
	nonce, ciphertext := ciphertext[:chacha20.NonceSizeX], ciphertext[chacha20.NonceSizeX:]

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	plaintext := make([]byte, len(ciphertext))
	cipher.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// Got from clause, generates 32-length password
func GeneratePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?"

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	// Ensure the password contains at least one of each character type
	types := []string{"[a-z]", "[A-Z]", "[0-9]", "[!@#$%^&*()-_=+\\[\\]{}|;:,.<>?]"}
	for _, t := range types {
		if !strings.ContainsAny(string(password), t) {
			randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
			charOfType, _ := rand.Int(rand.Reader, big.NewInt(int64(len(t)-2))) // -2 for the brackets
			password[randomIndex.Int64()] = t[charOfType.Int64()+1]             // +1 to skip the opening bracket
		}
	}

	return string(password), nil
}
