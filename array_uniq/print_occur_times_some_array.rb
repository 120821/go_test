def print_occur_times some_array

  temp = {}

  for i in some_array
    #puts i
    if temp[i] != nil
      temp[i] = temp[i] + 1
    else
      temp[i] = 1
    end
  end

  puts temp
end

my_array = [1,2,3,3,3,4,4,5,5,5,5,5,5,8,9,9]
print_occur_times(my_array)
