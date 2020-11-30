
print("helloworld,Lua")

-- this's my func
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

--[==[
  haha
  it's long comment
]==]
--[[nihao]]
print(Add(1, 3))

--[==[
  ss
]==]print([==[
  笑年郎
]==])
