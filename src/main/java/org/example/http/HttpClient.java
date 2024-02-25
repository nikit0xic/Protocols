package org.example.http;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

public class HttpClient {
    public static void main(String[] args) {

        try{
            String url = "http://109.167.241.225:8001/http_example/give_me_five";

            String wday = "2";
            String student = "17";

            URL fullUrl = new URL(url + "?wday=" + wday + "&student=" + student);

            HttpURLConnection connection = (HttpURLConnection) fullUrl.openConnection();

            connection.setRequestMethod("GET");

            connection.addRequestProperty("REQUEST_AGENT", "ITMO student");
            connection.addRequestProperty("COURSE", "Protocols");

            BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
            String inputLine;
            StringBuffer response = new StringBuffer();
            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();

            System.out.println(response.toString());

            System.out.println("Response Code: " + connection.getResponseCode());
            System.out.println("Response Message: " + connection.getResponseMessage());
            System.out.println("Response Body: " + response.toString());

            connection.disconnect();
        } catch (Exception e){
            e.printStackTrace();
        }
    }
}