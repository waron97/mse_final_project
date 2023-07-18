package org.mse;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

import java.io.IOException;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.Arrays;
import java.util.List;

public class Util {
    static void loadPageHtml() {}

    static boolean isLegalUrl(String url) {
        if (url == null) {
            return false;
        }
        List<String> filePatterns = Arrays.asList(".pdf", ".png", ".jpg");
        for (String pattern: filePatterns) {
            if (url.endsWith(pattern)) {
                return false;
            }
        }
        try {
            URL asUrl = new URL(url);
            String protocol = asUrl.getProtocol();
            return protocol.equals("http")  || protocol.equals("https");
        } catch (MalformedURLException e) {
            return false;
        }

    }

    static Robots getRobots(String url) {
        if (url == null) {
            return new Robots(null);
        }

        OkHttpClient client = new OkHttpClient();

        Request req = null;
        try {
            req = new Request.Builder().url(url + "/robots.txt").build();
        } catch (Exception exception) {
            return new Robots(null);
        }
        try (Response res = client.newCall(req).execute()) {
            return new Robots(res.body().string());
        } catch (IOException e) {
            return new Robots(null);
        }
    }

    static String getBaseUrl(String url) {
        if (url == null) {
            return null;
        }
        try {
            URL asUrl = new URL(url);
            return asUrl.getProtocol() + "://" + asUrl.getHost();
        } catch (MalformedURLException e) {
            return null;
        }
    }

    static String getRelativeUrl(String url) {
        if (url == null) {
            return null;
        }
        try {
            URL asUrl = new URL(url);
            String relativeUrl = asUrl.getPath();
            String query = asUrl.getQuery();
            if (query != null) {
                relativeUrl += "?" + query;
            }
            return relativeUrl;

        } catch (MalformedURLException e) {
            return null;
        }
    }

    static boolean isAbsoluteUrl(String url) {
        if (url == null) {
            return false;
        }
        return url.indexOf(":") > 1;
    }
}
