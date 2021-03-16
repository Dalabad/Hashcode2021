package com.google.hashcode.model;

import java.util.HashMap;
import java.util.Iterator;

/**
 * Created by neotreat on 23/02/2017.
 */
public class Endpoint {

    /**
     * Endpoint id.
     */
    private int id;

    /**
     * Latency to the mothership datacenter in ms.
     */
    private int datacenterLatency;

    /**
     * List of connected cache servers.
     */
    private HashMap<Integer, CacheServer> cacheServers;

    /**
     * List of latencies to the respective cache server.
     */
    private HashMap<Integer, Integer> cacheServerLatencies;

    /**
     *
     * @param id
     * @param datacenterLatency
     */
    public Endpoint(int id, int datacenterLatency) {
        this.id = id;
        this.datacenterLatency = datacenterLatency;
        this.cacheServers = new HashMap<Integer, CacheServer>();
        this.cacheServerLatencies = new HashMap<Integer, Integer>();
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
    public int getDatacenterLatency() {
        return datacenterLatency;
    }

    /**
     *
     * @return
     */
    public HashMap<Integer, CacheServer> getCacheServers() {
        return cacheServers;
    }

    /**
     *
     * @param cacheServer
     */
    public void addCacheServer(CacheServer cacheServer, int latency) {
        this.cacheServers.put(cacheServer.getId(), cacheServer);
        this.cacheServerLatencies.put(cacheServer.getId(), latency);
    }

    /**
     *
     * @param cacheServer
     * @return
     */
    public int getLatencyForCacheServer(CacheServer cacheServer) {
       return this.cacheServerLatencies.get(cacheServer.getId());
    }

    /**
     *
     * @param cacheServerId
     * @return
     */
    public int getLatencyForCacheServer(int cacheServerId) {
        return this.cacheServerLatencies.get(cacheServerId);
    }

    /**
     *
     * @return
     */
    public CacheServer getMinLatencyCacheServer() {

        Integer minLatency = Integer.MAX_VALUE;
        Integer cacheServerId = new Integer(-1);
        Iterator iterator = this.cacheServerLatencies.entrySet().iterator();

        while (iterator.hasNext()) {
            HashMap.Entry entry = (HashMap.Entry) iterator.next();
            Integer currentCacheServerId = (Integer) entry.getKey();
            Integer latency = (Integer) entry.getValue();

            if(latency < minLatency) {
                minLatency = latency;
                cacheServerId = currentCacheServerId;
            }
        }

        if(cacheServerId == -1) {
            return null;
        }

        return (CacheServer) this.cacheServers.get(cacheServerId);
    }

    /**
     *
     * @param video
     * @return
     */
    public CacheServer getMinLatencyCacheServer(Video video) {

        Integer minLatency = Integer.MAX_VALUE;
        Integer cacheServerId = new Integer(-1);
        Iterator iterator = this.cacheServerLatencies.entrySet().iterator();

        while (iterator.hasNext()) {
            HashMap.Entry entry = (HashMap.Entry) iterator.next();
            Integer currentCacheServerId = (Integer) entry.getKey();
            Integer latency = (Integer) entry.getValue();

            if(latency < minLatency && this.cacheServers.get(currentCacheServerId).getRemainingCapacity() >= video.getSize()) {
                minLatency = latency;
                cacheServerId = currentCacheServerId;
            }
        }

        if(cacheServerId == -1) {
            return null;
        }

        return (CacheServer) this.cacheServers.get(cacheServerId);
    }
}
