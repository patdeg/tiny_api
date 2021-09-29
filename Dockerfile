FROM scratch

# Copy application executable
COPY  application /app/application

# Run the application binary.
ENTRYPOINT ["/app/application"]