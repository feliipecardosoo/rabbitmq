FROM rabbitmq:4.1.0-management

ENV RABBITMQ_DEFAULT_USER=admin
ENV RABBITMQ_DEFAULT_PASS=admin

EXPOSE 5672 15672

HEALTHCHECK --interval=5s --timeout=5s --retries=10 \
  CMD rabbitmq-diagnostics ping || exit 1
