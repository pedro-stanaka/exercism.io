module BookKeeping
  VERSION = 3
end 

class Hamming
  def self.compute(original, compare)

    if original.length != compare.length 
      raise ArgumentError, "Can't calculate the distance for different length strings."
    end

    return 0 if original == compare

    distance = 0

    original.each_char.with_index do |o, i| 
      distance += 1 if original[i] != compare[i]
    end

    distance
  end
end