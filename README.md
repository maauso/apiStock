# apiStock (WIP)
## Command to run 
```bash
apiStock 
  -apikey string
        demo apikey (default "demo")
  -company string
        demo apikey (default "AAPL")
  -metric string
        Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric };(Required)
                                                                                 -provider string
Select the provider [ financialmodelingprep ] (Required)

```

```
go run cmd/apiStock/apiStock.go
https://financialmodelingprep.com/api/v3/company/discounted-cash-flow/AAPL?&apikey=demo
Symbol: AAPL
Date: 2021-02-13
DiscountCashFlow: 137.53884618136223
StockPrice: 135.37
"137.53884618136223"
```