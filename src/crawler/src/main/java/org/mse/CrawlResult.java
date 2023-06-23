package org.mse;

import java.util.List;

public class CrawlResult {
    private String url;
    private String title;

    private String date;

    public String getDate() {
        return date;
    }

    public void setDate(String date) {
        this.date = date;
    }

    private List<AnchorTag> links;
    private String bodyText;
    private String mainText;
    private String description;
    private List<String> keywords;

    private String rawHtml;

    public String getRawHtml() {
        return rawHtml;
    }

    public void setRawHtml(String rawHtml) {
        this.rawHtml = rawHtml;
    }

    public CrawlResult(String url,
                       String title,
                       List<AnchorTag> tags,
                       String bodyText,
                       String mainText,
                       String description,
                       List<String> keywords,
                       String rawHtml
    ) {
        this.url = url;
        this.title = title;
        this.links = tags;
        this.bodyText = bodyText;
        this.mainText = mainText;
        this.description = description;
        this.keywords = keywords;
        this.rawHtml = rawHtml;

    }

    public List<AnchorTag> getLinks() {
        return links;
    }

    public String getBodyText() {
        return bodyText;
    }

    public String getDescription() {
        return description;
    }

    public List<String> getKeywords() {
        return keywords;
    }

    public String getMainText() {
        return mainText;
    }

    public String getTitle() {
        return title;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public void setLinks(List<AnchorTag> links) {
        this.links = links;
    }

    public void setBodyText(String bodyText) {
        this.bodyText = bodyText;
    }

    public void setMainText(String mainText) {
        this.mainText = mainText;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public void setKeywords(List<String> keywords) {
        this.keywords = keywords;
    }
}
