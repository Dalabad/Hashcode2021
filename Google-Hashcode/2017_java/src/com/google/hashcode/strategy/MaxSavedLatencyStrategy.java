package com.google.hashcode.strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.model.CacheServer;
import com.google.hashcode.model.Endpoint;
import com.google.hashcode.model.Request;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;

/**
 * Created by janosch on 2/11/16.
 */
public class MaxSavedLatencyStrategy implements Strategy {


    @Override
    public void run(Simulation simulation, OutputExporter outputExporter) throws Exception {

        // Sort requests my getMaxSavedLatency metric
        List<Request> requests = new ArrayList<>(simulation.getRequests().values());
        Collections.sort(requests);

        for (Request request: requests) {

            Endpoint endpoint = request.getEndpoint();
            CacheServer minLatencyCacheServer = endpoint.getMinLatencyCacheServer(request.getVideo());

            try {
                minLatencyCacheServer.addVideo(request.getVideo());
            } catch(Exception e) {
                //System.out.println("Failed to add video, but will continue...");
            }

            //
        }

        System.out.println("-- Finished Strategy --");

    }
}
