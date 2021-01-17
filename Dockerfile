FROM golang:1.15.6-alpine

LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"

# change dir container to this directory
WORKDIR /app

# copy go.mod dan go.sum first 
COPY go.* ./

# download third_party from go mod
RUN go mod download 

# copy rest of the files
COPY . .

# build go aplication
RUN go build -o application .

# add permission executable
RUN chmod +x application

CMD ["./application"]