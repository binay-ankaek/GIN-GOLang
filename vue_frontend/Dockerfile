# # Stage 1: Build the Vue.js application
# FROM node:20.11.1 AS builder

# # Set the working directory
# WORKDIR /app

# # Copy package.json and yarn.lock (or package-lock.json) and install dependencies
# COPY vue_frontend/usermgmt/package.json vue_frontend/usermgmt/yarn.lock ./
# RUN yarn install

# # Copy the rest of the application code
# COPY vue_frontend/usermgmt ./

# # Build the Vue.js application for production
# RUN yarn build

# # Stage 2: Serve the application using nginx
# FROM nginx:alpine

# # Copy the built files from the build stage

# COPY --from=builder /app/dist /usr/share/nginx/html

# # Expose port 80 for the web server
# EXPOSE 80

# # Start nginx server
# CMD ["nginx", "-g", "daemon off;"]

# Stage 1: Build the Vue app
FROM node:20 AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY vue_frontend/usermgmt/package.json vue_frontend/usermgmt/yarn.lock ./

# Install dependencies
RUN yarn install

# Copy the rest of the frontend code
COPY vue_frontend/usermgmt .

# Build the Vue app
RUN yarn build

# Stage 2: Serve the Vue app
FROM nginx:alpine

# Copy built files from the builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Run Nginx
CMD ["nginx", "-g", "daemon off;"]

