<?php
/**
 * User: r00xx
 * Date: 2016/10/20
 * Time: 15:39
 */

/**
 * request services with balancing, ACL, etc..
 */
namespace TTsoa;

use Exception;

class Consumer
{
//  private $_engine = 'php';     //what language we use in service container
  private $_dsns;
  private $_engine;

  function __construct($dsns)
  {
    $this->_dsns = $dsns;   // read from the services list, array
  }


  /**
   * @param $balanType balancing type
   * @param $service services name
   * @return string
   * @throws Exception
   */
  public function runServ($serviceName, $serviceAct, $balanceType = 'random')
  {
    $enableServHost = $this->checkServ($this->_dsns, $serviceAct);     //fiter the enable serv host
//    echo '----------'.$this->_balancegServ.'-----------';
    $balanceHost = $this->balanceServ($enableServHost, $balanceType);
    $client = stream_socket_client("tcp://".$balanceHost, $errno, $errorMessage);

    if ($client === false) {
      throw new Exception("Failed to connect: $errorMessage");
    }

    fwrite($client, $this->_engine.'/'.$serviceName.'/'.$serviceAct."\n");
    $res = stream_get_contents($client);
    fclose($client);
    return $res;
  }

  /**
   * @param $type  type of the balancing request without check services exsit or not
   */
  public function balanceServ($enableServHost, $type) {
    $balancegServ = '';
    switch ($type)
    {
      case 'random':
        $rand = array_rand($enableServHost);
//        echo '---------------'.$rand.'----------------';
        $balancegServ = $enableServHost[$rand];
        break;
      case 'leastConn':     //least-ConnectionScheduling
        $balancegServ = $this->leastConnSche($this->_dsns);
        break;
    }

    return $balancegServ;
  }


  private function checkServ($dsns, $serviceAct) {
    $enableHost = array();
    foreach ($dsns as $k=>$v) {
      if (in_array($serviceAct, $v)) {
        $enableHost[] = $k;
      }
    }
    return $enableHost;
  }

  /**
   * arithmetic of least-Conn
   * @param $dsns
   */
  public function leastConnSche($dsns) {
    return false;
  }


  public function setEngine($engine) {
    $this->_engine = $engine;
    return $this;
  }



}










