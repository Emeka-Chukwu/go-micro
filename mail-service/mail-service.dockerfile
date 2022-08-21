

FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY mailerApp /app

CMD [ "/app/mailerApp" ]