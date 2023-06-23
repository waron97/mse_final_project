package org.mse;

import java.net.URL;
import java.util.ArrayList;
import java.util.List;

public class Robots {
    private String robotsTxt;
    private List<String> allowPatterns;
    private List<String> disallowPatterns;
    public Robots(String robotsTxt) {
        this.robotsTxt = robotsTxt;
        this.allowPatterns = new ArrayList<>();
        this.disallowPatterns = new ArrayList<>();
        parse();
    }

    private void parse() {
        if (this.robotsTxt == null) {
            return;
        }
        String[] lines = this.robotsTxt.split("\n");
        Boolean doRecord = false;
        for (String line:lines) {
            try {
                if (line.startsWith("User-agent")) {
                    String userAgent = line.split(": ")[1].strip();
                    if (userAgent.equals("*")) {
                        doRecord = true;
                    } else {
                        doRecord = false;
                    }
                } else if (line.startsWith("Allow: ")) {
                    if (!doRecord) {
                        continue;
                    }
                    allowPatterns.add(line.split(": ")[1]);
                } else if (line.startsWith("Disallow: ")) {
                    if (!doRecord) {
                        continue;
                    }
                    disallowPatterns.add(line.split(": ")[1]);
                }
            } catch (ArrayIndexOutOfBoundsException e) {
                continue;
            }

        }
    }

    public Boolean canCrawl(String url) {
        Boolean inAllowed = checkPattern(url, allowPatterns);
        Boolean inDisallowed = checkPattern(url, disallowPatterns);
        if (inAllowed) {
            return true;
        } else if (inDisallowed) {
            return false;
        }
        return true;
    }

    private Boolean checkPattern(String url, List<String> patterns) {
        Boolean inPattern = false;
        for (String pat: patterns) {
               if (pat.equals(url)) {
                   inPattern = true;
                   break;
               }
        }
        return inPattern;
    }

}
