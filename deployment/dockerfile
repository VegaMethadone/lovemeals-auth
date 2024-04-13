FROM postgres:latest


ENV POSTGRES_DB=testDB
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=0000


RUN echo "host all all 0.0.0.0/0 trust" >>  /etc/postgresql/postgresql.conf
RUN echo "local all all trust" >> /etc/postgresql/postgresql.conf

COPY ./docker-entrypoint-initdb.d/ /docker-entrypoint-initdb.d/ 


EXPOSE 5432


# docker build -t mypostgresql:latest . 
# docker run -d --name testPostgres -p 5432:5432 -e POSTGRES_PASSWORD=0000 -e POSTGRES_USER=postgres mypostgresql:latest