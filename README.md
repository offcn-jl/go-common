# Golang 通用组件
[![Status](https://img.shields.io/badge/Status-GA-brightgreen)](#当前版本) [![Version](https://img.shields.io/badge/Release-0.1.0-brightgreen)](#当前版本) [![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/offcn-jl/go-common)](https://goreportcard.com/report/github.com/offcn-jl/go-common) [![Master Build](https://github.com/offcn-jl/go-common/workflows/Master%20Build/badge.svg)](https://github.com/offcn-jl/go-common/actions?query=workflow%3A%22Master+Build%22) [![codecov](https://codecov.io/gh/offcn-jl/go-common/branch/master/graph/badge.svg)](https://codecov.io/gh/offcn-jl/go-common) [![Build](https://github.com/offcn-jl/go-common/workflows/Build/badge.svg)](https://github.com/offcn-jl/go-common/actions?query=workflow%3ABuild) [![codecov](https://codecov.io/gh/offcn-jl/go-common/branch/new-feature/graph/badge.svg)](https://codecov.io/gh/offcn-jl/go-common/branch/new-feature) 

## 目录结构 ( 按文件名排序 ) 及组件列表

- .github  ( Github 配置  
    - workflows ( 工作流配置 ) 
        - build.yml ( new-feature 分支构建配置 )  
        - master-build.yml ( master 分支构建配置 )  
- codes  ( 代码库 )  
    - error.go ( 错误代码 )  
    - README.md ( 说明 )  
- configer  ( 配置库 )  
    - configer.go ( 配置库 )  
- logger  ( 日志库 )  
    - logger.go ( 日志库 )  
- .gitignore ( GIT 配置文件，用于配置需要忽略提交的内容 )  
- go.mod ( go mod 配置文件 )  
- go.sum ( go mod 相关文件 )  
- LICENSE ( 版权声明 )  
- README.md ( 使用说明 )  
- version.go ( 通用组件版本 )

## 当前版本
当前版本 : 正式版 0.1.1

暂定版本发布流程 : Alpha -> Beta -> RC -> GA

> Alpha : 内部测试版, 一般不向外部发布  
> Beta : 也是测试版, 这个阶段的版本会一直加入新的功能  
> RC : 发行候选版本, 基本不再加入新的功能, 主要进行缺陷修复  
> GA : 正式发布的版本, 采用 Release X.Y.Z 作为发布版本号  
> 参考 : [Alpha、Beta、RC、GA版本的区别](http://www.blogjava.net/RomulusW/archive/2008/05/04/197985.html) 、 [软件版本GA、RC、beta等含义](https://blog.csdn.net/gnail_oug/article/details/79998154)
