package main

import (
	"context"
	"log"
	"time"

	"blockchain-monitoring-agent/analysis"
	"blockchain-monitoring-agent/metrics"
	"blockchain-monitoring-agent/monitoring"
	"blockchain-monitoring-agent/notifier"
	"blockchain-monitoring-agent/utils"
)

func main() {
	// Initialize logger
	utils.InitLogger()
	log.Println("Starting blockchain monitoring agent")

	// Create a new monitoring agent
	agent := monitoring.NewAgent()

	// Start the monitoring agent
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration for monitoring
	monitoringConfig, err := monitoring.LoadConfig()
	utils.CheckError(err, "Failed to load monitoring configuration")

	go agent.Start(ctx, monitoringConfig) // Pass monitoringConfig instead of cfg

	// Set up metrics (example function, needs implementation in metrics package)
	metrics.Setup()

	// Example block data for testing vulnerability detection
	block := &analysis.Block{
		Hash: "example_block_hash",
		Transactions: []analysis.Transaction{
			{From: "Alice", To: "Bob", Amount: 100, Hash: "tx1", Timestamp: time.Now()},
			{From: "Alice", To: "Alice", Amount: 50, Hash: "tx2", Timestamp: time.Now().Add(1 * time.Second)},
			{From: "Alice", To: "Alice", Amount: 50, Hash: "tx3", Timestamp: time.Now().Add(2 * time.Second)},
		},
	}

	// Detect vulnerabilities in the block
	log.Println("Starting vulnerability detection...")
	analysis.DetectVulnerabilities(block)
	log.Println("Vulnerability detection completed.")

	// Send notification
	notifier.Notify("Vulnerability detection completed")

	// Start the monitoring loop concurrently
	go monitoring.MonitoringLoop(ctx, monitoringConfig)

	// Keep the application running
	log.Println("Press Ctrl+C to stop...")
	select {}
}
