# UniTusk

A generic package for markov chain bots.

## Instructions

First, get the package:

```
go get github.com/wallnutkraken/unitusk
```

Afterwards, you can call `unitusk.New()` to create a Hivemind object. This object does not have any EndpointProviders.
EndpointProvider is an interface that the package uses for gathering
and sending messages. How it's implemented is up to you, but generally:

- There is no need to call `Send()` in your own code
- Queue should just expose a SendQueue object
- All errors must be self contained in an error slice and be accessible through `Errors()`