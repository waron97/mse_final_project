package org.mse;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

import java.io.IOException;

public class Frontier {
    private String location;
    OkHttpClient client;



    public Frontier(String location) {
        this.location = location;
        this.client = new OkHttpClient();
    }

    public String pop() {
        String url = this.location + "/frontier";
        Request request = new Request.Builder().url(url).build();
        try (Response response = client.newCall(request).execute()) {
            String body = response.body().string();
            return body;
        } catch (IOException e) {
            return null;
        }
    }

    public Boolean isAlive() {
        Request request = new Request.Builder().url(this.location).build();
        try (Response response = client.newCall(request).execute()) {
            return true;
        } catch (IOException e) {
            return false;
        }
    }
}
