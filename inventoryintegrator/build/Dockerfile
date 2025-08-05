FROM golang:1.24-alpine AS build
WORKDIR /srv
COPY ./golib ./golib
COPY ./inventoryintegrator ./inventoryintegrator
WORKDIR /srv/inventoryintegrator
RUN go mod download && go mod verify && go mod tidy
WORKDIR /srv/inventoryintegrator/cmd
RUN go build -o bootstrap

FROM golang:1.24-alpine
COPY --from=build /srv/inventoryintegrator/cmd/bootstrap /srv
WORKDIR /srv
CMD [ "./bootstrap" ]

# Checkout https://vsupalov.com/build-docker-image-clone-private-repo-ssh-key/
# ENTRYPOINT ["tail", "-f", "/dev/null"] # for just keep the container up and running to be able to exec into it
