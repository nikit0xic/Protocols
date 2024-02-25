package org.example.TCP;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;
import java.net.UnknownHostException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class TCPClient {

    public static void main(String[] args) {
        String ip = "109.167.241.255";
        int port = 6340;

        try (Socket socket = new Socket(ip, port);
             BufferedReader reader = new BufferedReader(new InputStreamReader(socket.getInputStream()))) {

            StringBuilder result = new StringBuilder();
            char[] buffer = new char[65536];
            int n;

            while ((n = reader.read(buffer)) != -1) {
                if (n != 4) {
                    result.append(buffer, 0, n);
                    break;
                }
                Thread.sleep(1000);
            }

            Pattern pattern = Pattern.compile("Student17\\s\\S+\\s\\S+");
            Matcher matcher = pattern.matcher(result.toString());
            int count = 0;

            while (matcher.find() && count < 10) {
                System.out.println(matcher.group());
                count++;
            }

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}
