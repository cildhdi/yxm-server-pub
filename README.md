# Server

[![Build Status](https://zyy45666.visualstudio.com/yxm/_apis/build/status/zyy45666.server?branchName=master)](https://zyy45666.visualstudio.com/yxm/_build/latest?definitionId=2&branchName=master)

所有推送到 master 分支的代码均会通过 AzurePipelines 编译，成功后会自动推送到服务器部署，日志请加入 [Team](https://zyy45666.visualstudio.com/yxm/_settings/teams) 后点击上方 Status badge 查看。


# pub

## 数据库，密钥等
修改 `config/config.go` 并修改相应服务器配置

## ci
修改 `azure-pipelines.yml` 中 ssh connection 为你在 AzurePipelines 中的设置，编译路径也要根据服务器的情况修改，不需要在服务器上安装 go 环境