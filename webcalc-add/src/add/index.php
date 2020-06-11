<?php
// header("Access-Control-Allow-Origin: *");
header("Content-type: application/json");
require('functions.inc.php');

$output = array(
	"error" => false,
	"string" => "",
	"answer" => 0
);

$x = $_REQUEST['x'];
$y = $_REQUEST['y'];

$x_is_int = false;
$y_is_int = false;

if ($x[0] == '-')
    $x_is_int = ctype_digit(substr($x, 1));
else
	$x_is_int = ctype_digit($x);


if ($y[0] == '-')
	$y_is_int = ctype_digit(substr($y, 1));
else
	$y_is_int = ctype_digit($y);


if ($x_is_int and $y_is_int)
	$answer=add($x,$y);
else
	$output['error']=true;

$output['string']=$x."+".$y."=".$answer;
$output['answer']=$answer;

echo json_encode($output);
exit();
