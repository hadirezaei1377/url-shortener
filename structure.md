- **main.go**: 
the main file that starts the HTTP server.

- **handler/url_handler.go**: 
where we define routes and handlers to receive and shorten links.

- **storage/storage.go**:
 Here we define the interfaces related to the storage of links.

- **storage/memory.go**:
 A simple implementation of the storage interface using memory to store links.

- **cache/cache.go**: 
Here we define interfaces related to cache.

- **cache/memory_cache.go**: 
A simple implementation of the cache interface using memory.

- **utils/shortener.go**: 
Contains link shortening algorithm.