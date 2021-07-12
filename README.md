# Dummy External Adapter

Useful for some basic testing and environment setup, this is a dead simple adapter for testing chainlink nodes and contracts.

## Endpoints

Default Port: `6060`

| Method | Endpoint                       | Description                                       |
| ------ | ------------------------------ | ------------------------------------------------- |
| GET    | /                              | Basic "is running" check                          |
| POST   | /random                        | Returns a random int withing 0-100                |
| POST   | /five                          | Returns 5                                         |
| POST   | /set_variable?var={your-value} | Allows you to set a variable to be later returned |
| POST   | /variable                      | Returns whatever variable you set                 |
