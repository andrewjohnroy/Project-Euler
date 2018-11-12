<?php
$largest = 0;
for ($x = 0; $x < 1000; $x++){
	for ($y = 0; $y < 1000; $y++){
		$str = strval($x * $y);
		if ( $str == strrev($str) and $x * $y > $largest ){
			$largest = $x * $y;
		}
	}
}
echo $largest;
