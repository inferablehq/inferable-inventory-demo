# Setup
## 1. Setup Dependencies

Ensure the following dependencies are available locally
`docker`, `docker-compose`, `git` `node` , `npm`


## 2. Init playwright browsers
Install the required playwright browsers
```
cd ./inventory && npm install && npx playwright install && cd ../
```

## 3. Start The services
```
# Get yours at https://app.inferable.ai/
export INFERABLE_API_SECRET=
export DB_CONNECTION_STRING=postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable
bash ./start.sh
```


## 4. Trigger a new Run
```
curl --request POST \
  --url https://api.inferable.ai/clusters/<CLUSTER_ID>/runs \
  --header 'Content-Type: application/json' \
  --header 'authorization: <API_SECRET>' \
  --data '{
    "message":  "Create a new order for Tablet with 1 items. Reserve the stock from the inventory before creating the order. The customers email is john.doe@example.com, name is John Doe and their address is 123 Main St, Anytown, USA.",
    "result": {
      "schema": {
        "type": "object",
        "properties": {
          "orderId": {
            "type": "string"
          },
          "customerId": {
            "type": "string"
          },
          "reservationId": {
            "type": "string"
          }
        },
        "required": ["orderId", "customerId", "reservationId"]
      }
    }
  }' | jq
```

## 5. Get the results
```
curl --request GET \
  --url https://api.inferable.ai/clusters/<CLUSTER_ID>/runs/<RUN_ID_FROM_PREVIOUS_STEP> \
  --header 'Content-Type: application/json' \
  --header 'authorization: <API_SECRET>' | jq
