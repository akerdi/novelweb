# Novelweb

提供全网所有小说阅读器

## Develop

    go run main.go
    
    cd frontend & yarn dev
    
## Build

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o novelweb_linux
    
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o novelweb_mac
    
    CGO_ENABLED=0 GOOS=windows GOARCH=64 go build -o novelweb_win
    
    
## Plan List

[ ] 首页推荐

建议放到redis，然后项目加个定时2天查看推荐列表

[ ] 小说搜索模板覆盖

[ ] 广告位

[ ] 页面优化

[ ] Content 数据建议放到redis