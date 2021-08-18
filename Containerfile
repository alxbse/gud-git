FROM scratch
COPY gud-git /gud-git

ENTRYPOINT ["/gud-git"]
