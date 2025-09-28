package pipeline

import (
	"os"
)

// Step interface
type Step interface {
	Run() error
}

type DownloadStep struct {
	InputPath  string
	OutputPath string
}

func (s *DownloadStep) Run() error {
	log.Infof("[DownloadStep] DownloadStep file: %s → %s", s.InputPath, s.OutputPath)
	// TODO: implement chunk splitting
	return nil
}

type SplitStep struct {
	InputPath  string
	OutputPath string
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

type CleanStep struct {
	TmpPath string
}

func (c *CleanStep) Run() error {
	log.Infof("[CleanStep] Cleaning up temporary files in %s\n", c.TmpPath)
	return os.RemoveAll(c.TmpPath)
}
