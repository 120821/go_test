def generate_small_array n
  result = []
  i = 1
  # 第一次循环
  loop do
    result.push(n)
    puts "====== 每个循环1, i: #{i}, result #{result.inspect}"

    i = i + 1
    if i > n
      break
    end
  end
  return result
end

def generate_array n
  result = []
  i = 1
  # 第二次循环
  loop do
    small_array = generate_small_array(i)
    puts "=== small_array #{small_array.inspect}"
    result.push(small_array)
    puts "=== 每个循环2， i: #{i}, small_array #{small_array.inspect} result #{result.inspect}"

    i =  i + 1
    if i > n
      break
    end
  end
  return result
end

puts generate_array(5).inspect
