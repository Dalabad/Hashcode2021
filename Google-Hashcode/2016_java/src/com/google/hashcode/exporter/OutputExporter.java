package com.google.hashcode.exporter;

import com.google.hashcode.models.Command;
import com.google.hashcode.models.Drone;
import com.google.hashcode.models.Location;
import com.google.hashcode.models.ProductType;

import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.StringJoiner;

/**
 * Created by janosch on 2/11/16.
 */
public class OutputExporter {

    private List<String> commands;
    private File file;

    public OutputExporter(File file) throws IOException {
        this.commands = new ArrayList<String>();
        this.file = file;

        if(!this.file.exists()) {
            this.file.createNewFile();
        }
    }

    public void addCommand(Drone drone, Command command, Location location, ProductType product, int number) {

        StringJoiner joiner = new StringJoiner(" ");
        joiner.add(String.valueOf(drone.getId()));
        joiner.add(String.valueOf(command.getCommand()));
        joiner.add(String.valueOf(location.getId()));
        joiner.add(String.valueOf(product.getId()));
        joiner.add(String.valueOf(number));

        this.commands.add(joiner.toString());
    }

    public void writeOutput() {
        try {
            PrintWriter printWriter = new PrintWriter(new FileWriter(this.file));

            printWriter.println(String.valueOf(this.commands.size()));

            for (String command : this.commands) {
                printWriter.println(command);
            }

            printWriter.flush();
            printWriter.close();

        } catch (java.io.IOException e) {
            System.out.println("Could not write output file");
            e.printStackTrace();
        }
    }

}
