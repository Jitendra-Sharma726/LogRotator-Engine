package main

import (
  "compress/gzip"
  "fmt"
  "io"
  "time"
)

// checkAndArchive compresses oversized logs and returns true on success.
func checkAndArchive(filename string, currentSize int64, maxSize int64) bool {

  if currentSize <= maxSize {
    return false
  }

  fmt.Printf("\n[ALERT] Log reached %d bytes. Rotating...\n", currentSize)

  //Step 1: Build timestamped archive name
  timestamp := time.Now().Format("2006-01-02_15-04-05")
  archiveName := fmt.Sprintf("server-%s.log.gz", timestamp)


  //Step 2: Open original log file
  original, err := os.Open(filename)
  if err != nil {
       fmt.Println("[Error] Could not open log file:", err)
       return false
  }

  //Step 3: Create .gz archieve file
  archive, err := os.Create(archiveName)
  if err != nil {
     fmt.Println("[Error] Could not create archive:", err)
     original.Close()
     return false
  }

  //Step 4: Wrap archive in a gzip Writer
  gzWriter := gzip.NewWriter(archive)

  //Step 5: Stream log data into archive
  _, err = io.Copy(gzWriter, original)
  if err != nil {
    fmt.Println("[ERROR] Could not compress log:", err)
    gzWriter.Close()
    archive.Close()
    original.Close()
    return false
  }

  //Step 6:





















    
    
