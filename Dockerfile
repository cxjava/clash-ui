FROM scratch
WORKDIR /
ADD clash-ui /
USER 1000
EXPOSE 8088
ENTRYPOINT ["/clash-ui"]