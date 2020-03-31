# user-api caching written in GO

## Requirements
Interface
http POST/<key> UTF-8.
http GET/<key> response is 404 if no such key.

Stored keys expire after 30 minutes
Large load with an 80/20 ratio of read/writes.
