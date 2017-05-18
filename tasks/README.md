# Memory lane tasks


## Local Setup

*Run a postgres containers*

        docker run --rm --name tasks -p 5432:5432 -e POSTGRES_PASSWORD=mysecrettasks -d postgres


*Connect to that postgres*

        docker run -it --rm --link tasks:thatpostgres -e PGPASSWORD=mysecrettasks postgres psql -h postgres -U postgres

*Run the API connecting to the database*

```
docker run -it --rm --link tasks:thatpostgres -p 8000:8000 \
    -e APP_DB_USERNAME=postgres \
    -e APP_DB_PASSWORD=mysecrettasks \
    -e APP_DB_NAME=tasks \
    -e APP_DB_HOST=thatpostgres \
    quay.io/ipedrazas/memorylane-tasks
```