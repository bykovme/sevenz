package sevenz

import (
	"log"
	"os/exec"
	"strings"
)

func Check7zAvailability() error {
	cmd := exec.Command(sevenCmd)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}

func Extract7zArchive(archiveFullPath string, destFullPath string) error {
	destPathParm := strings.Replace(destFolderTemplate, templatePattern, destFullPath, -1)
	cmd := exec.Command(sevenCmd, extractCmd, destPathParm, archiveFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}
