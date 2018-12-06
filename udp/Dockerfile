FROM goodrainapps/alpine:3.4

ADD server/server /run/server
ADD entrypoint.sh /run/entrypoint.sh

RUN chmod +x /run/entrypoint.sh

EXPOSE 38080

ENTRYPOINT ["/run/entrypoint.sh"]