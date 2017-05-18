# Memory lane tasks


## Local Setup

Run a postgres containers
docker run --rm --name tasks -p 5432:5432 -e POSTGRES_PASSWORD=mysecrettasks -d postgres


docker run -it --rm --link tasks:postgres -e PGPASSWORD=mysecrettasks postgres psql -h postgres -U postgres
