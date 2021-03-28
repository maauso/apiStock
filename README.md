# apiStock (WIP)
## Command to run 
```bash
go run cmd\apiStock\apiStock.go
  -apiKey string
        demo apikey
  -company string
        demo apikey (default "AAPL")
  -metric string
        Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric | UnderValuatedCompanies };(Required)
  -provider string
        Select the provider [ financialmodelingprep ] (Required)
  -yes string
        Populate companies Symbol (default "no")
```

```
go run cmd\apiStock\apiStock.go -provider financialmodelingprep -metric DiscountedCashFlow -company AAPL -apiKey demo
https://financialmodelingprep.com/api/v3/discounted-cash-flow/AAPL?apikey=demo
Symbol: AAPL , StockPrice: 121.21 , DiscountCashFlowValue: 123.37884618136223
```