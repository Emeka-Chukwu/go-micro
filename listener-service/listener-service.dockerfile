FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY listenerApp /app

CMD [ "/app/listenerApp" ]