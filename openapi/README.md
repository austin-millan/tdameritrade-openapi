# Go API client for openapi

TD Ameritrade API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.tdameritrade.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAndTradingApi* | [**CancelOrder**](docs/AccountsAndTradingApi.md#cancelorder) | **Delete** /accounts/{accountId}/orders/{orderId} | Cancel Order
*AccountsAndTradingApi* | [**DeleteSavedOrder**](docs/AccountsAndTradingApi.md#deletesavedorder) | **Delete** /accounts/{accountId}/savedorders/{savedOrderId} | Delete Saved Order
*AccountsAndTradingApi* | [**GetAccount**](docs/AccountsAndTradingApi.md#getaccount) | **Get** /accounts/{accountId} | Get Account
*AccountsAndTradingApi* | [**GetAccounts**](docs/AccountsAndTradingApi.md#getaccounts) | **Get** /accounts | Get Accounts
*AccountsAndTradingApi* | [**GetOrder**](docs/AccountsAndTradingApi.md#getorder) | **Get** /accounts/{accountId}/orders/{orderId} | Get Order
*AccountsAndTradingApi* | [**GetOrdersByPath**](docs/AccountsAndTradingApi.md#getordersbypath) | **Get** /accounts/{accountId}/orders | Get Orders By Path
*AccountsAndTradingApi* | [**GetOrdersByQuery**](docs/AccountsAndTradingApi.md#getordersbyquery) | **Get** /orders | Get Orders By Query
*AccountsAndTradingApi* | [**GetSavedOrder**](docs/AccountsAndTradingApi.md#getsavedorder) | **Get** /accounts/{accountId}/savedorders/{savedOrderId} | Get Saved Order
*AccountsAndTradingApi* | [**GetSavedOrdersByPath**](docs/AccountsAndTradingApi.md#getsavedordersbypath) | **Get** /accounts/{accountId}/savedorders | Get Saved Orders by Path
*AccountsAndTradingApi* | [**PlaceOrder**](docs/AccountsAndTradingApi.md#placeorder) | **Post** /accounts/{accountId}/orders | Place Order
*AccountsAndTradingApi* | [**ReplaceOrder**](docs/AccountsAndTradingApi.md#replaceorder) | **Put** /accounts/{accountId}/orders/{orderId} | Replace Order
*AccountsAndTradingApi* | [**ReplaceSavedOrder**](docs/AccountsAndTradingApi.md#replacesavedorder) | **Put** /accounts/{accountId}/savedorders/{savedOrderId} | Replace Saved Order
*AccountsAndTradingApi* | [**SaveOrder**](docs/AccountsAndTradingApi.md#saveorder) | **Post** /accounts/{accountId}/savedorders | Create Saved Order
*AuthenticationApi* | [**PostAccessToken**](docs/AuthenticationApi.md#postaccesstoken) | **Post** /oath2/token | Post Access Token
*InstrumentsApi* | [**GetInstrument**](docs/InstrumentsApi.md#getinstrument) | **Get** /instruments/{accountId} | Search for instrument and fundamental data
*InstrumentsApi* | [**SearchInstruments**](docs/InstrumentsApi.md#searchinstruments) | **Get** /instruments | Search for instrument and fundamental data
*MarketHoursApi* | [**GetHours**](docs/MarketHoursApi.md#gethours) | **Get** /marketdata/hours | Operating hours of markets
*MarketHoursApi* | [**GetMarketHours**](docs/MarketHoursApi.md#getmarkethours) | **Get** /marketdata/{market}/hours | Operating hours of markets
*MoversApi* | [**GetMovers**](docs/MoversApi.md#getmovers) | **Get** /marketdata/{index}/movers | Retrieve mover information by index symbol, direction type and change
*OptionChainsApi* | [**GetOptionChains**](docs/OptionChainsApi.md#getoptionchains) | **Get** /marketdata/chains | Get Option Chains for optionable symbols
*PriceHistoryApi* | [**GetPriceHistory**](docs/PriceHistoryApi.md#getpricehistory) | **Get** /marketdata/{symbol}/pricehistory | Historical price data for charts
*QuotesApi* | [**GetQuote**](docs/QuotesApi.md#getquote) | **Get** /marketdata/{symbol}/quotes | Search for instrument and fundamental data
*QuotesApi* | [**GetQuotes**](docs/QuotesApi.md#getquotes) | **Get** /marketdata/quotes | Search for instrument and fundamental data
*TransactionHistoryApi* | [**GetTransaction**](docs/TransactionHistoryApi.md#gettransaction) | **Get** /accounts/{accountId}/transactions/{transactionId} | APIs to access transaction history on the account
*TransactionHistoryApi* | [**GetTransactions**](docs/TransactionHistoryApi.md#gettransactions) | **Get** /accounts/{accountId}/transactions | APIs to access transaction history on the account
*UserInfoPreferencesApi* | [**GetPreferences**](docs/UserInfoPreferencesApi.md#getpreferences) | **Get** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences
*UserInfoPreferencesApi* | [**GetStreamerSubscriptionKeys**](docs/UserInfoPreferencesApi.md#getstreamersubscriptionkeys) | **Get** /userprincipals/streamersubscriptionkeys | APIs to access user-authorized accounts and their preferences
*UserInfoPreferencesApi* | [**GetUserPrincipals**](docs/UserInfoPreferencesApi.md#getuserprincipals) | **Get** /userprincipals | APIs to access user-authorized accounts and their preferences
*UserInfoPreferencesApi* | [**UpdatePreferences**](docs/UserInfoPreferencesApi.md#updatepreferences) | **Put** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences
*WatchlistApi* | [**CreateWatchlist**](docs/WatchlistApi.md#createwatchlist) | **Post** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**DeleteWatchlist**](docs/WatchlistApi.md#deletewatchlist) | **Delete** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**GetWatchlist**](docs/WatchlistApi.md#getwatchlist) | **Get** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**GetWatchlistMultipleAccounts**](docs/WatchlistApi.md#getwatchlistmultipleaccounts) | **Get** /accounts/watchlists | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**GetWatchlistSingleAccount**](docs/WatchlistApi.md#getwatchlistsingleaccount) | **Get** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**ReplaceWatchlist**](docs/WatchlistApi.md#replacewatchlist) | **Put** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
*WatchlistApi* | [**UpdateWatchlist**](docs/WatchlistApi.md#updatewatchlist) | **Patch** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist


## Documentation For Models

 - [Account](docs/Account.md)
 - [Bond](docs/Bond.md)
 - [CandleList](docs/CandleList.md)
 - [CandleListCandles](docs/CandleListCandles.md)
 - [CashAccount](docs/CashAccount.md)
 - [CashAccountCurrentBalances](docs/CashAccountCurrentBalances.md)
 - [CashAccountInitialBalances](docs/CashAccountInitialBalances.md)
 - [CashAccountOrderStrategies](docs/CashAccountOrderStrategies.md)
 - [CashAccountPositions](docs/CashAccountPositions.md)
 - [CashEquivalent](docs/CashEquivalent.md)
 - [CreateWatchlist](docs/CreateWatchlist.md)
 - [CreateWatchlistInstrument](docs/CreateWatchlistInstrument.md)
 - [CreateWatchlistWatchlistItems](docs/CreateWatchlistWatchlistItems.md)
 - [EasObject](docs/EasObject.md)
 - [Equity](docs/Equity.md)
 - [Etf](docs/Etf.md)
 - [Execution](docs/Execution.md)
 - [ExecutionExecutionLegs](docs/ExecutionExecutionLegs.md)
 - [ExpirationDate](docs/ExpirationDate.md)
 - [FixedIncome](docs/FixedIncome.md)
 - [Forex](docs/Forex.md)
 - [Fundamental](docs/Fundamental.md)
 - [FundamentalData](docs/FundamentalData.md)
 - [FundamentalFundamental](docs/FundamentalFundamental.md)
 - [Future](docs/Future.md)
 - [FutureOptions](docs/FutureOptions.md)
 - [Hours](docs/Hours.md)
 - [Index](docs/Index.md)
 - [Instrument](docs/Instrument.md)
 - [MarginAccount](docs/MarginAccount.md)
 - [MarginAccountCurrentBalances](docs/MarginAccountCurrentBalances.md)
 - [MarginAccountInitialBalances](docs/MarginAccountInitialBalances.md)
 - [Mover](docs/Mover.md)
 - [MutualFund](docs/MutualFund.md)
 - [Option](docs/Option.md)
 - [OptionChain](docs/OptionChain.md)
 - [OptionChainUnderlying](docs/OptionChainUnderlying.md)
 - [OptionDeliverables](docs/OptionDeliverables.md)
 - [OptionOptionDeliverables](docs/OptionOptionDeliverables.md)
 - [Order](docs/Order.md)
 - [OrderCancelTime](docs/OrderCancelTime.md)
 - [OrderChildOrderStrategies](docs/OrderChildOrderStrategies.md)
 - [OrderOrderLegCollection](docs/OrderOrderLegCollection.md)
 - [Preferences](docs/Preferences.md)
 - [SavedOrder](docs/SavedOrder.md)
 - [SavedOrderChildOrderStrategies](docs/SavedOrderChildOrderStrategies.md)
 - [SubscriptionKey](docs/SubscriptionKey.md)
 - [SubscriptionKeyKeys](docs/SubscriptionKeyKeys.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionTransactionItem](docs/TransactionTransactionItem.md)
 - [TransactionTransactionItemInstrument](docs/TransactionTransactionItemInstrument.md)
 - [Underlying](docs/Underlying.md)
 - [UpdatePreferences](docs/UpdatePreferences.md)
 - [UpdateWatchlist](docs/UpdateWatchlist.md)
 - [UpdateWatchlistWatchlistItems](docs/UpdateWatchlistWatchlistItems.md)
 - [UserPrincipal](docs/UserPrincipal.md)
 - [UserPrincipalAccounts](docs/UserPrincipalAccounts.md)
 - [UserPrincipalAuthorizations](docs/UserPrincipalAuthorizations.md)
 - [UserPrincipalPreferences](docs/UserPrincipalPreferences.md)
 - [UserPrincipalQuotes](docs/UserPrincipalQuotes.md)
 - [UserPrincipalStreamerInfo](docs/UserPrincipalStreamerInfo.md)
 - [UserPrincipalStreamerSubscriptionKeys](docs/UserPrincipalStreamerSubscriptionKeys.md)
 - [Watchlist](docs/Watchlist.md)
 - [WatchlistInstrument](docs/WatchlistInstrument.md)
 - [WatchlistWatchlistItems](docs/WatchlistWatchlistItems.md)


## Documentation For Authorization



## Bearer

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## Author

austin.millan@gmail.com
