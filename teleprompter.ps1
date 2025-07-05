$client = New-Object System.Net.Sockets.TcpClient
$client.Connect("localhost", 8080)
$stream = $client.GetStream()
$writer = New-Object System.IO.StreamWriter($stream)
$writer.AutoFlush = $true

while ($true) {
    $input = Read-Host "Enter message"
    if ($input -eq "exit") {
        break
    }
    $writer.WriteLine($input)
}

$writer.Close()
$stream.Close()
$client.Close()
