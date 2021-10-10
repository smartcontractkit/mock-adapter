# Dummy External Adapter

Pull from `public.ecr.aws/chainlink/mock-adapter`

Useful for some basic testing and environment setup, this is a dead simple adapter for testing chainlink nodes and contracts.

## Endpoints

Default Port: `6060`

| Method      | Endpoint                       | Description                                       |
| ----------- | ------------------------------ | ------------------------------------------------- |
| GET/POST    | /                              | Basic "is running" check                          |
| GET/POST    | /random                        | Returns a random int ranging from 0 to 100        |
| GET/POST    | /five                          | Returns 5                                         |
| POST        | /set_variable?var={your-value} | Allows you to set an integer to be later returned |
| GET/POST    | /variable                      | Returns whatever integer you set                  |
| POST   | /set_json_variable | Allows you to set a json body to be returned later |
| GET/POST   | /json_variable | returns whatever json you set |
