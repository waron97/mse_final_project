package org.mse;

import com.google.gson.Gson;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class Logger {
    private static Boolean sendLog(String level, String location, String message, Object data) {
        Gson gson = new Gson();
        Map<String, String> body = new HashMap<>();
        body.put("level", level);
        body.put("appId", Constants.logsAppId);
        body.put("location", location);
        body.put("message", message);

        if (data != null) {
            body.put("detail", gson.toJson(data));
        }

        OkHttpClient client = new OkHttpClient();
        String payload = gson.toJson(body);
        Request req = new Request.Builder()
                .url(Constants.logsUrl + "/logs")
                .post(RequestBody.create(payload, Constants.JSON))
                .header("Content-Type", "application/json")
                .header("Authorization", "apiKey " + Constants.logsKey)
                .build();

        try (Response res = client.newCall(req).execute()) {
            return true;
        } catch(Exception e) {
            System.out.println("Failed to send log");
            System.out.println(e);
            return false;
        }
    }

    public static Boolean debug(String location, String message, Object data) {
        return sendLog("debug", location, message, data);
    }

    public static Boolean info(String location, String message, Object data) {
        return sendLog("info", location, message, data);
    }

    public static Boolean warning(String location, String message, Object data) {
        return sendLog("warning", location, message, data);
    }

    public static Boolean error(String location, String message, Object data) {
        return sendLog("error", location, message, data);
    }

    public static Boolean critical(String location, String message, Object data) {
        return sendLog("critical", location, message, data);
    }

    public static Boolean isAlive() {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder().url(Constants.logsUrl).build();
        try (Response response = client.newCall(request).execute()) {
            return true;
        } catch (IOException e) {
            return false;
        }
    }
}
