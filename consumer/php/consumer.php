<?php
/**
 * User: r00txx
 * Date: 2016/10/19
 * Time: 16:49
 */

/**
 * a sample code use user service base on TTsoa
 */
require_once 'services_list.php';
require_once 'inc/ttConsumer.php';

$userSvs = $serviceList['user'];      //load user services config
print_r($userSvs);

$cs = new TTsoa\Consumer($userSvs['hosts']);
$res = $cs->balanceServ('random')->getService('/user/ulist');
echo $res;





