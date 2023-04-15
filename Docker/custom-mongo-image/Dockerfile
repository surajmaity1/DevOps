# STEP 1 :- SET THE baseImage
FROM alpine

# STEP 2 :- Install softwares that are required
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories

RUN apk update
RUN apk add mongodb=3.4.4-r0
RUN apk add binutils

# STEP 2.1 (Intermediate) :- configure softwares that needed
VOLUME [ "/data/db" ]
WORKDIR /data
EXPOSE 27017
RUN apk add isl

# STEP 3 :- Set Command
CMD [ "mongod" ]