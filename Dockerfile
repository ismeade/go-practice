FROM ismeade/alpine:3.6
MAINTAINER liyang "ismeade99@sina.com"
ADD ./time /time
EXPOSE 8080
CMD ["./time"]
