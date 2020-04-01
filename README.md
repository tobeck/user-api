<p align="left">
  <a href="https://github.com/tobeck/user-api"><img alt="Github Actions status" src="https://github.com/tobeck/user-api/workflows/build/badge.svg"></a>
</p>

# user-api caching written in GO

## Requirements
* http POST/<key> UTF-8.
* http GET/<key> response is 404 if no such key.
* Cached users expire after 30 minutes.
* Load with an 80/20 ratio of read/writes.
