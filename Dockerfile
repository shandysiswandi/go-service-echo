FROM golang:1.15.6-alpine

# make directory for this  
RUN mkdir /app  

# copy go.mod dan go.sum first 
COPY go.* /app 

# download third_party from go mod
RUN go mod download 

# copy rest of the files
COPY . /app

# change dir container to this directory
WORKDIR /app

# build go aplication
RUN go build -o server . 

CMD ["/app/server"]