package internal

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
)

// CreateMemoryFolderInHomeDir creates a folder at the user's home directory
func CreateMemoryFolderInHomeDir() string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	_ = os.Mkdir(fmt.Sprintf("%s/%s", dir, DomainName), os.ModePerm)
	return fmt.Sprintf("%s/%s", dir, DomainName)
}

// CreateFileInHomeDir creates a file at the user's home directory inside the
// xyz.memorystore folder
func CreateFolderInsideMemoryFolder(folderName string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	_ = os.Mkdir(fmt.Sprintf("%s/%s/%s", dir, DomainName, folderName), os.ModePerm)
	return fmt.Sprintf("%s/%s/%s", dir, DomainName, folderName)
}

func Generate(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
