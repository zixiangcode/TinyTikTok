# TinyTikTok

666
## 功能需求

具体要求参考飞书[说明文档](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc)

[API 接口文档](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c)

## 编码规范

本文档根据阿里巴巴团队出品的 [Java 开发手册](https://github.com/alibaba/p3c)制定如下编码规范，如有不完善之处，请阅读原文。
同时，[Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)也值得一读。这里仅作为**最基本的规范要求**，如有不完善之处，后续再补充。
### （一）命名风格

1. 代码中的命名均不能以下划线或美元符号开始，也不能以下划线或美元符号结束。  
    反例：`_name / __name / $name / name_ / name$ / name__`
2. 代码中的命名严禁使用拼音与英文混合的方式，更不允许直接使用中文的方式。  
   说明：正确的英文拼写和语法可以让阅读者易于理解，避免歧义。注意，即使**纯拼音命名方式也要避免采用**。  
   正例：`alibaba / taobao / youku / hangzhou` 等国际通用的名称，可视同英文。  
   反例：`DaZhePromotion [打折] / getPingfenByName() [评分] / int 某变量 = 3`
3. 公用的变量、类型、接口、结构、函数以及结构体的成员变量等命名使用 UpperCamelCase 风格，但以下情形例外：DO 等。  
   正例：`GolangStruct / UserDO / XmlService / TcpUdpDeal / TaPromotion`  
   反例：`Golangstruct / UserDo / XMLService / TCPUDPDeal / TAPromotion`
4. 私有的变量、类型、接口、结构、函数以及参数名、局部变量都统一使用 lowerCamelCase 风格，必须遵从驼峰形式。  
   正例： `localValue / getHttpMessage() / inputUserId`
5. 常量命名全部大写，单词间用下划线隔开，力求语义表达完整清楚，**不要嫌名字长**。  
   正例：`MAX_STOCK_COUNT`  
   反例：`MAX_COUNT`
6. 返回结果主要为布尔类型的函数，函数名可以 `is`、`has` 等开头。
7. 杜绝完全不规范的缩写，**避免望文不知义**。  
   反例：`AbstractClass` 缩写命名成 `AbsClass`；`condition` 缩写命名成 `condi`，此类随意缩写严重降低了代码的可阅读性。
8. 为了达到代码自解释的目标，任何自定义编程元素在命名时，使用**尽量完整的单词组合**来表达其意。  
   反例：`var a int` 的随意命名方式。

### （二）常量定义

1. 不允许任何魔法值（即未经预先定义的常量）直接出现在代码中。  
    反例1：
    ```go
    key := "Id#taobao_" + tradeId
    cache.put(key, value) // 缓存 get 时，由于在代码复制时，漏掉下划线，导致缓存击穿而出现问题
    ```
   反例2：
    ```go
   // 设置审查状态，status为2代表审核通过，为3代表退回修改
    switch adp.getStatus() {
        case "2" :
            adp.setStatus("审查通过");
            break;
        case "3" :
            adp.setStatus("退回修改");
            break;
    ```

### 代码格式

**请用 `gofmt` 自动格式化自己的编码！**

## 协助开发要求

1. 每次开发前请进行 `git pull` 获取最新版本的代码，避免造成冲突问题。
2. 开发时应在 `main` 分支以外的分支进行开发，开发之后通过 `pull requests` 与 `main` 分支进行合并。**严禁不经 `review` 强制合并分支！**
3. 推荐创建的开发分支以 `dev_xxx` 命名，建议多建分支，多用分支，每增加一个功能请及时 `commit` ，方便代码溯源。
4. 如有**涉及到对文件结构的更改**请及时沟通说明，同时请在 `pull requests` 时说明，包括但不限于：
   1. 移动文件夹
   2. 移动文件到另一个文件夹
   3. 删除文件夹
   4. 创建新的文件夹
5. 推荐使用 `GoLand` 进行开发，也可以使用 `VS code` 进行开发，设计到的 `git` 相关行为，请尽可能使用图形化界面操作，确保**明确自己在干什么**。
6. 如有不明白或不清楚的操作，请及时交流沟通。在询问他人之前，请通读[提问的智慧](https://github.com/ryanhanwu/How-To-Ask-Questions-The-Smart-Way/blob/main/README-zh_CN.md)和[别像弱智一样提问](https://github.com/tangx/Stop-Ask-Questions-The-Stupid-Ways/blob/master/README.md)。

总而言之，保证一个原则：**让别人明白你是谁，在什么时候，干了什么事**。