# Base Image
FROM node:alpine

# Working dir
WORKDIR /usr/my-node-app

# Copy Files
COPY ./package.json ./

# RUN Installer
RUN npm install

COPY ./ ./

# Default Command
CMD [ "npm", "start" ]