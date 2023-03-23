class Klass
  def hello(*args)
    "Hello " + args.join(' ')
  end
end
k = Klass.new
a = k.send :hello, "gentle", "readers"
#=> "Hello gentle readers"
puts a
