<?php

/**
 * @param Integer[] $nums
 * @return Integer
 */
function removeDuplicates(&$nums)
{
    $nums = array_keys(array_flip($nums));
    return count($nums);
}