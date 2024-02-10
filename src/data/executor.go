package data

import (
	"context"
	"os/exec"
	"strings"

	"github.com/PhoenixxZ2023/CheckerDtunnel-GO/src/domain/contract"
)

type bashExecutor struct {
}

func NewBashExecutor() contract.Executor {
	return &bashExecutor{}
}

func (b *bashExecutor) Execute(ctx context.Context, command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}
