1. Client which need to process by a chain. IE: a request need to validate data, check IP, log, decode, check authorization/user’s credentials, check caches.
2. Handler interface which defines generals work must be done by client.
    + Execute()
    + SetNext(Handler)
3. Implementation structs.
