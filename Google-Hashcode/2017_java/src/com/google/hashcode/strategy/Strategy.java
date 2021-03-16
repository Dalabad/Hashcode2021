package com.google.hashcode.strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;

/**
 * Created by janosch on 2/11/16.
 */
public interface Strategy {

    void run(Simulation simulation, OutputExporter outputExporter) throws Exception;

}
