## Latch Golang SDK 

This is a **NOT** official SDK for Latch for the [Go](https://golang.org/) programming language

* **Requirements**

    - You will need [Go installed](https://golang.org/doc/install). I am using version 1.8.3 on Ubuntu Linux
    - Read the official [Latch documentation](https://latch.elevenpaths.com/www/developers/doc_api)
    - In order to use this library, you will also need a [Latch account](https://latch.elevenpaths.com/), from there you will have access to your Application Id/User Id and the Secret values.
    - You will need to download this module:


        go get github.com/tuxotron/latch-sdk-go

* **Code Example**

    

    import "github.com/tuxotron/latch-sdk-go"

    credentials := latch.Credentials{"AppId", "Secret"}
    
    application := latch.NewLatchApplication(credentials)

    // Pairing using token
    
    application.PairWithToken(your-token-here)
    
    // Unpair
    
    application.Unpair(applicationId)
    
    // Check status
    
    application.Status(applicationId, false, false))
    
    
There are several methods (Application lock, unlock, history, User API, etc) I could not test myself because they require a *Gold* or *Platinum* subscription. 

**TODO**

Serialize server responses into a JSON object