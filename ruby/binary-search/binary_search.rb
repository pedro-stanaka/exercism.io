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
        return if data.length == 0
        return 0 if data.length == 1 and data[0] == item
        
        do_search(item, data)
    end

    private
    
    attr_reader :data
    
    def do_search(item, search_spc, offset = 0)
        return if search_spc.length == 1 and search_spc[0] != item
        
        mid_section = (search_spc.length/2).ceil
        return mid_section + offset if search_spc[mid_section] == item

        if item > data[mid_section]
            do_search(item, search_spc[mid_section..-1], mid_section + offset)
        else
            do_search(item, search_spc[0..mid_section-1], offset)
        end
    end
end