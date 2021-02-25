FROM gitpod/workspace-full

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/

RUN apt-get update -y && \
    apt-get install -y gcc libc6-dev libx11-dev xorg-dev libxtst-dev libpng++-dev xcb \
            libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev \
            libxkbcommon-dev \
            xsel xclip
