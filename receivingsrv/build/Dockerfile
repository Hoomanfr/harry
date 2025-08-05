FROM golang:1.24-alpine AS build
WORKDIR /srv
COPY ./golib ./golib
COPY ./receivingsrv ./receivingsrv
WORKDIR /srv/receivingsrv
RUN go mod download && go mod verify && go mod tidy
WORKDIR /srv/receivingsrv/cmd
RUN go build -o bootstrap

FROM golang:1.24-alpine
COPY --from=build /srv/receivingsrv/cmd/bootstrap /srv
WORKDIR /srv
EXPOSE 8080
CMD [ "./bootstrap" ]

# Checkout https://vsupalov.com/build-docker-image-clone-private-repo-ssh-key/
# ENTRYPOINT ["tail", "-f", "/dev/null"] # for just keep the container up and running to be able to exec into it
