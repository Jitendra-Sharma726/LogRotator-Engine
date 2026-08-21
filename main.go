package main 

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "time"
)

const LogFile = "server.log"
const MaxSize = 50 * 1024 // 50 KB

func main() {
  fmt.Println("=== LogRotator Engine Started ===")
  fmt.Prinln("Monitoring:", LogFile)
  fmt.Println("Max allowed size:", MaxSize, "bytes (50 KB)")
  fmt.Println("Press Ctrl+C to stop.")
  fmt.Println("-------------------------------------------")
  fmt.Println()

  rotationCount := 0

  go mockServerWorker()

  //Listen for Ctrl+C so we can shut cleanly
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

  ticker := time.NewTicker(2 * time.Second)

  for {
    select {
      //Ticker fired : inspect and rotate if needed
      case <-ticker.C
      info, err := os.Stat(LogFile)
      if err != nil {
        continueN
      }

      fmt.Printf("[MONITOR] server.log -> %d bytes\n", info.Size())

      wasArchived := CheckAndArchive(LogFile, info.Size(), int64(MaxSize))
      if wasArchived {
        rotationCount++
        fmt.Printf("[SUCCESS] Rotation #%d complete. \n\n", rotationCount)
        }

      //Ctrl+C received: Stop the ticker and exit
      case <-quit:
          ticker.Stop()
          fmt.Println("\n-----------------------------------------------")
          fmt.Printf("LogRotator stopped. Total rotations: %d\n", rotationCount)
          return
    }
  }
}


// === Mock Server Traffic (DO NOT MODIFY) ===
func mockServerWorker() {
  logLines := []string{
    "INFO:   User logged in. IP: 192.168.1.55\n",
    "INFO:   Page requested. URL: /dashboard\n",
    "WARN:   Slow query detected. Duration: 520ms\n",
    "INFO:   User logged out. IP: 192.168.1.55\n",
    "WARN:   High memory usage: 87%\n",
    "ERROR:  Database connection timeout. Retrying...\n",
    "INFO:   File uploaded successfully. Size: 2.4MB\n",
    "DEBUG:  Cache miss for key: user_profile_42\n",
    "INFO:   Password reset email sent to user@example.com\n",
    "FATAL:  Payment gateway API unreachable.\n",
  }

  i := 0
  for {
    f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err == nil {
      line := fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), LogLines[i%len(logLines)])
      f.WriteString(line)
      f.Close()
    }
    i++
    time.Sleep(10 * time.Millisecond)










                   






        
