# Products 

## Developing

### Requirements

- Golang
- Java 14

### Services

- [ x ] product-service
- Service written in Go

- [ x ] discount-calculator-service
- Service written in Java

- [ x ] user-service
- Service written in Go

## Running

1. Run the following command:
```
docker-compose up
```
2. You may run on your browser or curl. There are some examples that you might test:
```
- http://localhost:8001/product

// or with user param
- http://localhost:8001/product?user=41597637-8c33-409f-a869-a2090e87ec78
```

```
curl --request GET \
  --url 'http://localhost:8001/product

// or with user param
curl --request GET \
  --url 'http://localhost:8001/product?user=41597637-8c33-409f-a869-a2090e87ec78'
```

3. For testing the birthday discount use the following user:
`
4a07fb31-d908-411b-949e-6ae3effbe60b
`