package com.google.hashcode;

import com.google.hashcode.delivery_strategy.DeliveryStrategy;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.models.*;

import java.util.List;

/**
 * Created by Daniel Schosser on 2/11/16.
 */
public class Simulation {

    private int maxTurns, remainingTurns;
    private List<Drone> drones;
    private List<Warehouse> warehouses;
    private List<Order> orders;
    private List<ProductType> productTypes;
    private DeliveryStrategy deliveryStrategy;
    private OutputExporter outputExporter;

    public Simulation(int maxTurns, List<Drone> drones, List<Warehouse> warehouses, List<Order> orders, List<ProductType> productTypes) {
        this.maxTurns = maxTurns;
        this.remainingTurns = maxTurns;
        this.drones = drones;
        this.warehouses = warehouses;
        this.orders = orders;
        this.productTypes = productTypes;
    }

    public void setDeliveryStrategy(DeliveryStrategy deliveryStrategy) {
        this.deliveryStrategy = deliveryStrategy;
    }

    public void setOutputExporter(OutputExporter outputExporter) {
        this.outputExporter = outputExporter;
    }

    public List<Warehouse> getWarehouses() {
        return warehouses;
    }

    public List<Drone> getDrones() {
        return drones;
    }

    public List<ProductType> getProductTypes() {
        return productTypes;
    }


    public void run() {

        try {
            deliveryStrategy.deliverAll(this, outputExporter);
        } catch (Exception e) {
            e.printStackTrace();
        }

    }

    public int calculateDuration(Location start, Location destination) {
        return (int) Math.ceil(
                Math.sqrt(
                    Math.pow(
                            Math.abs(
                                start.row-destination.row), 2
                            ) +
                    Math.pow(
                            Math.abs(
                                start.column-destination.column
                            ), 2
                    ))
                );
    }


    public int getRemainingTurns() {
        return remainingTurns;
    }

    public void setRemainingTurns(int remainingTurns) {
        this.remainingTurns = remainingTurns;
    }

    public List<Order> getOrders() {
        return orders;
    }

    public OutputExporter getOutputExporter() {
        return outputExporter;
    }

    public Warehouse findNearestWarehouse(Location location) {
        Warehouse nearest = null;
        int nearestDuration = -1;
        for(Warehouse warehouse : warehouses) {
            int duration = calculateDuration(location, (Location) warehouse);
            if(duration < nearestDuration) {
                nearest = warehouse;
                nearestDuration = duration;
            }
        }

        return nearest;
    }

    public int getMaxTurns() {
        return maxTurns;
    }

    public boolean isAtWarehouse(Location location) {
        for(Warehouse warehouse : warehouses) {
            if(warehouse.row == location.row
                    && warehouse.column == location.column)
                return true;
        }
        return false;
    }
}
