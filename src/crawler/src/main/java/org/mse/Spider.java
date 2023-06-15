package org.mse;

import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Spider implements Runnable {
    private String initialUrl;
    private List<String> visited = new ArrayList<>();
    private List<String> next = new ArrayList<>();

    public Spider (String url) {
        this.initialUrl = url;
    }

    @Override
    public void run() {
        processPage(this.initialUrl);
    }

    private void processPage(String url) {
        Robots robots = Util.getRobots(this.initialUrl);
        String absoluteUrl = Util.getBaseUrl(url);
        try {
            Document doc = Jsoup.connect(this.initialUrl).get();
            String title = HtmlExtractor.getPageTitle(doc);
            List<AnchorTag> tags = HtmlExtractor.getLinks(doc, absoluteUrl);
            System.out.println(title);
            System.out.println(tags);

        } catch (IOException e) {

        }
    }
}
