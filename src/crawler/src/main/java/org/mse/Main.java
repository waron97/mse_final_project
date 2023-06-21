package org.mse;

import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) throws InterruptedException {
        Integer numThreads = 1;
        List<Thread> threads = new ArrayList<>();
        Frontier frontier = new Frontier(Constants.frontierUrl);

        while (!frontier.isAlive()) {
            Thread.sleep(100);
        }

        for (Integer i = 0; i < numThreads; i++) {
            Thread t = new Thread(new Spider(frontier.pop()));
            threads.add(t);
            t.start();
        }

        for (Thread thread : threads) {
            try {
                thread.join();
            } catch (InterruptedException e) {
                System.out.println("Interrupted");
            }

        }
    }
}