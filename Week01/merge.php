<?php



function merge(&$nums1, $m, $nums2, $n)
{
    $nums1 = array_merge_recursive($nums1,$nums2);
    rsort($nums1);
}