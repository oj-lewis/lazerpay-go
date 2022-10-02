# Go API client for Lazerpay crypto payment

This is a Go client for interacting with the lazerpay API for quick crypto payments for businesses

## Installation

```shell
    go get github.com/oj-lewis/lazerpay-go
```

## Usage

```Go
    import (
        "log"
        
        "github.com/oj-lewis/lazerpay-go"
    )


    func main() {
        // create a new lazerpay instance
        c, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
        if err != nil {
            log.Fatal(err)
        }
    }
```

## Availabe Methods on the sdk

**1**. **Payments 💵**
- Initialize Payment
- Verify Payment

**2**. **Payout 🏦**
- Crypto Payout 
- Bank Payout

**3**. **Swap ♻**
- Crypto Swap 
- Get Crypto swap amount out

**4**. **Payment Links 🔗**
- Create payment links 
- Get all payment links
- Update payment links
- Get a single payment link

## Payment

#### `Initialize Payment`
This describes how to initiate a payment foy your 
customer

```Go
import (
    "log"

    "github.com/oj-lewis/lazerpay-go"
    )


func main() {
     // create a new lazerpay instance
    lazerpay, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
    if err != nil {
        log.Fatal(err)
    }

    payment_tx := &PaymentOption{
	    Reference: "YOUR REFERENCE", 
	    CustomerName: "John Doe",         	
	    CustomerEmail: "johndoe@example.com",
	    Coin: "USDT", // DAI, BUSD, USDT or USDC
	    Currency: "USD", // USD, NGN, GBP, EUR
	    Amount: 100,
	    AcceptPartialPayment: true, // defaults to false	 	
    } 
    resp, err := lazerpay.Payment.Initialize(payment_tx)
    if err != nil {
        log.Fatal(err)
    }
} 
```


#### `Verify Payment`
This describes how to verify an initialized payment

```Go
import (
    "log"

    "github.com/oj-lewis/lazerpay-go"
    )


func main() {
     // create a new lazerpay instance
    lazerpay, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
    if err != nil {
        log.Fatal(err)
    }

    id := "the id generated from initializing the payment"
    
    resp, err := lazerpay.Payment.Verify(id)
    if err != nil {
        log.Fatal(err)
    }
} 
```


## Payout

#### `Crypto Payout`
This describes how to withdraw the crypto in thier lazerpay balance.

```Go
import (
    "log"

    "github.com/oj-lewis/lazerpay-go"
    )


func main() {
     // create a new lazerpay instance
    lazerpay, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
    if err != nil {
        log.Fatal(err)
    }

    payout_tx := &CrptoPayoutOptions{
	    Reference: "YOUR REFERENCE", 
	    Coin: "USDT", // DAI, BUSD, USDT or USDC
	    Amount: 100,
        Recipient: "THE CRYPTO ADDRESS OF THE RECIPIENT",
        Blockchain: "Binance Smart Chain",
	    	 	
    } 
    resp, err := lazerpay.Payout.Crypto(payout_tx)
    if err != nil {
        log.Fatal(err)
    }
} 
```

#### `Bank Payout`
This describes how to withdraw the crypto in thier lazerpay balance to their bank accounts.

```Go
import (
    "log"

    "github.com/oj-lewis/lazerpay-go"
    )


func main() {
     // create a new lazerpay instance
    lazerpay, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
    if err != nil {
        log.Fatal(err)
    }

    payout_tx := &BankOptions{
	    BankName: "Kuda Bank",
        BankCode: "KUDA BANK CODE",
        AccountName:"JOHN DOE",
        AccountNumber: "1234567890",
        Currency: "NGN",
        Country: "Nigeria",
        Default: true,
    } 

    // creates a new payout bank
    resp, err := lazerpay.Payout.Create(payout_tx)
    if err != nil {
        log.Fatal(err)
    }
} 
```
```Go
import (
    "log"

    "github.com/oj-lewis/lazerpay-go"
    )


func main() {
     // create a new lazerpay instance
    lazerpay, err := lazerpay.NewClient(LAZERPAY_PUBLIC_KEY, LAZERPAY_SECRET_KEY)
    if err != nil {
        log.Fatal(err)
    }

    initiate_tx := &PayoutOptions{
	    BankId: "",
        COIN: "USDT",
        Amount: 100,
    } 

    // initiates a new bank payout
    resp, err := lazerpay.Payout.Initiate(initiate_tx)
    if err != nil {
        log.Fatal(err)
    }
} 
```

## Author 

- Omiete John-lewis 
- Contact on twitter [@the_ojlewis](https://twitter.com/the_ojlewis)

## Contribution
Contributions are welcome if there is any bug🐛, you can a new issue or create a pull request