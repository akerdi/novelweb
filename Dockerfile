FROM centos:7

WORKDIR /app

ENV NODE_ENV production

# 首先安装依赖包
COPY staticBuilds/novelweb_linux /app/
COPY public /app/public
COPY rule.json /app

EXPOSE 8900

CMD ["./novelweb_linux"]
