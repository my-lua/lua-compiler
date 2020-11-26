
print("你好，世界，Lua")

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

print(Add(1, 3))
