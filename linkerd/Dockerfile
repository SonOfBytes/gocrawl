FROM buoyantio/linkerd:1.3.3
ADD disco /disco
COPY config.yaml /config.yaml
EXPOSE 8080 8088 8090 9990
CMD ["/config.yaml"]