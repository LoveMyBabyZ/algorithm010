<?php
function groupAnagrams($strs)
{
    $return = [];
    foreach ($strs as $str) {
        $s = str_split($str, 1);
        sort($s);
        $s2 = implode('', $s);
        $return[$s2][] = $str;
    }
    return array_values($return);
}