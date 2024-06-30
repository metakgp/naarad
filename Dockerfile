FROM binwiederhier/ntfy

RUN apk --no-cache add jq

# Copy metaploy configuration
COPY metaploy/naarad.metaploy.conf /
COPY metaploy/postinstall.sh /

# Set the postinstall script as executable
RUN chmod +x /postinstall.sh

EXPOSE 8000

ENTRYPOINT ["/postinstall.sh", "ntfy", "serve"]