
print("你好，世界，Lua")

-- 这是我编写的样例函数
function Add(a, b)
  local sum = (a + b) * b
  for i = 1, sum, 1 do
    print(i)
  end
  if sum > 10 then
    print("a")
  else
    print("b")
  end
  return sum
end

--[[
  哈哈
  这是一个可以换行的注释
]]
--[[nihao]]
print(Add(1, 3))
