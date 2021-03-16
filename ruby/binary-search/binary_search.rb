=begin
Write your code for the 'Binary Search' exercise in this file. Make the tests in
`binary_search_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/binary-search` directory.
=end

class BinarySearch
    def initialize(data)
        @data = data.sort
    end

    def search_for(item)
        return nil if @data.length == 0
        return 0 if @data.length == 1 and @data[0] == item
        
        do_search(item, @data)
    end

    def do_search(item, data, offset = 0)
        return nil if data.length == 1 and data[0] != item
        
        mid_section = (data.length/2).ceil
        return mid_section + offset if data[mid_section] == item

        if item > data[mid_section]
            do_search(item, data[mid_section..-1], mid_section + offset)
        else
            do_search(item, data[0..mid_section-1], offset)
        end
    end
end