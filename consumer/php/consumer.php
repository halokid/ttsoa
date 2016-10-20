<?php
/**
 * User: r00txx
 * Date: 2016/10/19
 * Time: 16:49
 */


$client = stream_socket_client("tcp://127.0.0.1:7777", $errno, $errorMessage);

if ($client === false) {
  throw new UnexpectedValueException("Failed to connect: $errorMessage");
}

fwrite($client, "hey\n");
echo stream_get_contents($client);
fclose($client);




