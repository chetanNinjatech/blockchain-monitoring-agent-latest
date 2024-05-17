Blockchain Monitoring Agent
Description
The Blockchain Monitoring Agent is a project developed to demonstrate Web3 Security skills by creating an external agent in Go. This agent monitors a Blockchain network, extracts relevant operational metrics, and identifies potential vulnerabilities, threats, and attacks in real-time. The project includes a local Blockchain setup using Docker and Docker Compose to simulate specific scenarios for monitoring purposes, ensuring compatibility with both open nodes and local blockchain setups. 

Table of Contents:-
Installation Instructions
Usage Instructions
Using Docker
Testing
Configuration Details
Contributing
License
Contact Information
Acknowledgements
Additional Resources

Installation Instructions
Prerequisites
Go 1.21 or higher (https://golang.org/doc/install)
Docker (https://docs.docker.com/get-docker/)
Docker Compose (https://docs.docker.com/compose/install/)

Steps
1) Clone the repository
git clone https://github.com/chetanNinjatech/blockchain-monitoring-agent.git
cd blockchain-monitoring-agent 

2) Install dependencies:-
go mod download

Usage Instructions
Running the Project

1)Build the project:
(go build -o blockchain-monitoring-agent)

2)Run the project:
()./blockchain-monitoring-agent)

Using Docker:-
Building the Docker Image
(docker-compose build)

Running the Docker Container
(docker-compose up)

This will set up and run the local Ethereum blockchain and the monitoring agent. The Ethereum node will be exposed on port 8545.

Testing
Running Tests
To run the tests for the project, use the following command:
go test ./...

Configuration Details
The configuration for the agent is specified in the config.yaml file:
api_endpoint: "http://172.17.0.2:8545"
notification_email: "chetan.n@ninjatechnolabs.com" 

api_endpoint: The endpoint of the blockchain node.
notification_email: Email address to receive notifications.

Customizing the Configuration
To customize the configuration, modify the values in config.yaml to match your environment. 

License
This project is licensed under the MIT License.

Acknowledgements
Ethereum Go Client (https://github.com/ethereum/go-ethereum)





