### 目前对gopher-lua的认识
##### 是一个golang版的lua vm 基于lua5.1
##### 提供了一套仿C-api 给go用于嵌入
##### 能实现互相调用 传递go struct给lua使用
##### 实现了byte流打包后可以相互传递


### 目前对gopher的测试
##### golang直接实现tsFib函数
##### golang调用lua里的Fib函数
##### 分别在gopher，原生lua和golang执行tsFib(35) 进行了执行时间对比
```
function Fib(n)
    if n == 0 then
        return 0
    end
    if n == 1 then
        return 1
    end
    return Fib(n-2)+Fib(n-1)
end


function tsFib(n)
    local start = os.clock()
    ret, n = 0, n-1
    for i=0, n, 1 do
        ret = Fib(i)
    end
    print(os.clock()-start)
    return ret
end

--start = os.clock()
--print(tsFib(35))
--print(os.clock()-start)

--gopher 比原生lua慢近9倍   比go慢160倍
```

### 系统设计
##### golang 连接启动收发go loop()
##### 收到数据处理成明文的単条协议后交给dispatch 
##### dispatch 分发给各*.lua解析
##### 封装goSend()给lua用于解析后打包。
##### 封装一套packet类提供GetInt SetInt之类方便读写

### 问题
##### 这样的测试是否足够合理
##### 上次提到的I/O测试是指什么
##### 第一版这么设计行么
##### 协议可配置具体是指什么
