FROM scratch
WORKDIR /
ADD clash-ui /clash-ui
USER 1001
EXPOSE 8088
ENTRYPOINT ["/clash-ui"]