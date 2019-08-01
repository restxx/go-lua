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
