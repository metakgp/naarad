FROM binwiederhier/ntfy

COPY server.yml /etc/ntfy/server.yml

EXPOSE 8000

# Copy metaploy configuration
COPY metaploy/naarad.metaploy.conf /
COPY metaploy/postinstall.sh /

# Set the postinstall script as executable
RUN chmod +x /postinstall.sh

ENTRYPOINT ["/postinstall.sh", "ntfy", "serve"]