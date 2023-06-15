package org.mse;

public class Spider implements Runnable {
    private String initialUrl;

    public Spider (String url) {
        this.initialUrl = url;
    }

    @Override
    public void run() {
        System.out.println("Thread started with url " + this.initialUrl);
    }
}
