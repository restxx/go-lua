print("hello, I`m from lua!")

--[[ 函数返回两个值的最大值 --]]
function max(num1, num2)
   print(num1, num2)
   if (num1 > num2) then
      result = num1;
   else
      result = num2;
   end

   return result;
end

print(double(20))
