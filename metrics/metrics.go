// monitoring/metrics.go
package metrics

import (
	"log"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
)

// Setup (optional, if needed for initialization)
func Setup() {
	// Add any initialization logic for the metrics package here (e.g., connecting to a metrics server)
	log.Println("Metrics package initialized.")
}

func GetBlockCount() (int64, error) {
	client, err := rpc.Dial(viper.GetString("api_endpoint"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	var blockCount int64
	err = client.Call(&blockCount, "eth_blockNumber")
	if err != nil {
		log.Fatalf("Failed to retrieve block count: %v", err)
	}

	return blockCount, nil
}
