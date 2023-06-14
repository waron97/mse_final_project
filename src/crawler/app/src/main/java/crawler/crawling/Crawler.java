package crawler.crawling;

import java.util.ArrayList;
import java.util.List;

public class Crawler implements Runnable {
    private String initialPage;
    List<String> visited;
    List<String> detected;

    public Crawler(String initialPage) {
        this.initialPage = initialPage;
        this.visited = new ArrayList<>();
        this.detected = new ArrayList<>();
    }

    @Override
    public void run() {

    }
}