package org.mse;

import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.util.ArrayList;
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


}
