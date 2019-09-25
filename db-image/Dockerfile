# Derived from official mysql image (our base image)
FROM mysql:8
# Add a database
ENV MYSQL_DATABASE grpc
# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
ADD ./sql-scripts/ /docker-entrypoint-initdb.d/

RUN chmod 777 -R /docker-entrypoint-initdb.d/