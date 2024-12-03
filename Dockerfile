FROM griffingrid/tdlib-golang-1.23.2 as builder

RUN mkdir /app
RUN echo $CGO_CFLAGS

COPY . /app
WORKDIR /app
RUN go mod tidy

RUN go build -o bin/tg_parser /app/cmd

FROM griffingrid/tdlib-golang-1.23.2

WORKDIR /app

COPY --from=builder /app/bin /app
COPY --from=builder /app/.tdlib /app/.tdlib

CMD ["/app/tg_parser"]