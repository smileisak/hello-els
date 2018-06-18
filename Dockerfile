# base image (scratch)
FROM scratch

# Image maintainer
MAINTAINER Ismail KABOUBI <iskab@smile.fr>

# Copy server binary to docker image
COPY server /

# Run server binary when a container start using this images
CMD ["/server"]
