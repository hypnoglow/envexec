FROM alpine:3.18

# NOTE: Using "latest" tags is not recommended. Prefer tagged versions.
COPY --from=hypnoglow/envexec:latest-scratch /envexec /usr/local/bin/envexec

ENTRYPOINT ["envexec", "--"]
CMD ["echo", "Hello from envexec!"]
