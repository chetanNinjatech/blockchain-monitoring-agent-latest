// utils/logger.go
package utils

import "log"

func InitLogger() {
	// Initialize logger configuration here
	// For example, setting log format, output destination, etc.
	log.SetPrefix("[Blockchain Monitoring] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
