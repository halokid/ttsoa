<?php
/**
 * User: r00xx
 * Date: 2016/10/21
 * Time: 15:55
 */

require_once __DIR__.'/inc/user.php';

echo $argc;
print_r($argv);

$u = new TTsoa\User();
$u->ulist();
$u->$argv[1]();

//echo 'xx';
