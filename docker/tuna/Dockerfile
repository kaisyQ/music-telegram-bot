FROM alpine/curl AS tuna-alpine
WORKDIR /var/www/app
RUN curl -sSLf https://get.tuna.am | sh
COPY .env.local .env
COPY ./sh/init-tunnel init-tunnel
ENTRYPOINT ./init-tunnel
