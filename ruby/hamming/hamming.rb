module BookKeeping
  VERSION = 3
end 

class Hamming
  def self.compute(original, compare)

    if original.length != compare.length 
      raise ArgumentError, "Can't calculate the distance for different length strings."
    end

    original.each_char.with_index.count do |_, i| 
     original[i] != compare[i]
   end
 end
end
