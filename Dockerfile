FROM debian
COPY bin/quarky-test /bin/quarky-test
ENTRYPOINT ["/bin/quarky-test"]
