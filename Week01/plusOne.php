<?php

$digits = [9,9,9];
var_dump(plusOne($digits));

function plusOne($digits) {


    for($l = count($digits) -1 ; $l>=0;$l--){
        if($digits[$l] != 9){
            $digits[$l]++ ;
            return $digits;
        }else{
            $digits[$l] = 0;
        }
    }
    array_unshift($digits,1);
    return $digits;

}