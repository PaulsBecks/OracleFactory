package models

import (
	"os/exec"
)

type Manifest string

func (m *Manifest) Run() error {
	// Run the manifest in a new docker
	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--network=oracle-factory-network",
		"oracle_blueprint",
		"/bin/bash",
		"-c",
		"echo \""+string(*m)+
			"\" > manifest.bloql; cat manifest.bloql; java -jar Blockchain-Logging-Framework/target/blf-cmd.jar extract manifest.bloql")
	return cmd.Run()
}
