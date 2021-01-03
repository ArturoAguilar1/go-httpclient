package gohttp

type Mock struct {
    Method string
    Url string
    RequestBody string

    Error error
    ResponseBody error
    ResponseStatusCode int
}
