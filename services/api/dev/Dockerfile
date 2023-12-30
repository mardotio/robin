FROM golang:1.21

# We need nodemon because things like "gow" don't work in docker bceause of inotify
ENV NODE_VERSION=16.13.0
ENV NVM_DIR=/root/.nvm
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN npm install -g nodemon

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download
COPY ../main.go ../src ./

CMD ["nodemon", "--signal", "SIGTERM", "--watch", "src", "--watch", "main.go", "--ext", "go", "--exec", "go run ."]
