# Go Client Library for GoCardless Pro Api

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
    gocardless "github.com/gocardless/gocardless-pro-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the gocardless-pro-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```
go get -u github.com/gocardless/gocardless-pro-go/v1.0.0
```

## Initializing the client

The client is initialised with an access token, and is configured to use GoCardless live environment by default:

```go
    token := "your_access_token"
    client, err := gocardless.New(token)
```


Optionally, the client can be customised with endpoint, for ex: sandbox environment
```go
    opts := gocardless.WithEndpoint(gocardless.SandboxEndpoint)
    client, err := gocardless.New(token, opts)
```

the client can also be initialised with a customised http client, for ex;
```go
    customHttpClient := &http.Client{
        Timeout: time.Second * 10,
    }
    opts := gocardless.WithClient(customHttpClient)
    client, err := gocardless.New(token, opts)
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
    customerListIterator := service.Customers.All(ctx, customerListParams)
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
    customerCreateParams := CustomerCreateParams{}
    customerCreateParams.AddressLine1 = "9 Acer Gardens"
    customerCreateParams.City = "Birmingham"
    customerCreateParams.CountryCode = "GB"
    customerCreateParams.Email = "bbr@example.xom"
    customerCreateParams.FamilyName = "Rodgriguez"
    customerCreateParams.GivenName = "Bender Bending"
    customerCreateParams.PostalCode = "B4 7NJ"

    customer, err := client.Customers.Create(ctx, customerCreateParams)
```

### Updating Resources

Resources can be updates with the `Update` method:

```go
    ctx := context.TODO()
    customerUpdateParams := CustomerUpdateParams{}
    customerUpdateParams.GivenName = "New name"

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
    customersCreateResult, err := service.Customers.Create(ctx, customerCreateParams, requestOption)
```

### Setting custom headers

You shouldn't generally need to customise the headers sent by the library, but you wish to
in some cases (for example if you want to send an `Accept-Language` header when [creating a mandate PDF](https://developer.gocardless.com/api-reference/#mandate-pdfs-create-a-mandate-pdf)).
```go
    headers := make(map[string]string)
    headers["Accept-Language"] = "fr"
    requestOption := gocardless.WithHeaders(headers)
    customersCreateResult, err := service.Customers.Create(ctx, customerCreateParams, requestOption)
```

Custom headers you specify will override any headers generated by the library itself (for
example, an `Idempotency-Key` header with a randomly-generated value or one you've configured
manually). Custom headers always take precedence.
```go
    requestOption := gocardless.WithIdempotencyKey("test-idemptency-key-123")
    customersCreateResult, err := service.Customers.Create(ctx, customerCreateParams, requestOption)
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

## Compatibility

This library requires go 1.16 and above.

## Documentation

TODO

## Contributing

This client is auto-generated from Crank, a toolchain that we hope to soon open source. Issues should for now be reported on this repository.  __Please do not modify the source code yourself, your changes will be overridden!__
