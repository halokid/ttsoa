<?php
/**
 * User: r00txx
 * Date: 2016/10/19
 * Time: 16:49
 */

/**
 * a sample code use user service base on TTsoa
 */
//require_once 'services_list.php';
require_once 'inc/ttConsumer.php';

$jStr = file_get_contents('./consumer.json');
$serviceList = json_decode($jStr, true);
//print_r($serviceList);exit();
$userSvs = $serviceList['user_php'];      //load user services config, PHP sample
//$userSvs = $serviceList['user_java'];      //load user services config, java sample
//print_r($userSvs);

$cs = new TTsoa\Consumer($userSvs['host_items']);

//php provider sample
//$res = $cs->setEngine($userSvs['engine'])->balanceServ('random')->getService('/user/ulist');
$res = $cs->setEngine($userSvs['engine'])->runServ($userSvs['name'], 'ulist');

//java provider sample
//$res = $cs->setEngine($userSvs['engine'])->balanceServ('random')->getService('/TTsoa/UserImpl');

echo $res;





