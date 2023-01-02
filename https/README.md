# Self Signed HTTPS Certificate

## How to

### 1. Write Ca Certificate (server.ca.crt) to System

### 2. Start Server

```bash
go run .
```

### 3. Verify
* Method 1: Open Browser
  * visit: https://zsxxx.com:9996
* Method 2: Curl
  * run: `curl --cacert $PWD/server.ca.crt --key $PWD/server.key https://zsxxx.com:9996`
