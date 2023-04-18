# Iris + GORM

Simple implementation of REST API with Iris framework and GORM

## Explaining test
- Use benchmarck tool in route /dogs
    - `autocannon http://localhost:4000/dogs`
- Populate database with 10 records
```curl
curl --location --request GET 'localhost:4000/dogs' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ralph",
    "breed":"Pastor",
    "age": 12,
    "isGoodBoy":true
}'
```

#### Benchmark
![bench](https://github.com/LeandroRezendeCoutinho/iris-gorm/blob/main/img/iris_gorm_bench.png)

#### Memory and CPU
![bench](https://github.com/LeandroRezendeCoutinho/iris-gorm/blob/main/img/iris_mem.png)

References: 
-https://www.iris-go.com/
- https://gorm.io/
