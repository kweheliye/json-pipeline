package pipeline

import (
	"io"
	"json-pipeline/internal/download"
	"json-pipeline/pkg/utils"
	"os"
	"path/filepath"
	"time"
)

// Step interface
type Step interface {
	Name() string
	Run() error
}

type DownloadStep struct {
	InputPath  string
	OutputPath string
}

func (s *DownloadStep) Run() error {
	log.Infof("[DownloadStep] DownloadStep file: %s → %s", s.InputPath, s.OutputPath)

	o := filepath.Dir(s.OutputPath)

	config := &download.DownloadConfig{
		ChunkSize:  1024 * 1024 * 50, // 50MB chunks
		Timeout:    2 * time.Minute,  // 2 minute timeout
		MaxRetries: 5,                // 5 retry attempts
	}

	err := os.MkdirAll(o, 0o755)
	utils.ExitOnError(err)

	d, err := download.DownloaderFactory(s.InputPath, config)
	utils.ExitOnError(err)

	rd, err := d.DownloadReader(s.InputPath)
	utils.ExitOnError(err)

	defer rd.Close()

	wr, err := os.Create(s.OutputPath)
	utils.ExitOnError(err)

	defer wr.Close()

	n, err := io.Copy(wr, rd)
	utils.ExitOnError(err)

	log.Infof("Downloaded %d bytes from %s to %s", n, s.InputPath, s.OutputPath)

	return nil
}

func (s *DownloadStep) Name() string {
	return "Download"
}

type SplitStep struct {
	InputPath  string
	OutputPath string
}

func (s *SplitStep) Name() string {
	return "Split"
}

func (s *SplitStep) Run() error {
	log.Infof("[SplitStep] Splitting file: %s → %s", s.InputPath, s.OutputPath)
	// TODO: implement chunk splitting
	return nil
}

type ParseStep struct {
	InputPath  string
	OutputPath string
}

func (p *ParseStep) Run() error {
	log.Infof("[ParseStep] Parsing chunks from %s → %s\n", p.InputPath, p.OutputPath)
	// TODO: implement JSON flatten + parquet write
	return nil
}

func (s *ParseStep) Name() string {
	return "Parse"
}

type CleanStep struct {
	TmpPath string
}

func (c *CleanStep) Run() error {
	log.Infof("[CleanStep] Cleaning up temporary files in %s\n", c.TmpPath)
	return nil
}

func (s *CleanStep) Name() string {
	return "Clean"
}
