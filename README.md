# microsoft-graph-client-go
**Unoffficial** Microsoft Graph Client for Go

I needed a client so I started writing one.
See example in the example folder.

This client is super super alpha, hence it will be v0 for the forseable future. Do not expect the json to be deserialized into structs. If you look at the source there are no structs to represent the data structures graph api exposes. **The api will return the raw json and in its current state the consumer is expected to provide the struct to be hydrated.**

## current focus
filling out the query options and testing it with client secret credentials. Allowing full CRUD.

## future version focus

v1 should support all of the above, in addition to the various authentication options available in MSAL (which itself is still alpha)

v2 should start filling in the various data structures for QOL. 