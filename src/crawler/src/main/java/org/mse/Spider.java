package org.mse;

import com.google.gson.Gson;
import okhttp3.*;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Spider implements Runnable {
    private String initialUrl;
    private List<String> visited = new ArrayList<>();
    private List<String> next = new ArrayList<>();

    private OkHttpClient client = new OkHttpClient();

    public Spider (String url) {
        this.initialUrl = url;
    }

    @Override
    public void run() {
        System.out.println("Thread starting with url " + this.initialUrl);
        processPage(this.initialUrl);
        while (!next.isEmpty()) {
            String nextPage = next.remove(0);
            Boolean isLegal = Util.isLegalUrl(nextPage);
            if (visited.contains(nextPage) || !Util.isLegalUrl(nextPage)) {
                continue;
            }
            processPage(nextPage);
        }
    }

    private void processPage(String url) {
        System.out.println("[" + Thread.currentThread().getName() + "] starting on " + url);
        Robots robots = Util.getRobots(url);
        String absoluteUrl = Util.getBaseUrl(url);
        try {
            Document doc = Jsoup.connect(url).get();
            String title = HtmlExtractor.getPageTitle(doc);
            List<AnchorTag> tags = HtmlExtractor.getLinks(doc, absoluteUrl);
            String bodyText = HtmlExtractor.getBodyText(doc);
            String mainText = HtmlExtractor.getMainText(doc);
            String description = HtmlExtractor.getDescription(doc);
            List<String> keywords = HtmlExtractor.getKeywords(doc);
            String rawHtml = doc.html();

            sendCrawl(
                    url,
                    title,
                    tags,
                    bodyText,
                    mainText,
                    description,
                    keywords,
                    rawHtml
            );

            next.addAll(
                    tags.stream()
                            .map(a -> a.getHref())
                            .filter(a -> !this.next.contains(a))
                            .filter(a -> !this.visited.contains(a))
                            .toList()
            );

        } catch (IOException e) {

        }
    }

    private void sendCrawl(String url, String title, List<AnchorTag> tags, String bodyText, String mainText, String description, List<String> keywords, String rawHtml) {
        CrawlResult result = new CrawlResult(
                url,
                title,
                tags,
                bodyText,
                mainText,
                description,
                keywords,
                rawHtml
        );
        Gson gson = new Gson();
        String payload = gson.toJson(result);
        RequestBody body = RequestBody.create(payload, Constants.JSON);
        Request request = new Request
                .Builder()
                .url(Constants.acceptorUrl)
                .post(body)
                .build();

        try (Response response = client.newCall(request).execute()) {
            // success
        } catch (IOException e) {
            System.out.println("[ERROR] could not register crawl data");
        }



    }
}
