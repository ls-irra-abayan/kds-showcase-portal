FROM alpine:3

ARG COMPONENT_NAME
ENV COMPONENT_NAME="$COMPONENT_NAME"

COPY "$COMPONENT_NAME" "$COMPONENT_NAME"
COPY docker-start.sh docker-start.sh

RUN chmod a+x "$COMPONENT_NAME" docker-start.sh

CMD ["./docker-start.sh"]
