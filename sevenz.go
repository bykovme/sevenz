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

func Extract(archiveFullPath string, destFullPath string) error {
	destPathParam := strings.Replace(destFolderTemplate, templatePattern, destFullPath, -1)
	cmd := exec.Command(sevenCmd, extractCmd, destPathParam, archiveFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}

func Compress(archiveFullPath string, srcFullPath string) error {
	cmd := exec.Command(sevenCmd, archiveCmd, archiveFullPath, srcFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}
