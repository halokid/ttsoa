<?php
/**
 * User: r00xx
 * Date: 2016/10/20
 * Time: 15:49
 */

$serviceList = array(
  'user' => array('name' => '用户服务',
                   'hosts' => array('127.0.0.1:7777', '127.0.0.1:7777'),

                  'host_items' => array( '127.0.0.1:7777' => array('ulist', 'detail'),
                                         '192.168.0.203:7777' => array('ulist', 'detail'),
                                   ),

                  'all_items' => array('ulist', 'detail', 'reg', 'login', 'logout'),
                  ),
);

echo json_encode($serviceList);