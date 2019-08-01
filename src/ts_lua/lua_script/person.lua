p = person.new("Steeve")
print(p:name()) -- "Steeve"

p:name("Alice", "Alice2")

p:buf({11,22,33,44,55,66,77,88,99,10})

print("-------------------------------------")

for k, v in pairs(p:buf()) do
    print(v)
end
