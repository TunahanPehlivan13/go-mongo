# go-mongo

An example of go, mongodb and in-memory database usage together.

Running version of this project under <a href="http://go-mongo-getir.herokuapp.com/">this URL</a>

You can find out the setup instructions in `Makefile`

#### Structure:
4 Domain layers:

- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## API:

### POST /records

Fetchs records

##### Example Input:
```
{
    "startDate": "2015-11-29T14:44:43.114Z",
    "endDate": "2019-11-29T14:44:45.114Z",
    "minCount": 169,
    "maxCount": 189999
} 
```

### POST /in-memory

Post key and value

##### Example Input:
```
{
    "key": "key",
    "value": "value"
}
```

### GET /in-memory?key=sample-key

Fetch value by key
