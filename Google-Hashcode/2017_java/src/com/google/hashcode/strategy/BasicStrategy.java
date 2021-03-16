package com.google.hashcode.strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.model.CacheServer;
import com.google.hashcode.model.Request;
import com.google.hashcode.model.Video;

import java.util.Map;

/**
 * Created by janosch on 2/11/16.
 */
public class BasicStrategy implements Strategy {

    private boolean logging = false;

    private void log(String string) {
        if (logging == true)
            System.out.println(string);
    }

    @Override
    public void run(Simulation simulation, OutputExporter outputExporter) throws Exception {

        for (Map.Entry<Integer, Request> entry : simulation.getRequests().entrySet())
        {
            entry.getValue().getVideo().increaseTotalRequests(entry.getValue().getAmount());
        }

        int maxRequests = 0;
        Video videoWithMaxRequests = null;

        for (Map.Entry<Integer, Video> entry : simulation.getVideos().entrySet())
        {
            int totalRequests = entry.getValue().getTotalRequests();
            if (totalRequests > maxRequests) {
                maxRequests = totalRequests;
                videoWithMaxRequests = entry.getValue();
                log("Video with max Requests is: " + entry.getValue().getId() + " with requests " + entry.getValue().getTotalRequests());
            }
        }

        for (Map.Entry<Integer, CacheServer> entry : simulation.getCacheServers().entrySet())
        {
            if (entry.getValue().getCapacity() > videoWithMaxRequests.getSize()) {
                entry.getValue().addVideo(videoWithMaxRequests);
            }
        }
    }
}
