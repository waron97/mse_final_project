package org.mse;

import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class HtmlExtractor {
    static String getPageTitle(Document doc) {
        return doc.title();
    }

    static List<AnchorTag> getLinks(Document doc, String baseUrl) {
        List<AnchorTag> links = new ArrayList<>();
        Elements elements = doc.select("a");
        for (Element element : elements) {
            String text = element.text();
            String href = element.attr("href");
            if (href == null || href.equals("")) {
                continue;
            } else if (href.startsWith("#")) {
                continue;
            } else if (!Util.isAbsoluteUrl(href)) {
                href = baseUrl + href;
            }
            links.add(new AnchorTag(text, href));
        }
        return links;
    }

    static String getBodyText(Document doc) {
        return doc.select("body").first().text();
    }

    static String getMainText(Document doc) {
        Elements element = doc.select("main");
        if (!element.isEmpty()) {
            element.first().text();
        }
        return "";
    }

    static String getDescription(Document doc) {
        Elements elements = doc.select("meta[name='description']");
        if (elements.isEmpty()) {
            return "";
        }
        return elements.first().attr("content");
    }

    static List<String> getKeywords(Document doc) {
        Elements elements = doc.select("meta[name='keywords']");
        if (elements.isEmpty()) {
            return new ArrayList<>();
        }
        String content = elements.first().attr("content");
        return Arrays.asList(content.split(","))
                .stream()
                .map(item -> item.strip())
                .toList();
    }


}
