package com.google.hashcode.model;

/**
 * Created by neotreat on 23/02/2017.
 */
public class Video {

    /**
     * Video ID.
     */
    private int id;

    /**
     * Size of video in MB.
     */
    private int size;

    /**
     * Total amount of requests for this video.
     */
    private int totalRequests;


    /**
     *
     */
    public Video(int id, int size) {
        this.id = id;
        this.size = size;
    }

    /**
     *
     * @return
     */
    public int getId() {
        return id;
    }

    /**
     *
     * @return
     */
    public int getSize() {
        return size;
    }

    /**
     *
     * @return
     */
    public int getTotalRequests() {
        return totalRequests;
    }

    /**
     *
     * @param amount
     */
    public void increaseTotalRequests(int amount) {
        this.totalRequests += amount;
    }
}
