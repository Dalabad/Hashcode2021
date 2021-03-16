package com.google.hashcode.model;

import java.util.HashMap;
import java.util.List;

/**
 * Created by neotreat on 23/02/2017.
 */
public class CacheServer {

    /**
     * Cache server ID.
     */
    private int id;

    /**
     * Capacity of the cache server in MB.
     */
    private int capacity;

    /**
     * Capacity left on the cache server in MB.
     */
    private int remainingCapacity;

    /**
     * List of currently stored videos.
     */
    private HashMap<Integer, Video> videos;

    /**
     *
     */
    public CacheServer(int id, int capacity) {
        this.id = id;
        this.capacity = capacity;
        this.remainingCapacity = capacity;
        this.videos = new HashMap<Integer, Video>();
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
    public int getCapacity() {
        return capacity;
    }

    /**
     *
     * @return
     */
    public int getRemainingCapacity() {
        return remainingCapacity;
    }

    /**
     *
     * @return
     */
    public HashMap<Integer, Video> getVideos() {
        return videos;
    }

    /**
     *
     * @param video
     */
    public void addVideo(Video video) throws Exception {

        // Check if there's enough space on the cache server left
        if (this.remainingCapacity < video.getSize()) {
            throw new Exception("Not enough space left on the cache server.");
        }

        // Add the video if there's enough space
        this.videos.put(new Integer(video.getId()), video);
        this.remainingCapacity -= video.getSize();
    }
}
