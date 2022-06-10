# Stream

`stream` is a Go library providing some utility functions on streams. This library implements *lazy* streams, 
meaning building the stream doesn't actually loop through it. Some functions, like `ForEach` or `Reduce` 
transform the stream into a finite state, and thus consume the stream. This makes it efficient to map, filter, 
reduce over a collection, since a single loop is needed for every operation. More unit tests are comming to 
get code coverage to 100%. If you cannot use a simple array as the stream's input, you can always define your own
implementation of `Stream` and pass it to this library.

## Caveats
this library's streams cannot be processed in multiple goroutines, or multiple times. A stream, as defined here, is single use.
Once the iteration is completed, it is no longer usable.
