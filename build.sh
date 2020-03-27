# 开始删除public 文件
echo `dirname $0`
echo "rm public files"
rm -r public && echo "rm public files success"

# 开始打包Go 执行文件到/staticBuilds
echo "start build linux package"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o staticBuilds/novelweb_linux
echo "build linux package success"

# 开始打包public 文件
echo "start build frontend"
cd frontend && echo `dirname $0` && yarn build && cp -r dist ../public && echo "cp -r dist ../public"

cd ../
echo `dirname $0`

echo "start build docker container..."
# 开始打包docker 文件；然后上传
docker build -t registry.gitlab.com/honglian/novelweb . && echo "build docker latest image complete!"
docker push registry.gitlab.com/honglian/novelweb && echo "push docker latest image complete!"