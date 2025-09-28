package cmd

import (
	"bufio"
	"fmt"
	"json-pipeline/internal/pipeline"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the JSON → Parquet pipeline",
	Run:   runPipeline,
}

func init() {
	runCmd.Flags().StringVar(&inputPath, "input", "", "Input file path or URL")
	runCmd.Flags().StringVar(&outputPath, "output", "", "Output Parquet file path")
}

func runPipeline(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	if inputPath == "" {
		fmt.Print("Enter input file path or URL: ")
		inputRaw, _ := reader.ReadString('\n')
		inputPath = strings.TrimSpace(inputRaw)
	}
	if outputPath == "" {
		fmt.Print("Enter output Parquet file path: ")
		outputRaw, _ := reader.ReadString('\n')
		outputPath = strings.TrimSpace(outputRaw)
	}

	if inputPath == "" || outputPath == "" {
		log.Fatal("Input and output paths cannot be empty")
	}

	log.Infof("Running pipeline with input: %s, output: %s", inputPath, outputPath)

	tmpDir := os.TempDir()

	steps := []pipeline.Step{
		&pipeline.DownloadStep{
			InputPath:  inputPath,
			OutputPath: filepath.Join(tmpDir, "data.json"),
		},
		&pipeline.SplitStep{
			InputPath:  inputPath,
			OutputPath: filepath.Join(tmpDir, "split"),
		},
		&pipeline.ParseStep{
			InputPath:  filepath.Join(tmpDir, "split"),
			OutputPath: outputPath,
		},
		&pipeline.CleanStep{
			TmpPath: tmpDir,
		},
	}

	p := pipeline.New(steps...)
	if err := p.Run(); err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}
}
