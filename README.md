# apiStock (WIP)

In this early version we can get the companies that are underValuated following the discounted cash flow value. 
We'll can get the companies under a specific sector.  

## Options to run 
```bash
go run cmd\apiStock\apiStock.go
  -apiKey string
        financialmodelingprep apikey (Required)
  -beta string
        Company Beta (default "1")
  -company string
        MSFT (default "MSFT")
  -marketCap string
        Market Capitalization  (default "1000000000")
  -metric string
        Metric { DiscountedCashFlow | UnderValuatedCompanies };(Required)
  -percentageOfGrowth float
        Set the minimum growth tha the company should have (default 10)
  -provider string
        Select the provider [ financialmodelingprep ] (Required) (default "financialmodelingprep")
  -sector string
        Company Sector: Consumer Cyclical - Energy - Technology - Industrials - Financial Services - Basic Materials - Communication Services - Consumer Defensive - Healthcare - Real Estate - Utilities - Industrial Goods - Financial - Services - Cong
lomerates  (default "Technology")
  -yes string
        Populate companies Symbol (default "no")

```

## How to use it

First off we must create a financialmodelingprep apikey, to do that we just need to create a user on [financialmodelingprep/developers](https://financialmodelingprep.com/developer/docs/)

After to have our apiKey we'll be able to use 2 types of metrics. 

### UnderValuatedCompanies
Using the sector and some trait we'll receive the groups of companies that are underValuated using DiscountCashFlow

```
go run cmd/apiStock/apiStock.go -provider financialmodelingprep -metric UnderValuatedCompanies -beta 1 -marketCap 1000000000 -percentageOfGrowth 10 -sector Financial -apiKey xxxxxxxxxxxxxxxxxxxx

Symbol: WSFS, StockPrice: 50.37 , DiscountCashFlowValue: 80.68206887303607,  Change : 60.18 %

Symbol: WTFC, StockPrice: 76.12 , DiscountCashFlowValue: 118.65331366389559,  Change : 55.88 %

Symbol: WU, StockPrice: 25.15 , DiscountCashFlowValue: 28.594283646888567,  Change : 13.69 %



```
Sectors: Consumer Cyclical - Energy - Technology - Industrials - Financial Services - Basic Materials - Communication Services - Consumer Defensive - Healthcare - Real Estate - Utilities - Industrial Goods - Financial - Services - Cong

### DiscountedCashFlow 
Using the company symbol we'll receive the stock values, and the  discounted cash flow value

```
go run cmd/apiStock/apiStock.go -metric DiscountedCashFlow -company MSFT -apiKey xxxxxxxxxxxx

Number of companies that we are checking:  1

Symbol: MSFT, StockPrice: 242.35 , DiscountCashFlowValue: 244.1170180918912,
```
