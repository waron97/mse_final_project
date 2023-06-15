package org.mse;

public class AnchorTag {
    private String text;
    private String href;


    public AnchorTag(String text, String href) {
        this.text = text;
        this.href = href;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public String getHref() {
        return href;
    }

    public void setHref(String href) {
        this.href = href;
    }

    @Override
    public String toString() {
        return "AnchorTag{" +
                "text='" + text + '\'' +
                ", href='" + href + '\'' +
                '}';
    }
}
