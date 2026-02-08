package all

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var protocols = []string{"hysteria", "ss", "trojan", "vless", "vmess", "other"}

func merge() error {
	baseDirs := []string{"subs/", "telegram/"} // Directories to read input files from
	outputDir := "all/"                        // Output directory

	allEntriesMap := make(map[string]struct{}) // For all.txt

	for _, proto := range protocols {
		fmt.Printf("[INFO] Processing protocol: %s\n", proto)
		entriesMap := make(map[string]struct{}) // Deduplicate entries per protocol
		totalEntries := 0
		for _, dir := range baseDirs {
			filePath := filepath.Join(dir, proto+".txt") // Input file path
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("[WARN] Missing file: %s\n", filePath)
				continue // Skip if file is missing
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text()) // Read and trim each line
				if line != "" {
					entriesMap[line] = struct{}{} // Add to map if not empty
					// Only add to allEntriesMap if not hysteria or other
					if proto != "hysteria" && proto != "other" {
						allEntriesMap[line] = struct{}{}
					}
					totalEntries++
				}
			}
			file.Close()
		}
		fmt.Printf("[INFO] Total entries: %d\n", totalEntries)

		outPath := filepath.Join(outputDir, proto+".txt") // Output file path
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("failed to create output file for %s: %w", proto, err)
		}
		// Write deduplicated entries
		for entry := range entriesMap {
			_, _ = outFile.WriteString(entry + "\n")
		}
		outFile.Close()
		fmt.Printf("[SUMMARY] %s: %d unique entries saved to %s\n", proto, len(entriesMap), outPath)
	}

	// Write all unique entries to all.txt (excluding hysteria and other)
	allOutPath := filepath.Join(outputDir, "all.txt")
	allOutFile, err := os.Create(allOutPath)
	if err != nil {
		return fmt.Errorf("failed to create all.txt: %w", err)
	}
	for entry := range allEntriesMap {
		_, _ = allOutFile.WriteString(entry + "\n")
	}
	allOutFile.Close()
	fmt.Printf("[SUMMARY] all.txt: %d unique entries saved to %s\n", len(allEntriesMap), allOutPath)

	return nil
}

func Run() {
	err := merge()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
