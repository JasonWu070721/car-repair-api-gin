# Car Repair API

This API will help you store information about cars stored in your factory (repair shop).

## Quick start

### Docker

#### Run on Dockerfile

```shell script

docker build -t car_repair_api .
docker network create -d bridge car-repair-network

docker run -p 6379:6379 --name car_repair_redis --network car-repair-network -d redis

docker run --name car_repair_postgresql --network car-repair-network -e POSTGRES_PASSWORD=your_password -e POSTGRES_USER=postgresql -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres

docker run --name car_repair_api  --network car-repair-network -p 443:443 -d car_repair_api
```

#### Run on docker-compose

```shell script
docker-compose down
docker-compose up
```

## TODO

- [x] Support Swagger of OpenAPI
- [x] Support Docker
- [x] Create customer API
- [x] Create user API
- [x] Support JWT
- [x] Support HTTPS
- [x] Support Redis cache
- [x] Support PostgreSQL
- [x] Support AWS RDS
- [x] Support AWS EC2 deployment
- [x] Support ORM(Object Relational Mapping)
- [x] Create Python API Test
- [x] Support golang unit test
- [x] Support Gitlab CI/CD
- [ ] Create CSRF token
- [ ] XSS (Cross-Site Scripting) protection
- [ ] C10K test
- [ ] Create vehicle API
- [ ] Create quotation API
- [ ] Create material API
- [ ] Create service API
- [ ] Injection protection
- [ ] Support K8S
