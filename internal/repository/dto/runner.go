package dto

import (
	"time"

	"github.com/Qwerty10291/rankode-runner/internal/repository/models"
)

type RunRequest struct {
	Image     string
	Code      string
	Input     []string
	Timeout   time.Duration
	MemoryLimit int
	MaxFilesSize int
	MaxOutputSize int
}

type RunResult struct {
	Status        models.AttemptStatus
	Error string
	Output        []string
	ExecutionTime int64
	MemoryUsage   int
}
