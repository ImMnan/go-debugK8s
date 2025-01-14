FROM --platform=linux/amd64 debian 

COPY ./app /app
RUN chmod +x /app
# Install necessary tools for debugging
# RUN apt-get update && apt-get install -y file
ENTRYPOINT ["/app"]