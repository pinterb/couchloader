FROM scratch

COPY couchloader /

ENTRYPOINT ["/couchloader"]
CMD ["--help"]
