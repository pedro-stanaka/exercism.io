module BookKeeping
  VERSION = 5
end

class Gigasecond
  NANO_SECOND = 1_000_000_000

  def self.from(past_time)
    past_time + NANO_SECOND
  end
end
