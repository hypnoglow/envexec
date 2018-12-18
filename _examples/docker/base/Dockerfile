# NOTE: Using "latest" tags is not recommended. Prefer tagged versions.
FROM hypnoglow/envexec:latest-alpine

COPY entrypoint.sh /usr/local/bin/

ENTRYPOINT ["envexec", "--"]
CMD ["entrypoint.sh"]
