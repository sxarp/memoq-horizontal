# version: 2.0
ARG GCLOUD_SDK_VERSION=228.0.0-alpine

# Use google cloud sdk
FROM google/cloud-sdk:$GCLOUD_SDK_VERSION

# Install Java 8 for Datastore emulator
RUN apk add --update --no-cache openjdk8-jre &&\
gcloud components install cloud-datastore-emulator beta --quiet

COPY start ./

EXPOSE 8081

CMD ["./start"]
