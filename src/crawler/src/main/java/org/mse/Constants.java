package org.mse;

import okhttp3.MediaType;

public class Constants {
    public static String frontierUrl = "http://frontier-acceptor:3000";
    public static String acceptorUrl = "http://frontier-acceptor:3000/acceptor";
    public static final MediaType JSON
            = MediaType.get("application/json; charset=utf-8");
}
