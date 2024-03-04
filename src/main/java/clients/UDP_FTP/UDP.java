package clients.UDP_FTP;

import java.net.*;

public class UDP {
    public static void main(String args[]) throws Exception {
        DatagramSocket clientSocket = new DatagramSocket();
        InetAddress IPAddress = InetAddress.getByName("109.167.241.225");
        int port = 61557;

        for (int i = 1; i <= 100; i++) {
            String data = "<Терехов>: " + generateData(i);
            byte[] sendData = data.getBytes();
            DatagramPacket sendPacket = new DatagramPacket(sendData, sendData.length, IPAddress, port);
            clientSocket.send(sendPacket);
            System.out.println("Sent: "+ sendData + " " + sendData.length + " " +  IPAddress + " " + port);
            Thread.sleep(1000);
        }

        clientSocket.close();
    }

    public static int generateData(int iteration) {
        return (int) Math.pow(2, iteration);
    }
}
