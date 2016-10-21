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
  private $_balancegServ;
  private $_engine;

  function __construct($dsns)
  {
    $this->_dsns = $dsns;   // read from the services list, array
  }

  /**
   * @param $service /user/list
   * @throws Exception
   */
  public function getService($service)
  {
    print_r($this->_dsns);
    echo '----------'.$this->_balancegServ.'-----------';
    $client = stream_socket_client("tcp://".$this->_balancegServ, $errno, $errorMessage);

    if ($client === false) {
      throw new Exception("Failed to connect: $errorMessage");
    }

    fwrite($client, $this->_engine.$service."\n");
    $res = stream_get_contents($client);
    fclose($client);
    return $res;
  }

  /**
   * @param $type  type of the balancing request
   */
  public function balanceServ($type) {
    switch ($type)
    {
      case 'random':
        $rand = array_rand($this->_dsns);
        echo '---------------'.$rand.'----------------';
        $this->_balancegServ = $this->_dsns[$rand];
        break;
      case 'leastConn':     //least-ConnectionScheduling
        $this->_balancegServ = $this->leastConnSche($this->_dsns);
        break;
    }

    return $this;
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










