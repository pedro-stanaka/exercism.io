
class HelloWorld 
	def self.hello(pname = nil)
		pname = "World" if pname.nil?
		"Hello, #{pname}!"
	end
end
