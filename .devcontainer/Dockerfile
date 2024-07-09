ARG UBUNTU_VERSION=22.04
FROM mcr.microsoft.com/devcontainers/base:ubuntu-${UBUNTU_VERSION}

# Update apt-get and install various needed utilities
RUN set -ex \
    apt-get update \
    && apt-get install --no-install-recommends -y curl \
        wget \
        xz-utils \
        make \
        gcc \
        git

# Install bridged provider prerequisites
# See README.md

# Install go
ARG GO_VERSION=1.21.12
ENV PATH=$PATH:/usr/local/go/bin
RUN set -ex \
    rm -rf /usr/local/go \
    && export arch=$(uname -m | awk '{ if ($1 == "x86_64") print "amd64"; else if ($1 == "aarch64" || $1 == "arm64") print "arm64"; else print "unknown" }') \
    && curl -L "https://go.dev/dl/go${GO_VERSION}.linux-${arch}.tar.gz" | tar -C /usr/local/ -xzvf - \
    && which go \
    && go version

# Install go linter
RUN mkdir -p $GOPATH/bin && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.46.2

ENV PATH=$PATH:$GOPATH/bin

# Install pulumictl
RUN set -ex \
    && export arch=$(uname -m | awk '{ if ($1 == "x86_64") print "amd64"; else if ($1 == "aarch64" || $1 == "arm64") print "arm64"; else print "unknown" }') \
    && export urlPulumiRelease="https://api.github.com/repos/pulumi/pulumictl/releases/latest" \
    && export urlPulumiVersion=$(curl -s ${urlPulumiRelease} | awk -F '["v,]' '/tag_name/{print $5}') \
    && export urlPulumiBase="https://github.com/pulumi/pulumictl/releases/download" \
    && export urlPulumiBin="pulumictl-v${urlPulumiVersion}-linux-${arch}.tar.gz" \
    && export urlPulumi="${urlPulumiBase}/v${urlPulumiVersion}/${urlPulumiBin}" \
    && curl -L ${urlPulumi} | tar xzvf - --directory /tmp \
    && chmod +x /tmp/pulumictl \
    && mv /tmp/pulumictl /usr/local/bin/ \
    && which pulumictl \
    && pulumictl version \
    && rm -rf /tmp/* \
    && true

# Install nodejs
RUN set -ex \
    && export NODE_MAJOR=18 \
    && curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key \
        | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg \
    && echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" \
        | tee /etc/apt/sources.list.d/nodesource.list \
    && apt-get update \
    && apt-get install --no-install-recommends nodejs \
    && apt-get clean \
    && apt-get autoremove -y \
    && apt-get purge -y --auto-remove \
    && rm -rf \
        /var/lib/{apt,dpkg,cache,log} \
        /usr/share/{doc,man,locale} \
        /var/cache/apt \
        /root/.cache \
        /var/tmp/* \
        /tmp/* \
    && node --version \
    && npm --version \
    && npm install --global yarn \
    && yarn --version \
    && true

# Install python and related items
RUN apt-get install -y --no-install-recommends python3 && \
    apt-get install -y --no-install-recommends python3-setuptools python3-venv python3-pip

# Install .NET
RUN set -ex \
    && apt-get update \
    && apt-get install -y --no-install-recommends dotnet-sdk-6.0 dotnet-runtime-6.0 \
    && apt-get clean \
    && apt-get autoremove -y \
    && apt-get purge -y --auto-remove \
    && rm -rf \
        /var/lib/{apt,dpkg,cache,log} \
        /usr/share/{doc,man,locale} \
        /var/cache/apt \
        /root/.cache \
        /var/tmp/* \
        /tmp/* \
    && true

# Install pulumi
RUN set -ex \
    && export arch=$(uname -m | awk '{ if ($1 == "x86_64") print "x64"; else if ($1 == "aarch64" || $1 == "arm64") print "arm64"; else print "unknown" }') \
    && export urlPulumiRelease="https://api.github.com/repos/pulumi/pulumi/releases/latest" \
    && export urlPulumiVersion=$(curl -s ${urlPulumiRelease} | awk -F '["v,]' '/tag_name/{print $5}') \
    && export urlPulumiBase="https://github.com/pulumi/pulumi/releases/download" \
    && export urlPulumiBin="pulumi-v${urlPulumiVersion}-linux-${arch}.tar.gz" \
    && export urlPulumi="${urlPulumiBase}/v${urlPulumiVersion}/${urlPulumiBin}" \
    && curl -L ${urlPulumi} | tar xzvf - --directory /tmp \
    && chmod +x /tmp/pulumi/* \
    && mv /tmp/pulumi/* /usr/local/bin/ \
    && which pulumi \
    && pulumi version \
    && rm -rf /tmp/* \
    && true
