package org.mse;

import okhttp3.MediaType;

public class Constants {
    public static String frontierUrl = "http://frontier-acceptor:3000";
    public static String acceptorUrl = "http://frontier-acceptor:3000/acceptor";
    public static String logsUrl = "http://logs:8080";
    public static String logsKey = System.getenv("LOGS_KEY");
    public static String logsAppId = System.getenv("LOGS_APP_NAME");
    public static final MediaType JSON
            = MediaType.get("application/json; charset=utf-8");
}
