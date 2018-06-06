FROM ponzu/ponzu:latest

ENV GOPATH /go
ENV GO_SRC $GOPATH/src
ENV TRIPEZIO_ROOT github.com/lsilvapvt/tripezio
ENV PROJECT_ROOT $GO_SRC/$TRIPEZIO_ROOT

RUN mkdir -p $PROJECT_ROOT

# All commands will be run inside of ponzu root
WORKDIR $PROJECT_ROOT

# Copy the ponzu source into ponzu root.
COPY . .

# the following runs the code inside of the $GO_SRC/$PONZU_GITHUB directory
RUN ponzu build

EXPOSE 9888

# Define the scripts we want run once the container boots
# CMD [ "ponzu","run","--port=9888" ]
