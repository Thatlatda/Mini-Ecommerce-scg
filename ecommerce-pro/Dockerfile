# Specify a base image
FROM node:alpine
# set working directory
WORKDIR /ecommerce-pro

# add `/app/node_modules/.bin` to $PATH
ENV PATH="./node_modules/.bin:$PATH"

COPY package.json .
COPY package-lock.json .

# add app
COPY . .

RUN npm install

# start app
CMD ["npm", "start"]