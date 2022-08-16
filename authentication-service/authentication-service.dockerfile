

FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY authApp /app

CMD [ "/app/authApp" ]