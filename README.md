# PhoenEX
### <u>Implementation</u>
* Implemented a `Matching Engine` from scratch in go.
* Enabled ETH transactions by setting up `Ganache server`
* Used `REST` architecture to build all APIs for interacting with the matching Engine.
* Also implemented **Unit Tests**.
--- 
### <u>Setup</u>
1. First, ```git clone https://github.com/kushagra-gupta01/PhoenEX.git``` 
2. Navigate to the cloned repo by doing cd into the repo in your local system.
3. Then to install all dependencies, ```go mod download```
4. Setup **Ganache Server** in your local system to setup your private BlockChain.
    * Install Ganache
    * Then create a Workspace in your Ganache Server.
    * Get the private keys and address from there and replace them in the code.
5. Then execute ```make run```