package org.mse;

import com.google.gson.Gson;
import okhttp3.*;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;
import java.net.ConnectException;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Spider implements Runnable {
    private String initialUrl;
    private List<String> visited = new ArrayList<>();
    private List<String> next = new ArrayList<>();

    private OkHttpClient client = new OkHttpClient();
    private Frontier frontier;

    public Spider (Frontier frontier, String initialUrl) {
        this.frontier = frontier;
        this.initialUrl = initialUrl;
    }

    @Override
    public void run() {
        startOnPage(this.initialUrl);
        while (true) {
            startOnPage(frontier.pop());
        }
    }

    private void startOnPage(String url) {
        Logger.info(
                "startPage",
                "[Spider " + Thread.currentThread().getName() + "] starting on URL " + url,
                null
        );

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
        Logger.debug(
                "processPage",
                "[Spider " + Thread.currentThread().getName() + "] processing URL " + url,
                null
        );
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

            Thread.sleep(500);
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
            if (e instanceof ConnectException) {
                ConnectException error = (ConnectException) e;
                Map<String, String> details = new HashMap<String, String>();
                details.put("url", url);
                details.put("message", error.getMessage());
                Logger.error("processPage", "Error while processing page at " + url, details);
            } else {
                Logger.error("processPage", "Error while processing page at " + url, null);
            }
        }  catch (Exception e) {
            Logger.error("processPage", "Error while processing page at " + url, null);
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
            if (e instanceof ConnectException) {
                Map<String, String> data = new HashMap<>();
                ConnectException error = (ConnectException) e;
                data.put("message", error.getMessage());
                Logger.error("sendCrawl", "Could not register crawl data", data);
                return;
            }
            Map<String, String> data = new HashMap<>();
            data.put("message", e.getMessage());
            Logger.error("sendCrawl", "Could not register crawl data", data);
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
            Map<String, String> data = new HashMap<String, String>();
            data.put("message", e.getMessage());
            Logger.debug("getExisting", "Could not retrieve crawl data", data);
            return null;
        } catch (Exception e) {
            return null;
        }
    }

}
