
# TwiML

The TwiML package wraps the functionality of [TwiML](https://www.twilio.com/docs/api/twiml).

TwiML objects implement the [HandlerFunc](http://golang.org/pkg/net/http/#HandlerFunc) interface which allows them to be passed directly to http.Handle.



## Forming TwiML







## Note on Caching

By default, TwiML structs are cached after they are rendered.
Due to how this caching is tracked, if you make an update to an inner TwiML block, it will not be cached.
If you make an edit to an inner block of TwiML after you have added it to the outer block, you must call `ClearCache()`.

```go
inner := new(twiml.DialTwiml)
inner.Number(NumberOpts{} "123")
response.Dial(DialOpts{}, inner)
// NOTE: this will alter "response" and add a Sip directive
inner.Sip(SipOpts{}, "sip:something")
```



