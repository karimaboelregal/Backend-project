# Use Ruby slim base image for smaller size
FROM ruby:3.2-slim AS base

# Set the working directory
WORKDIR /app

# Install required system dependencies
RUN apt-get update -qq && apt-get install -y \
  build-essential \
  libyaml-dev \
  libssl-dev \
  libreadline-dev \
  libsqlite3-dev \
  sqlite3 \
  nodejs \
  npm \
  default-libmysqlclient-dev \
  git \
  && rm -rf /var/lib/apt/lists/*

# Install psych before bundling
RUN gem install psych -v 5.2.3

# Copy only Gemfile to leverage Docker caching
COPY Gemfile Gemfile.lock ./

# Install application gems
RUN bundle install --jobs 4 --retry 3 && \
    rm -rf ~/.bundle/ "${BUNDLE_PATH}"/ruby/*/cache "${BUNDLE_PATH}"/ruby/*/bundler/gems/*/.git

# Copy application source code
COPY . .

# Ensure tmp directories exist and are writable
RUN mkdir -p /app/tmp/sockets /app/tmp/pids /app/tmp/cache

# 🚨 Make sure all files are owned by `rails` **before switching users**
RUN chown -R 1000:1000 /app

# 🔄 Switch to the `rails` user
USER 1000:1000

# Expose the port
EXPOSE 3000

# Entrypoint for DB preparation
CMD ["bin/rails", "server", "-b", "0.0.0.0"]