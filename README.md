# Receipt Processor

Run `go build -o receipt-processor
./receipt-processor` command to start server at `localhost:8080`

## Test POST /receipts/process
Run `curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d @data.json`
### data.json
```
{
    "retailer": "M&M Corner Market",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      }
    ],
    "total": "9.00"
  }
```

Output : `{"id":"1c98c8f1-c958-47ae-a0a9-bd11536b4298"}`

## Test GET /receipts/{id}/points
Run `curl -X GET http://localhost:8080/receipts/1c98c8f1-c958-47ae-a0a9-bd11536b4298/points` to get the points

Output : `{"points":109}`

