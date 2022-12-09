# 'source env.sh' for running on local, main.go looks for these
# for running in docker, docker-compose.yml has its own vars
# TODO: Future - place these in secure config mgmt
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_NAME="workers"
export DB_PASSWORD="postgres"