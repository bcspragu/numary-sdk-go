# Go API client for Numary

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.6.0
- Package version: v1.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/numary/go-ledgerclient
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ledgerclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AddMetadataToAccount**](docs/AccountsApi.md#addmetadatatoaccount) | **Post** /{ledger}/accounts/{address}/metadata | Add metadata to an account.
*AccountsApi* | [**CountAccounts**](docs/AccountsApi.md#countaccounts) | **Head** /{ledger}/accounts | Count the accounts from a ledger.
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /{ledger}/accounts/{address} | Get account by its address.
*AccountsApi* | [**ListAccounts**](docs/AccountsApi.md#listaccounts) | **Get** /{ledger}/accounts | List accounts from a ledger.
*BalancesApi* | [**GetBalances**](docs/BalancesApi.md#getbalances) | **Get** /{ledger}/balances | Get the balances from a ledger&#39;s account
*BalancesApi* | [**GetBalancesAggregated**](docs/BalancesApi.md#getbalancesaggregated) | **Get** /{ledger}/aggregate/balances | Get the aggregated balances from selected accounts
*MappingApi* | [**GetMapping**](docs/MappingApi.md#getmapping) | **Get** /{ledger}/mapping | Get the mapping of a ledger.
*MappingApi* | [**UpdateMapping**](docs/MappingApi.md#updatemapping) | **Put** /{ledger}/mapping | Update the mapping of a ledger.
*ScriptApi* | [**RunScript**](docs/ScriptApi.md#runscript) | **Post** /{ledger}/script | Execute a Numscript.
*ServerApi* | [**GetInfo**](docs/ServerApi.md#getinfo) | **Get** /_info | Show server information.
*StatsApi* | [**ReadStats**](docs/StatsApi.md#readstats) | **Get** /{ledger}/stats | Get Stats
*TransactionsApi* | [**AddMetadataOnTransaction**](docs/TransactionsApi.md#addmetadataontransaction) | **Post** /{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID.
*TransactionsApi* | [**CountTransactions**](docs/TransactionsApi.md#counttransactions) | **Head** /{ledger}/transactions | Count the transactions from a ledger.
*TransactionsApi* | [**CreateTransaction**](docs/TransactionsApi.md#createtransaction) | **Post** /{ledger}/transactions | Create a new transaction to a ledger.
*TransactionsApi* | [**CreateTransactions**](docs/TransactionsApi.md#createtransactions) | **Post** /{ledger}/transactions/batch | Create a new batch of transactions to a ledger.
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /{ledger}/transactions/{txid} | Get transaction from a ledger by its ID.
*TransactionsApi* | [**ListTransactions**](docs/TransactionsApi.md#listtransactions) | **Get** /{ledger}/transactions | List transactions from a ledger.
*TransactionsApi* | [**RevertTransaction**](docs/TransactionsApi.md#reverttransaction) | **Post** /{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID.


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountWithVolumesAndBalances](docs/AccountWithVolumesAndBalances.md)
 - [Config](docs/Config.md)
 - [ConfigInfo](docs/ConfigInfo.md)
 - [ConfigInfoResponse](docs/ConfigInfoResponse.md)
 - [Contract](docs/Contract.md)
 - [CreateTransaction400Response](docs/CreateTransaction400Response.md)
 - [CreateTransaction409Response](docs/CreateTransaction409Response.md)
 - [CreateTransactions400Response](docs/CreateTransactions400Response.md)
 - [Cursor](docs/Cursor.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GetAccount200Response](docs/GetAccount200Response.md)
 - [GetAccount400Response](docs/GetAccount400Response.md)
 - [GetBalances200Response](docs/GetBalances200Response.md)
 - [GetBalances200ResponseCursor](docs/GetBalances200ResponseCursor.md)
 - [GetBalances200ResponseCursorAllOf](docs/GetBalances200ResponseCursorAllOf.md)
 - [GetBalancesAggregated200Response](docs/GetBalancesAggregated200Response.md)
 - [GetBalancesAggregated400Response](docs/GetBalancesAggregated400Response.md)
 - [GetTransaction400Response](docs/GetTransaction400Response.md)
 - [GetTransaction404Response](docs/GetTransaction404Response.md)
 - [LedgerStorage](docs/LedgerStorage.md)
 - [ListAccounts200Response](docs/ListAccounts200Response.md)
 - [ListAccounts200ResponseCursor](docs/ListAccounts200ResponseCursor.md)
 - [ListAccounts200ResponseCursorAllOf](docs/ListAccounts200ResponseCursorAllOf.md)
 - [ListAccounts400Response](docs/ListAccounts400Response.md)
 - [ListTransactions200Response](docs/ListTransactions200Response.md)
 - [ListTransactions200ResponseCursor](docs/ListTransactions200ResponseCursor.md)
 - [ListTransactions200ResponseCursorAllOf](docs/ListTransactions200ResponseCursorAllOf.md)
 - [Mapping](docs/Mapping.md)
 - [MappingResponse](docs/MappingResponse.md)
 - [Posting](docs/Posting.md)
 - [Script](docs/Script.md)
 - [ScriptResult](docs/ScriptResult.md)
 - [Stats](docs/Stats.md)
 - [StatsResponse](docs/StatsResponse.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionData](docs/TransactionData.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [Transactions](docs/Transactions.md)
 - [TransactionsResponse](docs/TransactionsResponse.md)
 - [Volume](docs/Volume.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



