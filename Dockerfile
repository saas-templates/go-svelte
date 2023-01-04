# Build FrontEnd
FROM node:16 as fe-builder
RUN apt update -y && apt install -y golang && npm install -g vite
RUN mkdir /build
ADD ./ui /build/ui
WORKDIR /build/ui
RUN yarn && yarn build

# Build BackEnd
FROM golang:1.19-alpine3.16 as builder
RUN apk add --no-cache make
RUN mkdir /build
ADD ./ /build/
RUN rm -rf /build/ui
COPY --from=fe-builder /build/ui /build/ui
WORKDIR /build
RUN make be

# Package Staage
FROM alpine:3.11.3
COPY --from=builder /build/dist/go-svelte /app/go-svelte
COPY --from=builder /build/config.yml /app/config.yml
CMD [ "/app/go-svelte", "serve", "-c", "/app/config.yml" ]

