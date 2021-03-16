package com.google.hashcode.exporter;

import com.google.hashcode.model.CacheServer;
import com.google.hashcode.model.Request;
import com.google.hashcode.model.Video;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;

/**
 * Created by janosch on 2/11/16.
 */
public class OutputExporter {

    private File file;

    public OutputExporter(File file) throws IOException {
        this.file = file;

        if(!this.file.exists()) {
            this.file.createNewFile();
        }
    }

    public void writeOutput(HashMap<Integer, CacheServer> cacheServers) {
        try {

            PrintWriter printWriter = new PrintWriter(new FileWriter(this.file));

            printWriter.println(String.valueOf(cacheServers.size()));

            for (Map.Entry<Integer, CacheServer> cacheServerEntry : cacheServers.entrySet())
            {
                String line = String.valueOf(cacheServerEntry.getValue().getId());

                for (Map.Entry<Integer, Video> entry : cacheServerEntry.getValue().getVideos().entrySet())
                {
                    line += " " + String.valueOf(entry.getValue().getId());
                }

                printWriter.println(line);
            }

            printWriter.flush();
            printWriter.close();

        } catch (java.io.IOException e) {
            System.out.println("Could not write output file");
            e.printStackTrace();
        }
    }

}
