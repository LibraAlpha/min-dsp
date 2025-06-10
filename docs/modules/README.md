# DSP系统模块设计文档

本目录包含DSP系统各个模块的详细设计文档，每个文档都详细说明了模块的功能点实现路径。

## 模块文档列表

1. [流量接入层设计](./traffic/README.md)
   - [广告请求接收模块](./traffic/request_receiver.md)
   - [流量过滤模块](./traffic/filter.md)
   - [请求预处理模块](./traffic/preprocessor.md)

2. [决策引擎层设计](./decision/README.md)
   - [实时竞价模块](./decision/bidding.md)
   - [预算控制模块](./decision/budget.md)
   - [出价策略模块](./decision/strategy.md)
   - [定向规则引擎](./decision/targeting.md)

3. [数据处理层设计](./data/README.md)
   - [实时数据处理模块](./data/realtime.md)
   - [离线数据处理模块](./data/offline.md)
   - [数据存储模块](./data/storage.md)

4. [运营管理层设计](./operation/README.md)
   - [广告主管理模块](./operation/advertiser.md)
   - [广告位管理模块](./operation/adslot.md)
   - [投放管理模块](./operation/delivery.md)
   - [报表统计模块](./operation/report.md)

## 文档说明

每个模块的设计文档包含以下内容：
1. 模块概述
2. 功能列表
3. 详细设计
4. 接口定义
5. 数据结构
6. 实现路径
7. 性能考虑
8. 安全考虑
9. 测试方案
10. 部署方案

## 使用说明

1. 开发人员可以根据需要查看对应模块的设计文档
2. 文档会持续更新，请确保使用最新版本
3. 如有疑问或建议，请及时反馈 