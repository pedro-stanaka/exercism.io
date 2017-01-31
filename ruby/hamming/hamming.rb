module BookKeeping
  VERSION = 3
end

class Hamming
  def self.compute(original, compare)
    raise ArgumentError unless original.length == compare.length

    original.each_char.with_index.count do |_, i|
      original[i] != compare[i]
    end
  end
end
