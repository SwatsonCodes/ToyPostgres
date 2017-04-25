this is a really simple docker image that inserts a single row of toy data into
a separately running postgres instance via the environment variable POSTGRES_URL

try something like this:
docker run --link postgres:postgres --env 'POSTGRES_URL=postgresql://postgres:@postgres:5432/postgres?sslmode=disable' bowerswilkins/toypostgres
