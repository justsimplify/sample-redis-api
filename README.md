## NOTES:

- This performs basic operations related to Redis.
- Following are the endpoints available. 
- All requests are `GET`

- Required Headers:
  - `redis_host`- Redis Instance address to connect to
  - `redis_port`- Redis Instance port to connect to
  
- Endpoints:
  - `/add/:key/:value` - Adds key to the redis
  - `/delete/:key` - Deletes key from redis
  - `/get/:key` - Gets key from redis
  
- Response format:
  - ```json
    {
      "message": "<response-message>",
      "error": "<error-message>"
    }
    ```