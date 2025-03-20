# Use official Nginx image
FROM nginx:alpine

# Copy the built Vue files into Nginx's HTML directory
COPY my-project/dist /usr/share/nginx/html

# Copy a custom Nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]