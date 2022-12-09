# for running on local, main.go looks for these env vars
# for running in docker, docker-compose loads these vars 
# TODO: Future - place these in secure config mgmt
export DBHOST="localhost"
export DBPORT="5432"
export DBUSER="james"
export DBNAME="workers"