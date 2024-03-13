## 服务器
IP地址：10.100.33.239

域名：thingue.uino.com

用户名密码：root ue@uino

## 转口转发关系
https://ue-docs.thingjs.com:443 -> http://123.124.196.193:2077 -> http://10.100.33.239:80

域名解析关系

http://thingue.uino.com 等于 http://10.100.33.239

## nginx

访问地址：
- https://ue-docs.thingjs.com    (文档master分支)
- https://ue-docs.thingjs.com/next  (文档dev分支)
- https://ue-docs.thingjs.com/dufs  (文件服务)

配置文件：/etc/nginx/nginx.conf

## dufs文件服务

访问地址：
- http://thingue.uino.com:5000/dufs/  （可以上传下载）
- https://ue-docs.thingjs.com/dufs    （只可以下载）

点击右上角图标登录后可以进行拖拽上传，删除等操作

用户名密码：admin thingue

配置文件：/root/server/dufs/docker-compose.yaml

文件保存位置：/thingue/data/

## gitlab runner

安装路径：/usr/bin/gitlab-runner

注册名称：ue-runner

查看页面：https://git.uino.com/groups/thingjs_base/thingjs_api/ue/-/runners
