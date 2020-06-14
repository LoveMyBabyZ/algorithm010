<?php


$nums = [1,0,1];
$i = 0;


for($j = 1;$j< count($nums);$j++){
    if($nums[$i] == 0 && $nums[$j] != 0){
        $nums[$i] = $nums[$j];
        $nums[$j] = 0;
        $i++;
    }elseif ($nums[$i] != 0 ){
        $i++;
    }

}

echo $j;
var_dump($nums);