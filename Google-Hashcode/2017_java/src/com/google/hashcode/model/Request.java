package com.google.hashcode.model;

import com.sun.org.apache.regexp.internal.RE;

/**
 * Created by neotreat on 23/02/2017.
 */
public class Request implements Comparable<Request> {

    /**
     * Requested video.
     */
    private Video video;

    /**
     * Endpoint from which the video is requested.
     */
    private Endpoint endpoint;

    /**
     * Amount of times the video is requested from this endpoint.
     */
    private int amount;

    /**
     *
     * @param video
     * @param endpoint
     * @param amount
     */
    public Request(Video video, Endpoint endpoint, int amount) {
        this.video = video;
        this.endpoint = endpoint;
        this.amount = amount;
    }

    /**
     *
     * @return
     */
    public Video getVideo() {
        return video;
    }

    /**
     *
     * @return
     */
    public Endpoint getEndpoint() {
        return endpoint;
    }

    /**
     *
     * @return
     */
    public int getAmount() {
        return amount;
    }

    /**
     *
     */
    public int getMaxSavedLatency(Video video) {
        int datacenterLatency = 0;
        int minCacheServerLatency = 0;

        CacheServer minLatencyCacheServer = this.endpoint.getMinLatencyCacheServer(video);

        datacenterLatency = this.endpoint.getDatacenterLatency();

        if(minLatencyCacheServer == null) {
            minCacheServerLatency = datacenterLatency;
        } else {
            minCacheServerLatency = this.endpoint.getLatencyForCacheServer(minLatencyCacheServer);
        }

        return (datacenterLatency - minCacheServerLatency) * this.amount;
    }

    @Override
    public int compareTo(Request o) {
        return Integer.compare(o.getMaxSavedLatency(o.getVideo()), this.getMaxSavedLatency(this.getVideo()));
    }
}
