FROM alpine
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
ADD mikudos-schedule-srv /mikudos-schedule-srv
WORKDIR /
ENTRYPOINT [ "/mikudos-schedule-srv" ]
