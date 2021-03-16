/**
 * Created by Daniel Schosser on 2/11/16.
 */
package com.google.hashcode.Importer;

import com.google.hashcode.Simulation;
import com.google.hashcode.model.CacheServer;
import com.google.hashcode.model.Endpoint;
import com.google.hashcode.model.Request;
import com.google.hashcode.model.Video;

import java.io.*;
import java.util.HashMap;

public class InputImporter {

    private boolean logging = false;
    private int videoCount = 0;
    private int endpointCount = 0;
    private int requestCount = 0;
    private int cacheServerCount = 0;
    private int cacheServerCapacity = 0;

    public Simulation readFileData(File file) {

        int lineNumber = 0;
        int currentEndpointId = 0;
        int connectedCacheServers = 0;
        int requestId = 0;
        Endpoint currentEndpoint = null;

        HashMap<Integer,Video> videos = new HashMap<>();
        HashMap<Integer,Endpoint> endpoints = new HashMap<>();
        HashMap<Integer,Request> requests = new HashMap<>();
        HashMap<Integer,CacheServer> cacheServers = new HashMap<>();

        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String line;

            // read line by line
            while ((line = br.readLine()) != null) {

                switch (lineNumber) {
                    case 0:
                        log("Line 1: " + line);
                        extractSettings(line);
                        cacheServers = createCacheServers();
                        break;
                    case 1:
                        log("Line 2: " + line);
                        videos = generateVideos(line);
                        break;
                    default:
                        log("Linie "+ lineNumber +": "+ line);
                        String[] settings = line.split(" ");

                        if(settings.length == 2) {
                            if(connectedCacheServers == 0) {
                                currentEndpoint = new Endpoint(currentEndpointId++, Integer.parseInt(settings[0]));
                                connectedCacheServers = Integer.parseInt(settings[1]);
                                if(connectedCacheServers == 0) {
                                    endpoints.put(currentEndpoint.getId(), currentEndpoint);
                                }
                            } else {
                                currentEndpoint.addCacheServer(
                                        cacheServers.get(Integer.parseInt(settings[0])),
                                        Integer.parseInt(settings[1])
                                );

                                connectedCacheServers--;
                                if(connectedCacheServers == 0) {
                                    endpoints.put(currentEndpoint.getId(), currentEndpoint);
                                }
                            }
                        } else if(settings.length == 3) {
                            Video video = videos.get(Integer.parseInt(settings[0]));
                            Endpoint endpoint = endpoints.get(Integer.parseInt(settings[1]));

                            Request request = new Request(
                                    video,
                                    endpoint,
                                    Integer.parseInt(settings[2])
                            );

                            requests.put(requestId++, request);
                        }

                        break;
                }

                lineNumber++;
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        return new Simulation(videos, endpoints, requests, cacheServers);
    }

    private HashMap<Integer,CacheServer> createCacheServers() {
        HashMap<Integer,CacheServer> cacheServers = new HashMap<>();

        for (int i=0;i<cacheServerCount;i++) {
            CacheServer server = new CacheServer(i, cacheServerCapacity);
            cacheServers.put(i, server);
        }

        return cacheServers;
    }

    private HashMap<Integer,Video> generateVideos(String line) {
        HashMap<Integer,Video> videos = new HashMap<>();

        String[] videoSizes = line.split(" ");

        int videoId = 0;
        for (String size : videoSizes) {
            Video video = new Video(videoId++, Integer.parseInt(size));
            videos.put(video.getId(),video);
        }

        return videos;
    }

    private void extractSettings(String string) {
        String[] settings = string.split(" ");

        videoCount = Integer.parseInt(settings[0]);
        endpointCount = Integer.parseInt(settings[1]);
        requestCount = Integer.parseInt(settings[2]);
        cacheServerCount = Integer.parseInt(settings[3]);
        cacheServerCapacity = Integer.parseInt(settings[4]);
    }

    private void log(String string) {
        if (logging == true)
            System.out.println(string);
    }
}