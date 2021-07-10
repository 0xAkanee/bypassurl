<?php
echo "######################################\n";
echo "#         Bypass Shortlink           #\n";
echo "# Coded By Rndzx   fb.me/negevian.id #\n";
echo "######################################\n";
echo "Masukkan URL: ";
$input = fopen("php://stdin","r");
$link = trim(fgets($input));
$ouo = base64_encode($link);
function curl($url){
  $ch = curl_init();
  curl_setopt($ch, CURLOPT_URL, $url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
  $output = curl_exec($ch);
  curl_close($ch);
  return $output;
}

$send = curl("https://lp.nrmn.top/api/bypass?url=.$ouo");
$data = json_decode($send, TRUE);
$url = $data["url"];
echo "Hasil Bypass: ".$url."\n";
?>
