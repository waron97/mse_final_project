package org.mse;

import com.google.gson.Gson;
import okhttp3.*;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.List;

public class Spider implements Runnable {
    private List<String> visited = new ArrayList<>();
    private List<String> next = new ArrayList<>();

    private OkHttpClient client = new OkHttpClient();
    private Frontier frontier;

    public Spider (Frontier frontier) {
        this.frontier = frontier;
    }

    @Override
    public void run() {
        while (true) {
            startOnPage(frontier.pop());
        }
    }

    private void startOnPage(String url) {
        processPage(url);
        while (!next.isEmpty()) {
            String nextPage = next.remove(0);
            if (visited.contains(nextPage) || !Util.isLegalUrl(nextPage)) {
                continue;
            }
            processPage(nextPage);
        }
    }

    private void processPage(String url) {

        Robots robots = Util.getRobots(url);
        String absoluteUrl = Util.getBaseUrl(url);

        if (!robots.canCrawl(Util.getRelativeUrl(url))) {
            return;
        }

        try {
            CrawlResult existing = getExisting(url);
            Boolean shouldDownload = shouldFetchPage(existing);
            Document doc = null;

            if (!shouldDownload) {
                doc = Jsoup.parse(existing.getRawHtml());

                List<AnchorTag> tags = HtmlExtractor.getLinks(doc, absoluteUrl);
                next.addAll(
                        tags.stream()
                                .map(a -> a.getHref())
                                .filter(a -> !this.next.contains(a))
                                .filter(a -> !this.visited.contains(a))
                                .toList()
                );
                return;
            }

            doc = Jsoup.connect(url).get();


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

        } catch (Exception e) {

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

    private Boolean shouldFetchPage(CrawlResult existing) {
        if (existing == null) {
            return true;
        }
        return false;
    }

    private CrawlResult getExisting(String url) {
        Gson gson = new Gson();
        String encodedUrl = url;
        try {
            encodedUrl = URLEncoder.encode(url, "utf-8");
        } catch (Exception e) {}
        Request request = new Request.Builder()
                .url(Constants.acceptorUrl + "/" + encodedUrl)
                .build();

        try (Response response = client.newCall(request).execute()) {
            // success
            String body = response.body().string();
            CrawlResult result = gson.fromJson(body, CrawlResult.class);
            return result;
        } catch (IOException e) {
            System.out.println("[ERROR] could not register crawl data");
            return null;
        } catch (Exception e) {
            return null;
        }
    }

}
