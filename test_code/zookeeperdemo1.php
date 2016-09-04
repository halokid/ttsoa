<?php
 
class ZookeeperDemo extends Zookeeper {
 
  public function watcher( $i, $type, $key ) {
    echo "Insider Watcher\n";
 
    // Watcher gets consumed so we need to set a new one
    $this->get( '/test', array($this, 'watcher' ) );
  }
 
}
 
$zoo = new ZookeeperDemo('127.0.0.1:2181');
$zoo->get( '/test', array($zoo, 'watcher' ) );
 
while( true ) {
  echo '.';
	sleep(2);
}

?>

