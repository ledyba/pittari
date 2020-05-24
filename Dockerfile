FROM alpine:3.11

RUN apk add --no-cache wget
RUN ["wget", "https://github.com/ledyba/Pittari/releases/download/v20200524/Pittari_x86-64", "-O", "/Pittari"]
RUN ["chmod", "a+x", "/Pittari"]

EXPOSE 8080
CMD ["/Pittari", "-listen", ":8080"]
