package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/dsnet/compress/bzip2"
	"github.com/google/brotli/go/cbrotli"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
	"github.com/ulikunitz/xz"
	"github.com/xuri/excelize/v2"
)

type CompressionResult struct {
	Algorithm        string
	CompressedSizeMB float64
	Duration         time.Duration
	CompressionRatio float64
	CompressionSpeed float64
}

func bytesToMB(bytes int) float64 {
	return float64(bytes) / (1024 * 1024)
}

// compressRAR сжимает данные с использованием утилиты RAR
func compressZip(data []byte, originalSize int64) CompressionResult {
	start := time.Now()

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	w, err := zipWriter.Create("data")
	if err != nil {
		fmt.Println("Error creating ZIP writer:", err)
		return CompressionResult{}
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("Error writing to ZIP:", err)
		return CompressionResult{}
	}

	err = zipWriter.Close()
	if err != nil {
		fmt.Println("Error closing ZIP writer:", err)
		return CompressionResult{}
	}

	duration := time.Since(start)
	compressedSize := buf.Len()
	compressionRatio := float64(originalSize) / float64(compressedSize)
	compressionSpeed := float64(originalSize) / duration.Seconds()

	return CompressionResult{
		Algorithm:        "ZIP",
		CompressedSizeMB: bytesToMB(compressedSize),
		CompressionRatio: compressionRatio,
		CompressionSpeed: compressionSpeed,
		Duration:         duration,
	}
}

func compressGzip(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer := gzip.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "gzip",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressZlib(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer := zlib.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "zlib",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressBzip2(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer, _ := bzip2.NewWriter(&b, &bzip2.WriterConfig{Level: bzip2.BestCompression})
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "bzip2",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressLzw(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer := lzw.NewWriter(&b, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "lzw",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressLzma(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer, _ := xz.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "lzma",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressXZ(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer, _ := xz.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "xz",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressBrotli(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer := cbrotli.NewWriter(&b, cbrotli.WriterOptions{Quality: 11})
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "brotli",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressLz4(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer := lz4.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "lz4",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func compressZstd(data []byte, originalSize int64) CompressionResult {
	start := time.Now()
	var b bytes.Buffer
	writer, _ := zstd.NewWriter(&b)
	writer.Write(data)
	writer.Close()
	duration := time.Since(start)
	compressedSize := b.Len()
	return CompressionResult{
		Algorithm:        "zstd",
		CompressedSizeMB: bytesToMB(compressedSize),
		Duration:         duration,
		CompressionRatio: float64(originalSize) / float64(compressedSize),
		CompressionSpeed: float64(originalSize) / duration.Seconds(),
	}
}

func writeResultsToXLSX(results []CompressionResult, filename string) error {
	f := excelize.NewFile()
	sheetName := "Compression Results"

	index, _ := f.NewSheet(sheetName)

	// Записываем заголовки
	headers := []string{"Algorithm", "Original Size (MB)", "Compressed Size (MB)", "Compression Ratio", "Compression Speed (B/s)", "Time Taken (seconds)"}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// Записываем результаты
	for i, result := range results {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), result.Algorithm)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), result.CompressedSizeMB)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), result.CompressionRatio)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), result.CompressionSpeed)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), result.Duration.Seconds())
	}

	f.SetActiveSheet(index)
	return f.SaveAs(filename)
}

func main() {
	filePath := "testfile/test.avi" // Укажите путь к вашему файлу
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fileSize := fileInfo.Size()
	fmt.Printf("Original File Size: %.2f MB\n", bytesToMB(int(fileSize)))

	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}

	var wg sync.WaitGroup
	results := make(chan CompressionResult, 10)

	compressionFuncs := []func([]byte, int64) CompressionResult{
		compressGzip,
		compressZlib,
		compressBzip2,
		compressLzw,
		compressLzma,
		compressXZ,
		compressBrotli,
		compressLz4,
		compressZstd,
		compressZip,
		//	compressRAR,
		//	compress7z,
	}

	for _, compressFunc := range compressionFuncs {
		wg.Add(1)
		go func(compress func([]byte, int64) CompressionResult) {
			defer wg.Done()
			results <- compress(data, fileSize)
		}(compressFunc)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var bestTimeResult, bestSizeResult CompressionResult
	firstResult := true
	for result := range results {
		fmt.Printf("Algorithm: %s, Compressed Size: %.2f MB, Compression Ratio: %.2f, Compression Speed: %.2f B/s, Time Taken: %s\n",
			result.Algorithm, result.CompressedSizeMB, result.CompressionRatio, result.CompressionSpeed, result.Duration)

		if firstResult || result.Duration < bestTimeResult.Duration {
			bestTimeResult = result
		}
		if firstResult || result.CompressionRatio > bestSizeResult.CompressionRatio {
			bestSizeResult = result
		}
		firstResult = false
	}

	fmt.Printf("\nBest Time Result: %s (Time Taken: %s, Speed: %.2f B/s)\n", bestTimeResult.Algorithm, bestTimeResult.Duration, bestTimeResult.CompressionSpeed)
	fmt.Printf("Best Compression Result: %s (Compression Ratio: %.2f, Compressed Size: %.2f MB)\n", bestSizeResult.Algorithm, bestSizeResult.CompressionRatio, bestSizeResult.CompressedSizeMB)

	err = writeResultsToXLSX([]CompressionResult{bestTimeResult, bestSizeResult}, "compression_results.xlsx")
	if err != nil {
		fmt.Println("Error saving results to XLSX file:", err)
	}
}
