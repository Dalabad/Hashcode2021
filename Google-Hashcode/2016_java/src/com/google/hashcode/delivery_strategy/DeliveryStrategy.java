package com.google.hashcode.delivery_strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.models.Order;

/**
 * Created by janosch on 2/11/16.
 */
public interface DeliveryStrategy {

    void deliverAll(Simulation simulation, OutputExporter outputExporter) throws Exception;

}
