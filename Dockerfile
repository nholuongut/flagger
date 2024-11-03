FROM alpine:3.9
LABEL maintainer="Nho Luong <luongutnho@hotmail.com>"
RUN addgroup -S flagger \
    && adduser -S -g flagger flagger \
    && apk --no-cache add ca-certificates

WORKDIR /home/flagger

COPY /bin/flagger .

RUN chown -R flagger:flagger ./

USER flagger

ENTRYPOINT ["./flagger"]

