package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: gvm <version>\nExamples: 'gvm go1.21.1' or 'gvm go' (default installation)\n")
		return
	}

	version := os.Args[1]

	if !isVersionInstalled(version) {
		fmt.Printf("Version not found, install it using 'go install golang.org/dl/%s@latest' followed by '%s download'\n", version, version)
		return
	}

	bashrc := os.Getenv("HOME") + "/.bashrc"

	cleanVersion(bashrc)

	setVersion(bashrc, version)

	if version == "go" {
		fmt.Println("The default go installation was set as default for future sessions")
	} else {
		fmt.Printf("%s set was as default for future sessions\n", version)
	}
}

func isVersionInstalled(version string) bool {
	if version == "go" {
		return true
	}

	if err := exec.Command(getGoVersionPath(version), "version").Run(); err != nil {
		return false
	}

	return true
}

func cleanVersion(bashrc string) {
	commands := []string{
		fmt.Sprintf(`sed -i '/alias go=.*/d' %s`, bashrc),
		fmt.Sprintf(`sed -i '/PATH=\$GOROOT\/bin:\$PATH/d' %s`, bashrc),
		fmt.Sprintf(`sed -i '/GOROOT=.*/d' %s`, bashrc),
	}

	execCommands(commands)
}

func setVersion(bashrc, version string) {
	if version == "go" {
		return
	}
	commands := []string{
		fmt.Sprintf(`echo "GOROOT=$(%s env GOROOT)" >> "%s"`, getGoVersionPath(version), bashrc),
		fmt.Sprintf(`echo "PATH=\$GOROOT/bin:\$PATH" >> "%s"`, bashrc),
		fmt.Sprintf(`echo "alias go='%s'" >> "%s"`, version, bashrc),
	}

	execCommands(commands)
}

func execCommands(commands []string) {
	for _, command := range commands {
		if err := exec.Command("sh", "-c", command).Run(); err != nil {
			fmt.Printf("error executing command %s: %s\n", command, err.Error())
			return
		}
	}
}

func getGoVersionPath(version string) string {
	goRoot := os.Getenv("GOROOT")
	if goRoot != "" {
		return goRoot + "/bin/" + version
	} else {
		return os.Getenv("HOME") + "/go/bin/" + version
	}
}
