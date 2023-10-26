# Go Client Library for GoCardless Pro API [![CircleCI](https://circleci.com/gh/gocardless/gocardless-pro-go-template/tree/master.svg?style=svg&circle-token=68c31e704d9b0020a5f42b4b89b0a77a17bdac6c)](https://circleci.com/gh/gocardless/gocardless-pro-go-template/tree/master)

This library provides a simple wrapper around the [GoCardless API](http://developer.gocardless.com/api-reference).

- ["Getting started" guide](https://developer.gocardless.com/getting-started/api/introduction/?lang=go) with copy and paste Go code samples
- [API Reference](https://developer.gocardless.com/api-reference/2015-07-06)

## Getting started

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
go mod tidy
```

Then, reference gocardless-pro-go in a Go program with `import`:
``` go
import (
    gocardless "github.com/gocardless/gocardless-pro-go/v3"
)
```


Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the gocardless-pro-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```
go get -u github.com/gocardless/gocardless-pro-go@v3.6.0
```

## Initializing the client

The client is initialised with an access token, and is configured to use GoCardless live environment by default:

```go
    token := "your_access_token"
    config, err := gocardless.NewConfig(token)
    if err != nil {
        fmt.Printf("got err in initialising config: %s", err.Error())
        return
    }
    client, err := gocardless.New(config)
    if err != nil {
		fmt.Printf("error in initialisating client: %s", err.Error())
		return
	}
```


Optionally, the client can be customised with endpoint, for ex: sandbox environment
```go
    config, err := gocardless.NewConfig(token, gocardless.WithEndpoint(gocardless.SandboxEndpoint))
    if err != nil {
        fmt.Printf("got err in initialising config: %s", err.Error())
        return
    }
    client, err := gocardless.New(config)
    if err != nil {
		fmt.Printf("error in initialisating client: %s", err.Error())
		return
	}
```

the client can also be initialised with a customised http client, for ex;
```go
    customHttpClient := &http.Client{
        Timeout: time.Second * 10,
    }
    config, err := gocardless.NewConfig(token, gocardless.WithClient(customHttpClient))
    if err != nil {
        fmt.Printf("got err in initialising config: %s", err.Error())
        return
    }
    client, err := gocardless.New(config)
    if err != nil {
		fmt.Printf("error in initialisating client: %s", err.Error())
		return
	}
```

## Examples 

### Fetching resources

To fetch single item, use the `Get` method

```go
    ctx := context.TODO()
    customer, err := client.Customers.Get(ctx, "CU123")
```

### Listing resources

To fetch items in a collection, there are two options:

* Fetching items one page at a time using `List` method: 

```go
    ctx := context.TODO()
    customerListParams := gocardless.CustomerListParams{}
    customers, err := client.Customers.List(ctx, customerListParams)
    for _, customer := range customers.Customers {
        fmt.Printf("customer: %v", customer)
    }
    cursor := customers.Meta.Cursors.After
    customerListParams.After = cursor
    nextPageCustomers, err := client.Customers.List(ctx, customerRemoveParams)
```

* Iterating through all of the items using a `All` method to get a lazily paginated list.
  `All` will deal with making extra API requests to paginate through all the data for you:

```go
    ctx := context.TODO()
    customerListParams := gocardless.CustomerListParams{}
    customerListIterator := client.Customers.All(ctx, customerListParams)
    for customerListIterator.Next() {
        customerListResult, err := customerListIterator.Value(ctx)
        if err != nil {
            fmt.Printf("got err: %s", err.Error())
        } else {
            fmt.Printf("customerListResult is %v", customerListResult)
        }
    }
```

### Creating resources

Resources can be created with the `Create` method:

```go
    ctx := context.TODO()
    customerCreateParams := gocardless.CustomerCreateParams{
        AddressLine1: "9 Acer Gardens"
        City:         "Birmingham",
        PostalCode:   "B4 7NJ",
        CountryCode:  "GB",
        Email:        "bbr@example.com",
        GivenName:    "Bender Bending",
        FamilyName:   "Rodgriguez",
    }

    customer, err := client.Customers.Create(ctx, customerCreateParams)
```

### Updating Resources

Resources can be updates with the `Update` method:

```go
    ctx := context.TODO()
    customerUpdateParams := CustomerUpdateParams{
        GivenName: "New Name",
    }

    customer, err := client.Customers.Update(ctx, "CU123", customerUpdateParams)
```

### Removing Resources

Resources can be removed with the `Remove` method:

```go
    ctx := context.TODO()
    customerRemoveParams := CustomerRemoveParams{}

    customer, err := client.Customers.Remove(ctx, "CU123", customerRemoveParams)
``` 

### Retrying requests

The library will attempt to retry most failing requests automatically (with the exception of those which are not safe to retry).

`GET` requests are considered safe to retry, and will be retried automatically. Certain `POST` requests are made safe to retry by the use of an idempotency key, generated automatically by the library, so we'll automatically retry these too. Currently its retried for 3 times. If you want to override this behaviour(for example, to provide your own retry mechanism), then you can use the `WithoutRetries`.  This will not retry and return response object.

```go
    requestOption := gocardless.WithoutRetries()
    customersCreateResult, err := client.Customers.Create(ctx, customerCreateParams, requestOption)
```

### Setting custom headers

You shouldn't generally need to customise the headers sent by the library, but you wish to
in some cases (for example if you want to send an `Accept-Language` header when [creating a mandate PDF](https://developer.gocardless.com/api-reference/#mandate-pdfs-create-a-mandate-pdf)).
```go
    headers := make(map[string]string)
    headers["Accept-Language"] = "fr"
    requestOption := gocardless.WithHeaders(headers)
    customersCreateResult, err := client.Customers.Create(ctx, customerCreateParams, requestOption)
```

Custom headers you specify will override any headers generated by the library itself (for
example, an `Idempotency-Key` header with a randomly-generated value or one you've configured
manually). Custom headers always take precedence.
```go
    requestOption := gocardless.WithIdempotencyKey("test-idemptency-key-123")
    customersCreateResult, err := client.Customers.Create(ctx, customerCreateParams, requestOption)
```

### Handling webhooks

GoCardless supports webhooks, allowing you to receive real-time notifications when things happen in your account, so you can take automatic actions in response, for example:

* When a customer cancels their mandate with the bank, suspend their club membership
* When a payment fails due to lack of funds, mark their invoice as unpaid
* When a customer’s subscription generates a new payment, log it in their “past payments” list

The client allows you to validate that a webhook you receive is genuinely from GoCardless, and to parse it into Event objects(defined in event_service.go) which are easy to work with:

```go
    http.HandleFunc("/webhookEndpoint", func(w http.ResponseWriter, r *http.Request) {
        wh, err := NewWebhookHandler("secret", EventHandlerFunc(func(event gocardless.Event) error {
		    parseEvent(event)
        }))
        wh.ServeHTTP(w,r)
    })

    func parseEvents(event Event) {
        // work through list of events
    }
``` 

### Error Handling

When the library returns an `error` defined by us rather than the stdlib, it can be converted into a `gocardless.APIError` using `errors.As`:

```go
    billingRequest, err := client.BillingRequests.Create(ctx, billingRequestCreateParams)
	if err != nil {
		var apiErr *gocardless.APIError
		if errors.As(err, &apiErr) {
			fmt.Printf("got err: %v", apiErr.Message)
		}
		return nil, err
	}
```

## Compatibility

This library requires go 1.20 and above.

## Documentation

TODO

## Contributing

This client is auto-generated from Crank, a toolchain that we hope to soon open source. Issues should for now be reported on this repository.  __Please do not modify the source code yourself, your changes will be overridden!__
