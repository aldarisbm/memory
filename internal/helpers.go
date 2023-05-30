package internal

import (
	"fmt"
	"os"
	"os/user"
)

// CreateFileInHomeDir creates a file at the user's home directory
// inside the xyz.memorystore directory
func CreateFileInHomeDir(fileName string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	_ = os.Mkdir(fmt.Sprintf("%s/%s", dir, DomainName), os.ModePerm)
	return fmt.Sprintf("%s/%s/%s", dir, DomainName, fileName)
}
