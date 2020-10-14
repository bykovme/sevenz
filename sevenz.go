package sevenz

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Check7zAvailability() error {
	cmd := exec.Command(sevenCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
	}
	return err
}

func Extract(archiveFullPath string, destFullPath string) error {
	destPathParam := strings.Replace(destFolderTemplate, templatePattern, destFullPath, -1)
	cmd := exec.Command(sevenCmd, extractCmd, destPathParam, archiveFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}

func CompressStandard(archiveFullPath string, srcFullPath string) error {
	cmd := exec.Command(sevenCmd, archiveCmd, archiveType, archiveFullPath, srcFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}

func CompressInsane(archiveFullPath string, srcFullPath string) error {
	var args []string
	//cmd := exec.Command(sevenCmd, archiveCmd, archiveType, insaneCompressionParams, archiveFullPath, srcFullPath)
	cmd := exec.Command(sevenCmd)

	args = append(args, sevenCmd, archiveCmd, archiveType)
	args = append(args, strings.Split(insaneCompressionParams, " ")...)
	args = append(args, archiveFullPath, srcFullPath)
	cmd.Args = args
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}

func CompressCustom(archiveFullPath string, srcFullPath string, customParams string) error {
	cmd := exec.Command(sevenCmd, archiveCmd, archiveType, customParams, archiveFullPath, srcFullPath)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	return err
}
