# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity: $$O(nlogn)$$, this is the sorting log time of golang sort library (1 ms).
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: $$O(n)$$, The new array with the sorted keys in a case where each element on N is different (7.86 Mb)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```golang []
import "sort"

func topKFrequent(nums []int, k int) []int {
    var (
        counterHash = make(map[int]int)
    )

    // Count the occurences
    for _, num := range nums {
        counterHash[num] +=  1
    }

    var (
        keys = make([]int, 0, len(counterHash))
    )

    // Sort the map by creating a list of the keys
    for val, _ := range counterHash{
        keys = append(keys, val)
    }

    // then sorting the list of keys by their value on the hashmap
    sort.Slice(keys, func(i,j int) bool {
        return counterHash[keys[i]] > counterHash[keys[j]]
    })

    // Return the list but only the top K
    return keys[0:k]


}
```