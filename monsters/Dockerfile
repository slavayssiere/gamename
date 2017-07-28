FROM golang:1.8.3-onbuild
ONBUILD RUN latestv=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r .tag_name) \
    && curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/$latestv/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64 \
    && chmod +x /usr/local/bin/swagger \
    && cd /go/src/app \
    && swagger generate spec -o ./public/index.json
CMD ["app"]
