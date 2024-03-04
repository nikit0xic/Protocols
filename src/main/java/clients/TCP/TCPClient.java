package clients.TCP;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;
import java.net.UnknownHostException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class TCPClient {
    public static void main(String[] args) {
        String ip = "109.167.241.225";
        int port = 601;

        try {
            Socket socket = new Socket(ip, port);
            BufferedReader reader = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            StringBuilder result = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                result.append(line);
            }

            socket.close();

            Pattern pattern = Pattern.compile("Student17\\s\\S+\\s\\S+");
            Matcher matcher = pattern.matcher(result.toString());

            int count = 0;
            while (matcher.find() && count < 10) {
                System.out.println(matcher.group());
                count++;
            }
        } catch (UnknownHostException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
