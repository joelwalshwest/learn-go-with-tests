FROM joelwalshwest/my-development-environment

WORKDIR /code

COPY . .

# Copy local nvim configurations
COPY .nvim/local.lua /root/.config/nvim/lua/joelwalshwest/
RUN echo "require('joelwalshwest.local')" >> /root/.config/nvim/lua/joelwalshwest/init.lua
