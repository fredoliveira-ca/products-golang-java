# Products 
![services](/assets/products-golan-java.png)

## Developing

### Requirements

- Golang
- Java 14

### Services

- [x] product-service
  - Service written in Go

- [x] discount-calculator-service
  - Service written in Java

- [x] user-service
  - Service written in Go

## Running
> If you prefer, you can access the environment in the Google Cloud which I've created in order to test the commands below.
> You can access that through this link: http://35.223.114.52:8001/products
1. Run the following command:
> :warning: Check if the files 'schema.sql' and 'data.sql' are in the same folder that you will run the command. It will create the tables and the data for testing.
```
docker-compose up
```

2. You may run on your browser or curl. There are some examples that you might test:
```
- http://localhost:8001/products
// or with user param
- http://localhost:8001/products?user=41597637-8c33-409f-a869-a2090e87ec78
```

```
curl --request GET \
  --url 'http://localhost:8001/products
// or with user param
curl --request GET \
  --url 'http://localhost:8001/products?user=41597637-8c33-409f-a869-a2090e87ec78'
```

3. For testing the birthday discount you may use the following user:
`
4a07fb31-d908-411b-949e-6ae3effbe60b
`
> :warning: **According to the script, you're adding this user with birthday date equal today!**: See this on line 40 in the data.sql file! If you want to test it afterwards, please update the user.
