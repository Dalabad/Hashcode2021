package com.google.hashcode;

import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.model.CacheServer;
import com.google.hashcode.model.Endpoint;
import com.google.hashcode.model.Request;
import com.google.hashcode.model.Video;
import com.google.hashcode.strategy.Strategy;

import javax.print.attribute.IntegerSyntax;
import javax.print.attribute.standard.RequestingUserName;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/**
 * Created by Daniel Schosser on 2/11/16.
 */
public class Simulation {

    private Strategy strategy;
    private OutputExporter outputExporter;

    HashMap<Integer, Video> videos = new HashMap<Integer, Video>();
    HashMap<Integer, Endpoint> endpoints = new HashMap<Integer, Endpoint>();
    HashMap<Integer, Request> requests = new HashMap<Integer, Request>();
    HashMap<Integer, CacheServer> cacheServers = new HashMap<Integer, CacheServer>();
    public Strategy getStrategy() {
        return strategy;
    }
    public void setStrategy(Strategy strategy) {
        this.strategy = strategy;
    }
    public OutputExporter getOutputExporter() {
        return outputExporter;
    }
    public void setOutputExporter(OutputExporter outputExporter) {
        this.outputExporter = outputExporter;
    }
    public HashMap<Integer, Video> getVideos() {
        return videos;
    }
    public void setVideos(HashMap<Integer, Video> videos) {
        this.videos = videos;
    }
    public HashMap<Integer, Endpoint> getEndpoints() {
        return endpoints;
    }
    public void setEndpoints(HashMap<Integer, Endpoint> endpoints) {
        this.endpoints = endpoints;
    }
    public HashMap<Integer, Request> getRequests() {
        return requests;
    }
    public void setRequests(HashMap<Integer, Request> requests) {
        this.requests = requests;
    }
    public HashMap<Integer, CacheServer> getCacheServers() {
        return cacheServers;
    }
    public void setCacheServers(HashMap<Integer, CacheServer> cacheServers) {
        this.cacheServers = cacheServers;
    }


    public Simulation(HashMap videos, HashMap endpoints, HashMap requests, HashMap cacheServers) {
        this.setVideos(videos);
        this.setEndpoints(endpoints);
        this.setRequests(requests);
        this.setCacheServers(cacheServers);
    }


    public void run() {

        try {
            strategy.run(this, outputExporter);
        } catch (Exception e) {
            e.printStackTrace();
        }

    }

}
