### 1、读取yaml/json

### 2、封装执行方法

#### 1、定位方式

### 3、封装断言方法

#### 1、断言方式

#### 2、断言结果失败继续/停止

### 4、封装执行

### 5、日志

### 6、保存报告到json

### 7、集成到平台

统一执行->读取yaml/json->定位/断言->执行->日志->保存报告到json

## 待处理问题

- [x] 怎么全局初始化配置文件，不用每次读取配置文件

## 需要完成的任务 todo

- [x] locator定位单测
- [x] assert断言单测
- [x] 多个期望值 简单
- [x] 断言失败是否继续 复杂
- [x] 断言超时时间 简单
- [ ] 基于工具官方的断言方式 困难
- [ ] 组合定位器，参考python版的定位器 困难
- [ ] 断言结果保存到文件中，包含期望值、实际值、断言方式 复杂
- [ ] 使用Chrome浏览器 简单
- [ ] 网络监听 困难
- [ ] 项目集成到平台 困难
- [ ] 实例项目（react） 一般
