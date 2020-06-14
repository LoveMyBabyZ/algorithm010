<?php

function rotate(&$nums, $k) {
    if (count($nums) != 1 && $k != 0 && count($nums) != $k){
        $k = $k%count($nums);
        $nums = array_merge(array_slice($nums, 0 - $k, count($nums) - 1), array_slice($nums, 0, count($nums) - $k));
    }
}


function rotate2(&$nums, $k) {
    $count = count($nums);
    $k = $k % $count;
    $this->reverse($nums, 0, $count-1);
    $this->reverse($nums, $k, $count-1);
    $this->reverse($nums, 0, $k-1);
}

function reverse(&$nums, $s, $e){
    while ($s<$e) {
        $tmp = $nums[$s];
        $nums[$s] = $nums[$e];
        $nums[$e] = $tmp;
        $s++;
        $e--;
    }
}

