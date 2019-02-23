however, it should be noted that in the case of insufficient capacity, appendthe operation will result in reallocation of memory, which may result in huge memory allocation and copy data cost. Even if the capacity is sufficient, you still need appendto update the slice itself with the return value of the function, because the length of the new slice has changed.


Therefore, the performance of adding elements from the beginning of a slice is generally much worse than the performance of adding elements from the tail.

You can use copyand appendcombine to avoid creating intermediate temporary slices, as well as completing the addition of elements