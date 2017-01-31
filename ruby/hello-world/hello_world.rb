
class HelloWorld 
	def self.hello(pname = nil)
		if pname.nil?
			pname = "World"	
		end

		"Hello, #{pname}!"
	end
end
